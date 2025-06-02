package config

import (
	"context"
	"database/sql"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

var (
	once  sync.Once
	DB    *gorm.DB
	Mongo *mongo.Client
	SqlDB *sql.DB
)

func RelationalDatabase(envFiles ...string) *gorm.DB {
	requiredEnvVars := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_NAME",
		"DB_DRIVER",
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("missing env var: %s", envVar)
		}
	}

	switch os.Getenv("DB_DRIVER") {
	case "mysql":
		mySqlDriver(envFiles...)
	case "postgres":
		postgreSqlDriver(envFiles...)
	default:
		log.Fatalf("unknown database driver: %s", os.Getenv("DB_DRIVER"))
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error: failed to get SQL DB object from GORM: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(15 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	SqlDB = sqlDB

	log.Println("connected to database")

	return DB
}

func CloseDB() {
	if SqlDB != nil {
		if err := SqlDB.Close(); err != nil {
			log.Printf("failed to close database connection: %v", err)
		}

		log.Println("closed database connection")
	}

	if Mongo != nil {
		ctx := context.Background()
		if err := Mongo.Disconnect(ctx); err != nil {
			log.Printf("failed to close mongodb connection: %v", err)
		}

		log.Println("closed mongodb connection")
	}
}

func MongoDB(envFiles ...string) *mongo.Client {
	requiredEnvVars := []string{
		"MONGO_HOST",
		"MONGO_PORT",
		"MONGO_SECURITY",
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("missing env var: %s", envVar)
		}
	}

	mongoDB()

	return Mongo
}
