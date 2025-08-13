package exampleworker

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type Service interface {
	Example(ctx context.Context, name string) error
}

type service struct {
	localRepository LocalRepository
	db              *gorm.DB
}

func NewService(
	db *gorm.DB,
	localRepository LocalRepository,
) Service {
	return &service{
		db:              db,
		localRepository: localRepository,
	}
}

func (s *service) Example(ctx context.Context, name string) error {
	log.Println("Hello, ", name)
	return nil
}
