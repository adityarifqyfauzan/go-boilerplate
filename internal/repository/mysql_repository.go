package repository

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

type mysqlRepository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) RelationalRepository[T] {
	return &mysqlRepository[T]{db: db}
}

func (r *mysqlRepository[T]) FindOneBy(ctx context.Context, criteria map[string]interface{}) (*T, error) {
	tr := otel.Tracer("find-one-by-repository")
	spanName := fmt.Sprintf("FindOneByMainRepository<%T>", *new(T))
	ctx, span := tr.Start(ctx, spanName)
	defer span.End()

	var err error
	defer func(err error) {
		if err != nil {
			span.RecordError(err)
		}
	}(err)

	var entity T
	err = r.db.WithContext(ctx).Where(criteria).First(&entity).Error
	if err != nil {
		return nil, fmt.Errorf("find one by failed: %w", err)
	}
	return &entity, nil
}

func (r *mysqlRepository[T]) FindBy(ctx context.Context, criteria map[string]interface{}, orderBy string, page, size int) ([]*T, error) {
	tr := otel.Tracer("find-by-repository")
	spanName := fmt.Sprintf("FindByMainRepository<%T>", *new(T))
	ctx, span := tr.Start(ctx, spanName)
	defer span.End()

	var err error
	defer func(err error) {
		if err != nil {
			span.RecordError(err)
		}
	}(err)

	var entities []*T
	query := r.db.WithContext(ctx).Where(criteria)

	if orderBy != "" {
		query = query.Order(orderBy)
	}

	query = query.Offset((page - 1) * size)
	if size != 0 {
		query = query.Limit(size)
	}

	err = query.Find(&entities).Error
	if err != nil {
		return nil, fmt.Errorf("find by failed: %w", err)
	}

	return entities, nil
}

func (r *mysqlRepository[T]) Create(ctx context.Context, m *T, tx *gorm.DB) (*T, error) {
	tr := otel.Tracer("create-repository")
	spanName := fmt.Sprintf("CreateMainRepository<%T>", *new(T))
	ctx, span := tr.Start(ctx, spanName)
	defer span.End()

	var err error
	defer func(err error) {
		if err != nil {
			span.RecordError(err)
		}
	}(err)

	err = tx.WithContext(ctx).Create(m).Error
	if err != nil {
		return nil, fmt.Errorf("create failed: %w", err)
	}
	return m, nil
}

func (r *mysqlRepository[T]) Update(ctx context.Context, m *T, tx *gorm.DB) error {
	tr := otel.Tracer("update-repository")
	spanName := fmt.Sprintf("UpdateMainRepository<%T>", *new(T))
	ctx, span := tr.Start(ctx, spanName)
	defer span.End()

	var err error
	defer func(err error) {
		if err != nil {
			span.RecordError(err)
		}
	}(err)

	err = tx.WithContext(ctx).Save(m).Error
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}
	return nil
}

func (r *mysqlRepository[T]) Delete(ctx context.Context, m *T, tx *gorm.DB) error {
	tr := otel.Tracer("delete-repository")
	spanName := fmt.Sprintf("DeleteMainRepository<%T>", *new(T))
	ctx, span := tr.Start(ctx, spanName)
	defer span.End()

	var err error
	defer func(err error) {
		if err != nil {
			span.RecordError(err)
		}
	}(err)

	err = tx.WithContext(ctx).Delete(m).Error
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	return nil
}
