package container

import (
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/database"
	"go.uber.org/dig"
)

func BuildPackageContainer(container *dig.Container) error {
	// provide all package instance
	interfaces := []interface{}{
		database.InitDB,
	}

	for _, i := range interfaces {
		if err := container.Provide(i); err != nil {
			return err
		}
	}

	return nil
}
