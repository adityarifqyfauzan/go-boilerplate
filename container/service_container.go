package container

import (
	"github.com/adityarifqyfauzan/go-boilerplate/app/service"
	"go.uber.org/dig"
)

func BuildRServiceContainer(container *dig.Container) error {
	// provide all services interfaces
	interfaces := []interface{}{
		service.NewHelloWorldService,
	}

	for _, i := range interfaces {
		if err := container.Provide(i); err != nil {
			return err
		}
	}

	return nil
}
