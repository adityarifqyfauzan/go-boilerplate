package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func mongoDB() {
	once.Do(func() {
		dsn := fmt.Sprintf("mongodb://%s:%s",
			os.Getenv("MONGO_HOST"),
			os.Getenv("MONGO_PORT"),
		)

		clientOption := options.Client().ApplyURI(dsn)
		if os.Getenv("MONGO_SECURITY") == "true" {
			clientOption = clientOption.SetAuth(options.Credential{
				Username: os.Getenv("MONGO_USER"),
				Password: os.Getenv("MONGO_PASS"),
			})
		}

		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		Mongo, err = mongo.Connect(clientOption)
		if err != nil {
			log.Printf("failed to connect to database: %v", err)
			return
		}

		err = Mongo.Ping(ctx, nil)
		if err != nil {
			log.Printf("failed to ping mongodb: %v", err)
			return
		}

		log.Println("connected to mongodb")
	})
}
