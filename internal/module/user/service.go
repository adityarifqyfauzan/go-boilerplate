// grpc services implementation

package user

import (
	"context"

	"github.com/adityarifqyfauzan/go-boilerplate/internal/model"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/repository"
	"gorm.io/gorm"
)

type Service interface {
	FindOneByID(ctx context.Context, id int) (*UserResponse, error)
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
}

type service struct {
	db       *gorm.DB
	userRepo repository.RelationalRepository[model.User]
}

func NewService(
	db *gorm.DB,
	userRepo repository.RelationalRepository[model.User],
) Service {
	return &service{
		db:       db,
		userRepo: userRepo,
	}
}

func (s *service) FindOneByID(ctx context.Context, id int) (*UserResponse, error) {
	user, err := s.userRepo.FindOneBy(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		User: *user,
	}, nil
}

func (s *service) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	user, err := s.userRepo.FindOneBy(ctx, map[string]interface{}{"email": req.Email})
	if err != nil {
		return nil, err
	}
	_ = user
	return nil, nil
}
