package main

import (
	"github.com/adityarifqyfauzan/go-boilerplate/internal/bootstrap"
	_ "github.com/adityarifqyfauzan/go-boilerplate/internal/database/migrations"
)

func main() {
	bootstrap.Init()
}
