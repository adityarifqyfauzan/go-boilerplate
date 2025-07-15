package authentication

import "go.uber.org/dig"

func InitContainer(container *dig.Container) {
	if err := container.Provide(NewLocalRepository); err != nil {
		panic(err)
	}

	if err := container.Provide(NewService); err != nil {
		panic(err)
	}
}
