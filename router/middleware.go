package router

import (
	"github.com/adityarifqyfauzan/go-boilerplate/app/middleware"
	"github.com/gin-gonic/gin"
)

func registerMiddleware(route *gin.RouterGroup) {
	// define middleware here
	// default middleware
	route.Use(gin.Logger())
	route.Use(gin.ErrorLogger())
	route.Use(middleware.ExceptionMiddleware())

	// custom middleware

}
