package authentication

import (
	"context"
	"net/http"

	"github.com/adityarifqyfauzan/go-boilerplate/internal/helper"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/helper/constant"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/model"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/repository"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/jwt"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/translator"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Login(ctx context.Context, request LoginRequest) *helper.ApiResponse
	Register(ctx context.Context, request RegisterRequest) *helper.ApiResponse
	RefreshToken(ctx context.Context, refreshToken string) *helper.ApiResponse
}

type service struct {
	localRepo    LocalRepository
	userRepo     repository.RelationalRepository[model.User]
	userRoleRepo repository.RelationalRepository[model.UserRole]
	roleRepo     repository.RelationalRepository[model.Role]
	db           *gorm.DB
	jwtService   *jwt.JWTService
}

func NewService(
	db *gorm.DB,
	localRepository LocalRepository,
	userRepository repository.RelationalRepository[model.User],
	userRoleRepository repository.RelationalRepository[model.UserRole],
	roleRepository repository.RelationalRepository[model.Role],
) Service {
	return &service{
		db:         db,
		localRepo:  localRepository,
		userRepo:   userRepository,
		jwtService: jwt.NewJWTService(),
	}
}

func (s *service) Login(ctx context.Context, request LoginRequest) *helper.ApiResponse {
	translate := translator.NewTranslator(ctx.Value(translator.LOCALIZER).(*i18n.Localizer))

	// Find user by email
	user, err := s.userRepo.FindOneBy(ctx, map[string]interface{}{"email": request.Email})
	if err != nil {
		return helper.NewApiResponse(http.StatusUnauthorized, translate.T("auth.invalid_credentials", nil), nil)
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return helper.NewApiResponse(http.StatusUnauthorized, translate.T("auth.invalid_credentials", nil), nil)
	}

	if user.UserStatusID == constant.USER_STATUS_INACTIVE_ID {
		return helper.NewApiResponse(http.StatusUnauthorized, translate.T("auth.user_inactive", nil), nil)
	}

	// get user role
	_, err = s.userRoleRepo.FindOneBy(ctx, map[string]interface{}{"user_id": user.ID})
	if err != nil {
		return helper.NewApiResponse(http.StatusUnauthorized, translate.T("auth.user_not_found", nil), nil)
	}

	// get user role
	role, err := s.roleRepo.FindOneBy(ctx, map[string]interface{}{"id": user.UserStatusID})
	if err != nil {
		return helper.NewApiResponse(http.StatusUnauthorized, translate.T("auth.user_not_found", nil), nil)
	}

	// Generate JWT tokens
	accessToken, err := s.jwtService.GenerateToken(
		user.ID,
		user.Email,
		user.Name,
		role.Slug,
	)
	if err != nil {
		return helper.NewApiResponse(http.StatusInternalServerError, translate.T("auth.failed_generate_tokens", nil), nil)
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		return helper.NewApiResponse(http.StatusInternalServerError, translate.T("auth.failed_generate_refresh_token", nil), nil)
	}

	return helper.NewApiResponse(http.StatusOK, translate.T("auth.login_successful", nil), gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		},
	})
}

func (s *service) Register(ctx context.Context, request RegisterRequest) *helper.ApiResponse {
	translate := translator.NewTranslator(ctx.Value(translator.LOCALIZER).(*i18n.Localizer))

	// Check if user already exists
	existingUser, _ := s.userRepo.FindOneBy(ctx, map[string]interface{}{"email": request.Email})
	if existingUser != nil {
		return helper.NewApiResponse(http.StatusConflict, translate.T("auth.user_already_exists", nil), nil)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return helper.NewApiResponse(http.StatusInternalServerError, translate.T("auth.failed_hash_password", nil), nil)
	}

	// Create user
	user := &model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := s.userRepo.Create(ctx, user, s.db)
	if err != nil {
		return helper.NewApiResponse(http.StatusInternalServerError, translate.T("auth.failed_create_user", nil), nil)
	}

	// Generate JWT tokens
	accessToken, err := s.jwtService.GenerateToken(
		createdUser.ID,
		createdUser.Email,
		createdUser.Name,
		"user",
	)
	if err != nil {
		return helper.NewApiResponse(http.StatusInternalServerError, translate.T("auth.failed_generate_tokens", nil), nil)
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(createdUser.ID)
	if err != nil {
		return helper.NewApiResponse(http.StatusInternalServerError, translate.T("auth.failed_generate_refresh_token", nil), nil)
	}

	return helper.NewApiResponse(http.StatusCreated, translate.T("auth.registration_successful", nil), gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":    createdUser.ID,
			"email": createdUser.Email,
			"name":  createdUser.Name,
		},
	})
}

func (s *service) RefreshToken(ctx context.Context, refreshToken string) *helper.ApiResponse {
	translate := translator.NewTranslator(ctx.Value(translator.LOCALIZER).(*i18n.Localizer))

	// Validate refresh token and generate new access token
	newAccessToken, err := s.jwtService.RefreshToken(refreshToken)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnauthorized, translate.T("auth.invalid_refresh_token", nil), nil)
	}

	return helper.NewApiResponse(http.StatusOK, translate.T("auth.token_refreshed", nil), gin.H{
		"access_token": newAccessToken,
	})
}
