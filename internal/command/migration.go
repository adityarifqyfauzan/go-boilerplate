package command

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/helper"
	"github.com/ahmadfaizk/schema"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v2"
)

var MigrateUpCommand = &cli.Command{
	Name:  "migrate:up",
	Usage: "Run migrations",
	Action: func(c *cli.Context) error {
		if !helper.ConfirmProductionAction() {
			return nil
		}

		return Up()
	},
}

var MigrateDownCommand = &cli.Command{
	Name:  "migrate:down",
	Usage: "Run migrations",
	Action: func(c *cli.Context) error {
		if !helper.ConfirmProductionAction() {
			return nil
		}

		return Down()
	},
}

var MigrateRefreshCommand = &cli.Command{
	Name:  "migrate:refresh",
	Usage: "Run migrations",
	Action: func(c *cli.Context) error {
		if !helper.ConfirmProductionAction() {
			return nil
		}

		if err := Down(); err != nil {
			return err
		}

		return Up()
	},
}

var MigrateCreateCommand = &cli.Command{
	Name:    "migrate:create",
	Aliases: []string{"mc"},
	Usage:   "Migrate create",
	Action: func(c *cli.Context) error {
		name := c.Args().Get(0)
		if name == "" {
			return fmt.Errorf("name is required")
		}
		return Create(name)
	},
}

var MigrateStatusCommand = &cli.Command{
	Name:  "migrate:status",
	Usage: "Run migrations",
	Action: func(c *cli.Context) error {
		return Status()
	},
}

var MigrateDownToCommand = &cli.Command{
	Name:  "migrate:down-to",
	Usage: "Run migrations",
	Action: func(c *cli.Context) error {
		if !helper.ConfirmProductionAction() {
			return nil
		}

		to := c.Args().Get(0)
		toInt, err := strconv.Atoi(to)
		if err != nil {
			return err
		}
		return DownTo(int64(toInt))
	},
}

var MigrateUpToCommand = &cli.Command{
	Name:  "migrate:up-to",
	Usage: "Run migrations",
	Action: func(c *cli.Context) error {
		if !helper.ConfirmProductionAction() {
			return nil
		}

		to := c.Args().Get(0)
		toInt, err := strconv.Atoi(to)
		if err != nil {
			return err
		}
		return UpTo(int64(toInt))
	},
}

func Up() error {
	m, err := newMigrator()
	if err != nil {
		return err
	}
	return goose.Up(m.db, m.dir)
}

func Create(name string) error {
	m, err := newMigrator()
	if err != nil {
		return err
	}
	return goose.Create(m.db, m.dir, name, m.migrationType)
}

func Down() error {
	m, err := newMigrator()
	if err != nil {
		return err
	}
	return goose.Reset(m.db, m.dir)
}

func Status() error {
	m, err := newMigrator()
	if err != nil {
		return err
	}
	return goose.Status(m.db, m.dir)
}

func DownTo(name int64) error {
	m, err := newMigrator()
	if err != nil {
		return err
	}
	return goose.DownTo(m.db, m.dir, name)
}

func UpTo(name int64) error {
	m, err := newMigrator()
	if err != nil {
		return err
	}
	return goose.UpTo(m.db, m.dir, name)
}

type migrator struct {
	dir           string
	dialect       string
	tableName     string
	migrationType string
	db            *sql.DB
}

func newMigrator() (*migrator, error) {
	db := config.RelationalDatabase()
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	m := &migrator{
		dir:           "internal/database/migrations",
		dialect:       os.Getenv("DB_DRIVER"),
		tableName:     "migrations",
		migrationType: "go",
		db:            sqlDB,
	}
	if err := m.init(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *migrator) init() error {
	goose.SetDialect(m.dialect)
	goose.SetTableName(m.tableName)
	if err := goose.SetDialect(m.dialect); err != nil {
		return err
	}
	if err := schema.SetDialect(m.dialect); err != nil {
		return err
	}
	return nil
}
