package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestFindOneBy(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	repo := NewRepository[model.User](gormDB)

	query := "SELECT \\* FROM `users` WHERE `email` = \\? ORDER BY `users`.`id` LIMIT \\?"
	rows := sqlmock.NewRows([]string{"id", "email", "name", "password"}).
		AddRow(1, "test2@test.com", "test", "hashed_password")

	mock.ExpectQuery(query).
		WithArgs("test2@test.com", 1).
		WillReturnRows(rows)

	user, err := repo.FindOneBy(context.Background(), map[string]any{"email": "test2@test.com"})
	if err != nil {
		t.Fatalf("error finding user: %v", err)
	}

	if user.Email != "test2@test.com" {
		t.Errorf("expected email to be %s, got %s", "test2@test.com", user.Email)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindBy(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	repo := NewRepository[model.User](gormDB)

	query := "SELECT \\* FROM `users` WHERE `email` = \\? ORDER BY id DESC LIMIT \\? OFFSET \\?"
	rows := sqlmock.NewRows([]string{"id", "email", "name", "password"}).
		AddRow(1, "test2@test.com", "test", "hashed_password").
		AddRow(2, "test3@test.com", "test3", "hashed_password")

	mock.ExpectQuery(query).
		WithArgs("test2@test.com", 10, 10).
		WillReturnRows(rows)

	users, err := repo.FindBy(context.Background(), map[string]any{"email": "test2@test.com"}, "id DESC", 2, 10)
	if err != nil {
		t.Fatalf("error finding users: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	repo := NewRepository[model.User](gormDB)

	user := &model.User{
		Email:    "new@test.com",
		Password: "password",
		Name:     "new user",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	createdUser, err := repo.Create(context.Background(), user, gormDB)
	if err != nil {
		t.Fatalf("error creating user: %v", err)
	}

	if createdUser.ID != 1 {
		t.Errorf("expected user ID 1, got %d", createdUser.ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	repo := NewRepository[model.User](gormDB)

	user := &model.User{
		BaseModel: model.BaseModel{ID: 1},
		Email:     "updated@test.com",
		Password:  "password",
		Name:      "updated user",
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Update(context.Background(), user, gormDB)
	if err != nil {
		t.Fatalf("error updating user: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	repo := NewRepository[model.User](gormDB)

	user := &model.User{
		BaseModel: model.BaseModel{ID: 1},
	}

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `users`").
		WithArgs(user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Delete(context.Background(), user, gormDB)
	if err != nil {
		t.Fatalf("error deleting user: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
