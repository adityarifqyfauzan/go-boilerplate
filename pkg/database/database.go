package database

import (
	"fmt"
	"sync"

	"github.com/adityarifqyfauzan/go-boilerplate/app/domain/model"
	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// instance the database connection with singleton
func InitDB(conf *config.AppConfig) *gorm.DB {
	once.Do(func() {
		var err error

		// host=%s port=%s dbname=%s user=%s password=%s sslmode=%s
		dsn := fmt.Sprintf(conf.GetString("database.dsn"),
			conf.GetString("database.host"),
			conf.GetString("database.port"),
			conf.GetString("database.name"),
			conf.GetString("database.user"),
			conf.GetString("database.pass"),
			conf.GetString("database.sslmode"),
		)
		db, err = gorm.Open(postgres.Open(dsn))
		if err != nil {
			logger.Console.Error(err)
		}
	})

	autoMigrate()

	return db
}

func autoMigrate() error {
	migrations := []error{
		db.AutoMigrate(&model.User{}),
	}

	for _, err := range migrations {
		if err != nil {
			return err
		}
	}

	return nil
}
