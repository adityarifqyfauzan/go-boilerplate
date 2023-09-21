package main

import "github.com/adityarifqyfauzan/go-boilerplate/router"

func main() {
	c := InitContainer()

	if err := c.Provide(NewNetListener); err != nil {
		panic(err)
	}

	if err := c.Provide(NewGRPCServer); err != nil {
		panic(err)
	}

	if err := c.Invoke(router.InitGRPCRouter); err != nil {
		panic(err)
	}

	if err := c.Invoke(func(s *GRPCServer) {
		s.Run()
	}); err != nil {
		panic(err)
	}
}
