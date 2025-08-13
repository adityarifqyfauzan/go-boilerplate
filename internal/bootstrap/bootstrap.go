package bootstrap

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/command"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/opentelemetry"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/translator"
	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load env file: %v", err)
	}

	conf := config.New()
	defer config.CloseDB()

	translator.Init("locales")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// setup OpenTelemetry
	otelShutdown, err := opentelemetry.SetupOTelSDK(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "worker":
			Worker(ctx, conf, os.Args[2:]...)
			return
		default:
			cmd := command.App
			err := cmd.Run(os.Args)
			if err != nil {
				log.Default().Fatalf("failed to run app: %v", err)
			}
			return
		}
	}

	// default behavior is to run rest-api
	RestAPI(ctx, conf)
}
