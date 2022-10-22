package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rjahon/labs-rmq/trigger_listener/helper"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helper.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logger", // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	helper.FailOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := helper.ParseId(os.Args) // TODO: fix parseId function

	PublishMsg(id, ch, ctx)
}

func PublishMsg(id int, ch *amqp.Channel, ctx context.Context) {
	body, statusCode, err := helper.GetPhone(id)
	helper.LogOnError(err, "Failed to get phone")

	switch statusCode {
	case http.StatusNotFound:
		{
			err = ch.PublishWithContext(ctx,
				"logs_direct", // exchange
				"info",        // routing key
				false,         // mandatory
				false,         // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        body,
				})
			helper.FailOnError(err, "Failed to publish a message")
		}
	case http.StatusInternalServerError:
		{
			PublishMsg(id, ch, ctx)
		}
	case http.StatusFound:
		{
			err = ch.PublishWithContext(ctx,
				"logs_direct", // exchange
				"info",        // routing key
				false,         // mandatory
				false,         // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        body,
				})
			helper.FailOnError(err, "Failed to publish a message")
		}
	default:
		{
			helper.FailOnError(nil, "Unexpected response")
		}
	}

	log.Printf(" [x] Sent %s", string(body))
}
