package container

import (
	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"go.uber.org/dig"
)

func configContainer(container *dig.Container) {
	if err := container.Provide(config.RelationalDatabase); err != nil {
		panic(err)
	}

	if err := container.Provide(config.New); err != nil {
		panic(err)
	}
}
