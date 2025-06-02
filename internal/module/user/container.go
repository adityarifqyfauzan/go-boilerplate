package user

import "go.uber.org/dig"

func InitContainer(container *dig.Container) {
	if err := container.Provide(NewRepository); err != nil {
		panic(err)
	}

	if err := container.Provide(NewService); err != nil {
		panic(err)
	}
}
