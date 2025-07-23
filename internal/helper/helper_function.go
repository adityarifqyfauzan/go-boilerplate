package helper

import (
	"fmt"
	"os"
)

func ConfirmProductionAction() bool {
	// check APP_ENV, if prod or production, give warning and options to rollback
	if IsProduction() {
		fmt.Println("You are running production environment, are you sure? (y/N)")
		input := "N"
		fmt.Scanln(&input)
		if input != "y" && input != "Y" {
			return false
		}
	}

	return true
}

func IsProduction() bool {
	return os.Getenv("APP_ENV") == "prod" || os.Getenv("APP_ENV") == "production" || os.Getenv("APP_ENV") == "PROD" || os.Getenv("APP_ENV") == "PRODUCTION"
}
