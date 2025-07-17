package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type mysqlRepository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) RelationalRepository[T] {
	return &mysqlRepository[T]{db: db}
}

func (r *mysqlRepository[T]) FindOneBy(ctx context.Context, criteria map[string]interface{}) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).Where(criteria).First(&entity).Error; err != nil {
		return nil, fmt.Errorf("find one by failed: %w", err)
	}
	return &entity, nil
}

func (r *mysqlRepository[T]) FindBy(ctx context.Context, criteria map[string]interface{}, orderBy string, page, size int) ([]*T, error) {
	var entities []*T
	query := r.db.WithContext(ctx).Where(criteria)

	if orderBy != "" {
		query = query.Order(orderBy)
	}

	query = query.Offset((page - 1) * size)
	if size != 0 {
		query = query.Limit(size)
	}

	if err := query.Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("find by failed: %w", err)
	}

	return entities, nil
}

func (r *mysqlRepository[T]) Create(ctx context.Context, m *T, tx *gorm.DB) (*T, error) {
	if err := tx.WithContext(ctx).Create(m).Error; err != nil {
		return nil, fmt.Errorf("create failed: %w", err)
	}
	return m, nil
}

func (r *mysqlRepository[T]) Update(ctx context.Context, m *T, tx *gorm.DB) error {
	if err := tx.WithContext(ctx).Save(m).Error; err != nil {
		return fmt.Errorf("update failed: %w", err)
	}
	return nil
}

func (r *mysqlRepository[T]) Delete(ctx context.Context, m *T, tx *gorm.DB) error {
	if err := tx.WithContext(ctx).Delete(m).Error; err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	return nil
}
