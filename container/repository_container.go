package container

import "go.uber.org/dig"

func BuildRepositoryContainer(container *dig.Container) error {
	// provide all repository interfaces
	interfaces := []interface{}{}

	for _, i := range interfaces {
		if err := container.Provide(i); err != nil {
			return err
		}
	}

	return nil
}
