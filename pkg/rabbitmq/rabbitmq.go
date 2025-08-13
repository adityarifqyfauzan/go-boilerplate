package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connection() (*amqp.Connection, error) {
	host := os.Getenv("RABBITMQ_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("RABBITMQ_PORT")
	if port == "" {
		port = "5672"
	}
	user := os.Getenv("RABBITMQ_USER")
	if user == "" {
		user = "guest"
	}
	pass := os.Getenv("RABBITMQ_PASS")
	if pass == "" {
		pass = "guest"
	}
	vhost := os.Getenv("RABBITMQ_VHOST")
	if vhost == "/" || vhost == "" {
		vhost = ""
	}

	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", user, pass, host, port, vhost)
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	log.Println("connected to rabbitmq")

	return conn, nil
}

type PublishOption struct {
	Topic            string
	Publishing       amqp.Publishing
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	Args             amqp.Table
}

func declareQueue(ch *amqp.Channel, opt *PublishOption) (amqp.Queue, error) {
	// declaring queue, will be created if not exists
	q, err := ch.QueueDeclare(
		opt.Topic,
		opt.Durable,
		opt.DeleteWhenUnused,
		opt.Exclusive,
		opt.NoWait,
		opt.Args,
	)
	if err != nil {
		return amqp.Queue{}, err
	}

	return q, nil
}

func PublishWithContext(ctx context.Context, ch *amqp.Channel, opt *PublishOption) error {
	q, err := declareQueue(ch, opt)
	if err != nil {
		return err
	}

	if err := ch.PublishWithContext(ctx, "", q.Name, false, false, opt.Publishing); err != nil {
		return err
	}

	return nil
}

func Consume(ch *amqp.Channel, opt *PublishOption) (<-chan amqp.Delivery, error) {
	q, err := declareQueue(ch, opt)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
