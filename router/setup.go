package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func setup(engine *gin.Engine) *gin.RouterGroup {
	// cek route
	health := engine.Group("")
	health.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"acknowledge": true,
		})
	})

	// set mode
	if os.Getenv("BUILD_ENV") != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// set endpoint prefix api
	route := engine.Group("api/")

	// register middleware
	registerMiddleware(route)

	return route
}
