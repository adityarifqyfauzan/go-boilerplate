// Don't modify this file directly
package database

import (
	"fmt"
	"log"
	"sort"
	"time"

	"gorm.io/gorm"
)

type Seeder struct {
	Name     string
	Run      func(db *gorm.DB) error
	Rollback func(db *gorm.DB) error
	Batch    int64
}

var SeederList = []Seeder{
	{},
}

type seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) seeder {
	return seeder{db: db}
}

func (s *seeder) ensureSeedsTable() error {
	return s.db.Exec(`
		CREATE TABLE IF NOT EXISTS seeds (
			id INT PRIMARY KEY AUTO_INCREMENT,
			filename VARCHAR(255) NOT NULL,
			batch BIGINT NOT NULL,
			seeded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`).Error
}

func (s *seeder) getLastSeedBatch() (int64, error) {
	var res struct{ Batch int64 }
	if err := s.db.
		Raw("SELECT COALESCE(MAX(batch),0) AS batch FROM seeds").
		Scan(&res).Error; err != nil {
		return 0, err
	}
	return res.Batch, nil
}

func (s *seeder) isSeedApplied(name string) (bool, error) {
	var count int64
	if err := s.db.
		Raw("SELECT COUNT(*) FROM seeds WHERE filename = ?", name).
		Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *seeder) RunAllSeeders() error {
	if err := s.ensureSeedsTable(); err != nil {
		return err
	}
	_, err := s.getLastSeedBatch()
	if err != nil {
		return err
	}
	newBatch := time.Now().Unix()

	var pending []Seeder
	for _, seeder := range SeederList {
		applied, err := s.isSeedApplied(seeder.Name)
		if err != nil {
			return err
		}
		if !applied {
			seeder.Batch = newBatch
			pending = append(pending, seeder)
		}
	}
	sort.Slice(pending, func(i, j int) bool { return pending[i].Name < pending[j].Name })

	for _, seeder := range pending {
		log.Println("🌱 Seeding:", seeder.Name)
		if err := seeder.Run(s.db); err != nil {
			return fmt.Errorf("failed to run seeder %s: %w", seeder.Name, err)
		}
		if err := s.db.
			Exec("INSERT INTO seeds (filename, batch) VALUES (?, ?)", seeder.Name, seeder.Batch).
			Error; err != nil {
			return err
		}
	}
	log.Printf("✅ Seed batch %d applied.\n", newBatch)
	return nil
}

func (s *seeder) RollbackSeedBatch(batch int64) error {
	if err := s.ensureSeedsTable(); err != nil {
		return err
	}

	var rows []struct{ Filename string }
	if err := s.db.
		Raw("SELECT filename FROM seeds WHERE batch = ? ORDER BY id DESC", batch).
		Scan(&rows).Error; err != nil {
		return err
	}
	if len(rows) == 0 {
		log.Printf("⚠️ No seeders in batch %d\n", batch)
		return nil
	}

	for _, r := range rows {
		log.Println("🔄 Rolling back seeder:", r.Filename)
		for _, seeder := range SeederList {
			if seeder.Name == r.Filename {
				if seeder.Rollback != nil {
					if err := seeder.Rollback(s.db); err != nil {
						return fmt.Errorf("rollback seeder %s failed: %w", seeder.Name, err)
					}
				}
				break
			}
		}
		if err := s.db.
			Exec("DELETE FROM seeds WHERE filename = ? AND batch = ?", r.Filename, batch).
			Error; err != nil {
			return err
		}
	}
	log.Printf("✅ Seeder batch %d rolled back.\n", batch)
	return nil
}

func (s *seeder) RollbackLastSeedBatch() error {
	b, err := s.getLastSeedBatch()
	if err != nil {
		return err
	}
	if b == 0 {
		log.Println("⚠️ No seed batch to rollback.")
		return nil
	}
	return s.RollbackSeedBatch(b)
}
