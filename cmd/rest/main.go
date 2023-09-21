package main

import (
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/logger"
	"github.com/adityarifqyfauzan/go-boilerplate/router"
)

func main() {

	// init IoC
	c := InitContainer()

	// init route
	if err := c.Invoke(router.InitRoute); err != nil {
		panic(err)
	}

	// invoke server
	if err := c.Provide(NewServer); err != nil {
		panic(err)
	}

	// run server
	if err := c.Invoke(func(s *Server) {
		s.Run()
	}); err != nil {
		logger.Console.Error(err)
	}

}
