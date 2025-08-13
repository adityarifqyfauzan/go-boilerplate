package worker

import (
	"context"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/module/exampleworker"
	amqp "github.com/rabbitmq/amqp091-go"
)

type WorkerFunc func(ctx context.Context, ch *amqp.Channel, config *config.Config)

// workerName as key, workerFunc as value
var Workers map[string]WorkerFunc

// register all worker here 👇
func init() {
	Workers = make(map[string]WorkerFunc)

	Workers["example-worker"] = exampleworker.ExampleWorker()
}
