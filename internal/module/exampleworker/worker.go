package exampleworker

import (
	"context"
	"encoding/json"
	"log"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

func ExampleWorker() func(ctx context.Context, ch *amqp.Channel, conf *config.Config) {
	return func(ctx context.Context, ch *amqp.Channel, conf *config.Config) {
		service := NewService(
			config.DB,
			NewLocalRepository(config.DB),
		)

		message, err := rabbitmq.Consume(ch, &rabbitmq.PublishOption{
			Topic: "example",
		})
		if err != nil {
			panic(err)
		}

		log.Println("example worker started")

		for msg := range message {
			tr := otel.Tracer("example-worker")
			ctx, span := tr.Start(ctx, "ExampleWorker")

			log.Println("example worker received message")

			var request ExampleRequest
			if err := json.Unmarshal(msg.Body, &request); err != nil {
				panic(err)
			}

			if err := service.Example(ctx, request.Name); err != nil {
				log.Printf("failed to example: %v", err)
			}

			span.End()

			log.Println("example worker processed message")
		}
	}
}
