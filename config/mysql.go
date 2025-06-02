package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func mySqlDriver(envFiles ...string) {
	once.Do(func() {
		if len(envFiles) == 0 {
			err := godotenv.Load(envFiles...)
			if err != nil {
				log.Printf("failed to load env file: %v", err)
			}
			log.Println("loaded env file")
		}

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: 500 * time.Millisecond, // Log slow queries
				LogLevel:      logger.Warn,            // Warning level logging
				Colorful:      true,                   // Colorful log output
			},
		),
			PrepareStmt: true, // Prepared statement caching
		})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}
		DB = db

	})
}
