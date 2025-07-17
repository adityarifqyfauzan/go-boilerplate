package authentication

import (
	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/helper/constant"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/model"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/repository"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.RouterGroup, config *config.Config) {

	service := NewService(
		config.DB,
		NewLocalRepository(config.DB),
		repository.NewRepository[model.User](config.DB),
		repository.NewRepository[model.UserRole](config.DB),
		repository.NewRepository[model.Role](config.DB),
	)

	handler := NewHandler(service)

	authenticationRoute := route.Group("authentication")
	authenticationRoute.POST("/login", handler.Login)
	authenticationRoute.POST("/register", handler.Register)
	authenticationRoute.POST("/forgot-password", handler.ForgotPassword)
	authenticationRoute.POST("/refresh-token", handler.RefreshToken)

	authenticationRoute.Use(middleware.AuthMiddleware())
	authenticationRoute.Use(middleware.RoleMiddleware(constant.ROLE_USER_SLUG))
	authenticationRoute.GET("/me", handler.Me)
}
