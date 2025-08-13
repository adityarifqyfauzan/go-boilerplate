package exampleworker

import (
	"gorm.io/gorm"
)

type LocalRepository interface {
}

type localRepository struct {
	db *gorm.DB
}

func NewLocalRepository(
	db *gorm.DB,
) LocalRepository {
	return &localRepository{
		db: db,
	}
}
