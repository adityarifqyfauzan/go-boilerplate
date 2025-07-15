package config

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

type Config struct {
	DB    *gorm.DB
	Mongo *mongo.Client
}

func New() *Config {

	// Connect to database
	DB := RelationalDatabase()
	Mongo := MongoDB()

	return &Config{
		DB:    DB,
		Mongo: Mongo,
	}
}
