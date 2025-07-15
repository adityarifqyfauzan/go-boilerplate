package routes

import (
	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/module/authentication"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine, config *config.Config) {

	v1 := engine.Group("/v1")

	// register all module routes here
	authentication.InitRoute(v1, config)
}
