package main

import (
	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/container"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

func InitContainer() *dig.Container {
	c := dig.New()

	if err := c.Provide(grpc.NewServer); err != nil {
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

	if err := container.BuildRServiceContainer(c); err != nil {
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
