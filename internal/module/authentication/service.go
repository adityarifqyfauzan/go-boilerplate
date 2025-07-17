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
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Login(ctx context.Context, request LoginRequest) *helper.ApiResponse
	Register(ctx context.Context, request RegisterRequest) *helper.ApiResponse
	RefreshToken(ctx context.Context, refreshToken string) *helper.ApiResponse
	Me(ctx context.Context) *helper.ApiResponse
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
		db:           db,
		localRepo:    localRepository,
		userRepo:     userRepository,
		userRoleRepo: userRoleRepository,
		roleRepo:     roleRepository,
		jwtService:   jwt.NewJWTService(),
	}
}

func (s *service) Login(ctx context.Context, request LoginRequest) *helper.ApiResponse {
	translate := translator.NewTranslator(ctx.Value(translator.LOCALIZER).(*i18n.Localizer))

	// Find user by email
	user, err := s.userRepo.FindOneBy(ctx, map[string]interface{}{"email": request.Email})
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.invalid_credentials", nil), nil)
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return helper.NewApiResponse(http.StatusBadRequest, translate.T("auth.invalid_credentials", nil), nil)
	}

	if user.UserStatusID == constant.USER_STATUS_INACTIVE_ID {
		return helper.NewApiResponse(http.StatusUnauthorized, translate.T("auth.user_inactive", nil), nil)
	}

	// get user role
	userRoles, err := s.userRoleRepo.FindBy(ctx, map[string]interface{}{"user_id": user.ID}, "", 0, 0)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.user_not_found", nil), nil)
	}

	roleIDs := make([]int, 0)
	for _, userRole := range userRoles {
		roleIDs = append(roleIDs, userRole.RoleID)
	}

	// get user role
	roles, err := s.roleRepo.FindBy(ctx, map[string]interface{}{"id": roleIDs}, "", 0, 0)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.user_not_found", nil), nil)
	}

	roleNames := make([]string, 0)
	for _, role := range roles {
		roleNames = append(roleNames, role.Slug)
	}

	// Generate JWT tokens
	token, err := s.jwtService.GenerateToken(jwt.Claims{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Name,
		Roles:    roleNames,
	})
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.failed_generate_tokens", nil), nil)
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.failed_generate_refresh_token", nil), nil)
	}

	return helper.NewApiResponse(http.StatusOK, translate.T("auth.login_successful", nil), LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		User: MeResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
			Roles: roles,
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

	tx := s.db.Begin()
	defer func() {
		tx.Rollback()
	}()

	// Create user
	user := &model.User{
		Name:         request.Name,
		Email:        request.Email,
		Password:     request.Password, // will hash in BeforeCreate hook (see User model)
		UserStatusID: constant.USER_STATUS_PENDING_ID,
	}

	createdUser, err := s.userRepo.Create(ctx, user, tx)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.failed_create_user", nil), nil)
	}

	// Create user role
	userRole := &model.UserRole{
		UserID: createdUser.ID,
		RoleID: constant.ROLE_USER_ID,
	}
	_, err = s.userRoleRepo.Create(ctx, userRole, tx)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.failed_create_user_role", nil), nil)
	}

	// Generate JWT tokens
	token, err := s.jwtService.GenerateToken(jwt.Claims{
		UserID:   createdUser.ID,
		Email:    createdUser.Email,
		Username: createdUser.Name,
		Roles:    []string{constant.ROLE_USER_SLUG},
	})
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.failed_generate_tokens", nil), nil)
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(createdUser.ID)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.failed_generate_refresh_token", nil), nil)
	}

	if err := tx.Commit().Error; err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("error.422", nil), nil)
	}

	return helper.NewApiResponse(http.StatusCreated, translate.T("auth.registration_successful", nil), RegisterResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

func (s *service) RefreshToken(ctx context.Context, refreshToken string) *helper.ApiResponse {
	translate := translator.NewTranslator(ctx.Value(translator.LOCALIZER).(*i18n.Localizer))

	// validate refresh token and generate new access token
	_, err := s.jwtService.ValidateToken(refreshToken)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnauthorized, translate.T("auth.invalid_refresh_token", nil), nil)
	}

	userID, err := s.jwtService.ExtractUserID(refreshToken)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnauthorized, translate.T("auth.invalid_refresh_token", nil), nil)
	}

	user, err := s.userRepo.FindOneBy(ctx, map[string]interface{}{"id": userID})
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.invalid_credentials", nil), nil)
	}

	if user.UserStatusID == constant.USER_STATUS_INACTIVE_ID {
		return helper.NewApiResponse(http.StatusUnauthorized, translate.T("auth.user_inactive", nil), nil)
	}

	userRoles, err := s.userRoleRepo.FindBy(ctx, map[string]interface{}{"user_id": user.ID}, "", 0, 0)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.user_not_found", nil), nil)
	}

	roleIDs := make([]int, 0)
	for _, userRole := range userRoles {
		roleIDs = append(roleIDs, userRole.RoleID)
	}

	roles, err := s.roleRepo.FindBy(ctx, map[string]interface{}{"id": roleIDs}, "", 0, 0)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.user_not_found", nil), nil)
	}

	roleNames := make([]string, 0)
	for _, role := range roles {
		roleNames = append(roleNames, role.Slug)
	}

	// Generate JWT tokens
	token, err := s.jwtService.GenerateToken(jwt.Claims{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Name,
		Roles:    roleNames,
	})
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.failed_generate_tokens", nil), nil)
	}

	return helper.NewApiResponse(http.StatusOK, translate.T("auth.token_refreshed", nil), RefreshTokenResponse{
		Token: token,
	})
}

func (s *service) Me(ctx context.Context) *helper.ApiResponse {
	translate := translator.NewTranslator(ctx.Value(translator.LOCALIZER).(*i18n.Localizer))
	userID := ctx.Value("user_id").(int)
	user, err := s.userRepo.FindOneBy(ctx, map[string]interface{}{"id": userID})
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.user_not_found", nil), nil)
	}

	userRoles, err := s.userRoleRepo.FindBy(ctx, map[string]interface{}{"user_id": userID}, "", 0, 0)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.user_roles_not_found", nil), nil)
	}

	roleIDs := make([]int, 0)
	for _, userRole := range userRoles {
		roleIDs = append(roleIDs, userRole.RoleID)
	}

	roles, err := s.roleRepo.FindBy(ctx, map[string]interface{}{"id": roleIDs}, "", 0, 0)
	if err != nil {
		return helper.NewApiResponse(http.StatusUnprocessableEntity, translate.T("auth.roles_not_found", nil), nil)
	}

	return helper.NewApiResponse(http.StatusOK, translate.T("success", nil), MeResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Roles: roles,
	})
}
