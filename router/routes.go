package router

import (
	"github.com/adityarifqyfauzan/go-boilerplate/app/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRoute(
	conf *viper.Viper,
	engine *gin.Engine,

	// register all handler interface here
	helloHandler handler.HelloWorldHandler,

) {
	route := setup(engine)

	// api v1
	v1 := route.Group("v1")
	v2 := route.Group("v2")
	_ = v2

	// hello world handler
	helloRoutes := v1.Group("hello") // define prefix hello for all helloworld routes
	{
		helloRoutes.GET(":name", helloHandler.Hello)
		helloRoutes.POST("exception", helloHandler.Exception)
	}

}
