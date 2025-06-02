package container

import "go.uber.org/dig"

func BuildContainer(container *dig.Container) {
	configContainer(container)
	repositoryContainer(container)
	moduleContainer(container)
}
