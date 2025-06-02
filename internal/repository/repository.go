package repository

import (
	"context"

	"gorm.io/gorm"
)

type RelationalRepository[T any] interface {
	FindOneBy(ctx context.Context, criteria map[string]interface{}) (*T, error)
	FindBy(ctx context.Context, criteria map[string]interface{}, orderBy string, page, size int) ([]*T, error)
	Create(ctx context.Context, m *T, tx *gorm.DB) (*T, error)
	Update(ctx context.Context, m *T, tx *gorm.DB) error
	Delete(ctx context.Context, m *T, tx *gorm.DB) error
}

type DocumentRepository[T any] interface {
	FindOneBy(ctx context.Context, filter interface{}) (*T, error)
	FindBy(ctx context.Context, filter interface{}, orderBy string, page, size int) ([]*T, error)
	Create(ctx context.Context, m *T) (*T, error)
	Update(ctx context.Context, filter interface{}, update interface{}) error
	Delete(ctx context.Context, filter interface{}) error
}
