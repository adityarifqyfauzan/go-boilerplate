package main

import (
	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/container"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

func InitContainer() *dig.Container {
	c := dig.New()

	if err := c.Provide(gin.New); err != nil {
		panic(err)
	}

	if err := c.Provide(viper.New); err != nil {
		panic(err)
	}

	if err := c.Provide(config.New); err != nil {
		panic(err)
	}

	if err := container.BuildPackageContainer(c); err != nil {
		panic(err)
	}

	if err := container.BuildHandlerContainer(c); err != nil {
		panic(err)
	}

	if err := container.BuildRUsecaseContainer(c); err != nil {
		panic(err)
	}

	if err := container.BuildRepositoryContainer(c); err != nil {
		panic(err)
	}

	return c
}
