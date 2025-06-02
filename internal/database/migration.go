// Don't modify this file directly
package database

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"gorm.io/gorm"
)

const (
	upMarker   = "--- up"
	downMarker = "--- down"
)

var (
	ErrMigrationFileNotFound   = fmt.Errorf("migration file not found")
	ErrMigrationAlreadyApplied = fmt.Errorf("migration already applied")
	ErrInvalidMigrationFormat  = fmt.Errorf("invalid Migration format")
	ErrMigrationFailed         = fmt.Errorf("migration failed")
	ErrRollbackFailed          = fmt.Errorf("rollback failed")
)

type Migration struct {
	db *gorm.DB
}

func NewMigration(db *gorm.DB) Migration {
	return Migration{db: db}
}

func (m *Migration) ensureMigrationsTable() error {
	if err := m.db.Exec(`
        CREATE TABLE IF NOT EXISTS migrations (
            id INT PRIMARY KEY AUTO_INCREMENT,
            filename VARCHAR(255) NOT NULL,
            batch INT NOT NULL,
            migrated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `).Error; err != nil {
		return fmt.Errorf("failed to ensure migrations table: %w", err)
	}
	return nil
}

func (m *Migration) getLastBatch() (int, error) {
	var res struct{ Batch int }
	if err := m.db.Raw("SELECT COALESCE(MAX(batch),0) AS batch FROM migrations").Scan(&res).Error; err != nil {
		return 0, fmt.Errorf("failed to get last batch: %w", err)
	}
	return res.Batch, nil
}

func (m *Migration) isMigrationApplied(filename string) (bool, error) {
	var cnt int64
	if err := m.db.Raw("SELECT COUNT(*) FROM migrations WHERE filename = ?", filename).Scan(&cnt).Error; err != nil {
		return false, fmt.Errorf("failed to check if migration is applied: %w", err)
	}
	return cnt > 0, nil
}

func (m *Migration) parseMigrationFile(content string) (upStmts, downStmts []string) {
	parts := strings.Split(content, downMarker)
	upPart := parts[0]
	downPart := ""
	if len(parts) > 1 {
		downPart = parts[1]
	}
	upPart = strings.Replace(upPart, upMarker, "", 1)
	return m.parseSQLStatements(upPart), m.parseSQLStatements(downPart)
}

func (m *Migration) RunMigration(filename string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/internal/database/migrations/%s.sql", wd, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrMigrationFileNotFound, err)
	}
	ups, _ := m.parseMigrationFile(string(data))
	for _, sql := range ups {
		if err := m.db.Exec(sql).Error; err != nil {
			return fmt.Errorf("%w: %v", ErrMigrationFailed, err)
		}
	}
	return nil
}

func (m *Migration) RollbackMigration(filename string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/internal/database/migrations/%s.sql", wd, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrMigrationFileNotFound, err)
	}
	_, downs := m.parseMigrationFile(string(data))
	for _, sql := range downs {
		if err := m.db.Exec(sql).Error; err != nil {
			return fmt.Errorf("%w: %v", ErrRollbackFailed, err)
		}
	}
	return nil
}

func (m *Migration) parseSQLStatements(content string) []string {
	lines := strings.Split(content, "\n")
	cleanedLines := []string{}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if idx := strings.Index(line, "--"); idx != -1 {
			line = line[:idx]
		}
		if idx := strings.Index(line, "#"); idx != -1 {
			line = line[:idx]
		}

		line = strings.TrimSpace(line)
		if line != "" {
			cleanedLines = append(cleanedLines, line)
		}
	}

	cleanedContent := strings.Join(cleanedLines, " ")
	rawStatements := strings.Split(cleanedContent, ";")

	finalStatements := []string{}
	for _, stmt := range rawStatements {
		s := strings.TrimSpace(stmt)
		if s != "" {
			finalStatements = append(finalStatements, s)
		}
	}

	return finalStatements
}

func (m *Migration) RunAllMigrations() error {
	if err := m.ensureMigrationsTable(); err != nil {
		return fmt.Errorf("failed to run all migrations: %w", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	files, err := os.ReadDir(fmt.Sprintf("%s/internal/database/migrations", wd))
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}
	last, err := m.getLastBatch()
	if err != nil {
		return fmt.Errorf("failed to get last batch: %w", err)
	}
	batch := last + 1
	var toRun []string
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".sql") {
			name := strings.TrimSuffix(f.Name(), ".sql")
			applied, err := m.isMigrationApplied(name)
			if err != nil {
				return fmt.Errorf("failed to check Migration status: %w", err)
			}
			if !applied {
				toRun = append(toRun, name)
			}
		}
	}
	sort.Strings(toRun)
	for _, name := range toRun {
		log.Println("🚀 Running", name)
		if err := m.RunMigration(name); err != nil {
			return fmt.Errorf("failed to run Migration %s: %w", name, err)
		}
		if err := m.db.Exec("INSERT INTO migrations(filename,batch) VALUES(?,?)", name, batch).Error; err != nil {
			return fmt.Errorf("failed to record Migration: %w", err)
		}
	}
	log.Printf("✅ Batch %d applied", batch)
	return nil
}

func (m *Migration) RunAllRollbacks() error {
	if err := m.ensureMigrationsTable(); err != nil {
		return fmt.Errorf("failed to run all rollbacks: %w", err)
	}
	last, err := m.getLastBatch()
	if err != nil {
		return fmt.Errorf("failed to get last batch: %w", err)
	}
	for b := last; b >= 1; b-- {
		if err := m.RollbackBatch(b); err != nil {
			return fmt.Errorf("failed to rollback batch %d: %w", b, err)
		}
	}
	return nil
}

func (m *Migration) RollbackBatch(batch int) error {
	if err := m.ensureMigrationsTable(); err != nil {
		return fmt.Errorf("failed to rollback batch: %w", err)
	}
	var rows []struct{ Filename string }
	if err := m.db.Raw("SELECT filename FROM migrations WHERE batch=? ORDER BY id DESC", batch).Scan(&rows).Error; err != nil {
		return fmt.Errorf("failed to get migrations for batch %d: %w", batch, err)
	}
	for _, r := range rows {
		log.Println("🔄 Rollback", r.Filename)
		if err := m.RollbackMigration(r.Filename); err != nil {
			return fmt.Errorf("failed to rollback %s: %w", r.Filename, err)
		}
		if err := m.db.Exec("DELETE FROM migrations WHERE filename=?", r.Filename).Error; err != nil {
			return fmt.Errorf("failed to delete Migration record: %w", err)
		}
	}
	log.Printf("✅ Batch %d rolled back", batch)
	return nil
}

func (m *Migration) RollbackLastBatch() error {
	last, err := m.getLastBatch()
	if err != nil {
		return fmt.Errorf("failed to get last batch: %w", err)
	}
	if last == 0 {
		log.Println("⚠️ No batch to rollback")
		return nil
	}
	return m.RollbackBatch(last)
}

func (m *Migration) FreshMigrations() error {
	if err := m.ensureMigrationsTable(); err != nil {
		return fmt.Errorf("failed to ensure migrations table: %w", err)
	}
	if err := m.db.Exec("TRUNCATE migrations").Error; err != nil {
		return fmt.Errorf("failed to truncate migrations table: %w", err)
	}
	return m.RunAllMigrations()
}
