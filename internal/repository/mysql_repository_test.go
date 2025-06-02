package repository

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/model"
	"github.com/joho/godotenv"
)

func TestFindOneBy(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load env file: %v", err)
	}

	db := config.RelationalDatabase()

	repo := NewRepository[model.User](db)
	tx := db.Begin()
	defer tx.Rollback()

	repo.Create(context.Background(), &model.User{
		Email:    "test2@test.com",
		Password: "test",
		Name:     "test",
	}, tx)

	user, err := repo.FindOneBy(context.Background(), map[string]interface{}{"email": "test2@test.com"})
	if err != nil {
		t.Fatalf("error finding user: %v", err)
	}

	t.Logf("user: %v", user)

	if user.Email != "test2@test.com" {
		t.Errorf("expected email to be %s, got %s", "test2@test.com", user.Email)
	}

	tx.Commit()
	fmt.Println(user.ID)
}
