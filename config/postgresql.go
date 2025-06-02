package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func postgreSqlDriver(envFiles ...string) {
	once.Do(func() {
		if len(envFiles) == 0 {
			err := godotenv.Load(envFiles...)
			if err != nil {
				log.Printf("failed to load env file: %v", err)
			}
			log.Println("loaded env file")
		}

		sslMode := os.Getenv("DB_SSLMODE")
		if sslMode == "" {
			sslMode = "disable"
		}

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
			sslMode,
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.New(
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
