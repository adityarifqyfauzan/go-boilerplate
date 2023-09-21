package container

import (
	"github.com/adityarifqyfauzan/go-boilerplate/app/handler"
	"go.uber.org/dig"
)

func BuildHandlerContainer(container *dig.Container) error {
	// provide all handler interfaces
	interfaces := []interface{}{
		handler.NewHelloWorldHandler,
	}

	for _, i := range interfaces {
		if err := container.Provide(i); err != nil {
			return err
		}
	}

	return nil
}
