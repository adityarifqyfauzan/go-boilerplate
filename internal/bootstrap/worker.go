package bootstrap

import (
	"context"
	"log"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/worker"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/rabbitmq"
)

var ()

func Worker(ctx context.Context, conf *config.Config, workerNames ...string) {
	conn, err := rabbitmq.Connection()
	if err != nil {
		log.Fatalf("failed to connect to rabbitmq: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	// recover from panic
	defer func() {
		if r := recover(); r != nil {
			log.Printf("worker panic: %v", r)
		}
	}()

	workers := worker.Workers
	if len(workerNames) > 0 {
		filteredWorkers := make(map[string]worker.WorkerFunc)
		for _, workerName := range workerNames {
			if worker, exists := workers[workerName]; exists {
				filteredWorkers[workerName] = worker
			}
		}
		workers = filteredWorkers
	}

	for _, worker := range workers {
		go worker(ctx, ch, conf)
	}

	// block until context is cancelled
	log.Println("worker started")
	<-ctx.Done()
	log.Println("worker shutting down:", ctx.Err())
}
