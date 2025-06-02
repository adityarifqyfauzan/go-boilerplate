package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type mongoRepository[T any] struct {
	collection *mongo.Collection
}

// Constructor
func NewMongoRepository[T any](collection *mongo.Collection) DocumentRepository[T] {
	return &mongoRepository[T]{collection: collection}
}

func (r *mongoRepository[T]) FindOneBy(ctx context.Context, filter interface{}) (*T, error) {
	var entity T
	err := r.collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("find one by failed: %w", err)
	}
	return &entity, nil
}

func (r *mongoRepository[T]) FindBy(ctx context.Context, filter interface{}, orderBy string, page, size int) ([]*T, error) {
	var entities []*T

	findOptions := options.Find().
		SetSkip(int64((page - 1) * size)).
		SetLimit(int64(size))

	if orderBy != "" {
		findOptions.SetSort(bson.M{orderBy: 1}) // 1 = ascending, -1 = descending
	}

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("find by failed: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var entity T
		if err := cursor.Decode(&entity); err != nil {
			return nil, fmt.Errorf("decode failed: %w", err)
		}
		entities = append(entities, &entity)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return entities, nil
}

func (r *mongoRepository[T]) Create(ctx context.Context, m *T) (*T, error) {
	_, err := r.collection.InsertOne(ctx, m)
	if err != nil {
		return nil, fmt.Errorf("create failed: %w", err)
	}
	return m, nil
}

func (r *mongoRepository[T]) Update(ctx context.Context, filter interface{}, update interface{}) error {
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no document matched")
	}
	return nil
}

func (r *mongoRepository[T]) Delete(ctx context.Context, filter interface{}) error {
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document deleted")
	}
	return nil
}
