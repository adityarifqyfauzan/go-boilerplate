package container

import (
	"github.com/adityarifqyfauzan/go-boilerplate/app/usecase"
	"go.uber.org/dig"
)

func BuildRUsecaseContainer(container *dig.Container) error {
	// provide all usecase interfaces
	interfaces := []interface{}{
		usecase.NewHelloWorldUsecase,
	}

	for _, i := range interfaces {
		if err := container.Provide(i); err != nil {
			return err
		}
	}

	return nil
}
