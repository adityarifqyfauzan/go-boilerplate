package user

import (
	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.Engine, config *config.Config) {
	handler := NewHandler()

	userRoute := route.Group("user")
	userRoute.GET("/me", handler.Me)
}
