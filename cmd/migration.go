package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/database"
	"github.com/urfave/cli/v2"
)

var MigrationCommand = &cli.Command{
	Name:  "migrate",
	Usage: "Run a specific migration file",
	Flags: []cli.Flag{&cli.StringFlag{Name: "file", Required: true}},
	Action: func(c *cli.Context) error {
		name := c.String("file")
		fmt.Println("🚀 Migrate:", name)
		m := database.NewMigration(config.DB)
		return m.RunMigration(name)
	},
}

var RollbackCommand = &cli.Command{
	Name:  "rollback",
	Usage: "Rollback a specific migration",
	Flags: []cli.Flag{&cli.StringFlag{Name: "file", Required: true}},
	Action: func(c *cli.Context) error {
		name := c.String("file")
		fmt.Println("🔄 Rollback:", name)
		m := database.NewMigration(config.DB)
		return m.RollbackMigration(name)
	},
}

var MakeMigrationCommand = &cli.Command{
	Name:  "make:migration",
	Usage: "Create new migration template",
	Action: func(c *cli.Context) error {
		if c.Args().Len() < 1 {
			return fmt.Errorf("nama migration dibutuhkan")
		}
		return CreateMigration(c.Args().First())
	},
}

var MigrateAllCommand = &cli.Command{
	Name:  "migrate:all",
	Usage: "Run all pending migrations",
	Action: func(c *cli.Context) error {
		fmt.Println("🚀 Migrate all")
		m := database.NewMigration(config.DB)
		return m.RunAllMigrations()
	},
}

var RollbackAllCommand = &cli.Command{
	Name:  "rollback:all",
	Usage: "Rollback all batches",
	Action: func(c *cli.Context) error {
		fmt.Println("🔄 Rollback all")
		m := database.NewMigration(config.DB)
		return m.RunAllRollbacks()
	},
}

var RollbackBatchCommand = &cli.Command{
	Name:  "rollback:batch",
	Usage: "Rollback specific batch",
	Flags: []cli.Flag{&cli.IntFlag{Name: "batch"}},
	Action: func(c *cli.Context) error {
		b := c.Int("batch")
		m := database.NewMigration(config.DB)
		if b == 0 {
			return m.RollbackLastBatch()
		}
		return m.RollbackBatch(b)
	},
}

var MigrateFreshCommand = &cli.Command{
	Name:  "migrate:fresh",
	Usage: "Reset and re-run all migrations",
	Action: func(c *cli.Context) error {
		fmt.Println("🔄 Fresh: rollback all then migrate all")
		m := database.NewMigration(config.DB)
		if err := m.RunAllRollbacks(); err != nil {
			return err
		}
		return m.RunAllMigrations()
	},
}

func CreateMigration(name string) error {
	ts := time.Now().Format("20060102150405")
	fname := fmt.Sprintf("%s_%s.sql", ts, name)
	dir, _ := os.Getwd()
	path := fmt.Sprintf("%s/internal/database/migrations/%s", dir, fname)
	up, down := getMigrationTemplate(name)
	content := fmt.Sprintf("%s\n%s\n%s\n%s", upMarker, up, downMarker, down)
	return os.WriteFile(path, []byte(content), 0644)
}

const (
	upMarker   = "--- up"
	downMarker = "--- down"
)

func getMigrationTemplate(name string) (string, string) {
	if strings.HasPrefix(name, "create_") {
		tbl := strings.TrimPrefix(name, "create_")
		tbl = strings.TrimSuffix(tbl, "_table")
		up := fmt.Sprintf(`CREATE TABLE %s (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL DEFAULT NULL
);`, tbl)
		down := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tbl)
		return up, down
	}

	if strings.HasPrefix(name, "alter_") {
		tbl := strings.TrimPrefix(name, "alter_")
		tbl = strings.TrimSuffix(tbl, "_table")
		up := fmt.Sprintf(`ALTER TABLE %s 
-- ADD COLUMN new_column_name DATA_TYPE;
`, tbl)
		down := fmt.Sprintf(`ALTER TABLE %s 
-- DROP COLUMN new_column_name;
`, tbl)
		return up, down
	}

	// Default fallback
	return "-- up SQL here", "-- down SQL here"
}
