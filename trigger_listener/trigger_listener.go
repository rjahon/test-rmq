package main

import (
	"context"
	"log"
	"net/http"
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
		"trigger", // name
		"direct",  // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	helper.FailOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare( //
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	helper.FailOnError(err, "Failed to declare a queue")

	routing_key := "record_id"
	exchange := "trigger"
	log.Printf("Binding queue %s to exchange %s with routing key %s",
		q.Name, exchange, routing_key)
	err = ch.QueueBind(
		q.Name,      // queue name
		routing_key, // routing key
		exchange,    // exchange
		false,
		nil)
	helper.FailOnError(err, "Failed to bind a queue")

	// id := rand.Intn(5)
	msgs, err := ch.Consume( //
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	helper.FailOnError(err, "Failed to register a consumer")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var forever chan struct{}

	// var id int
	go func() {
		for d := range msgs {
			// id, err = strconv.Atoi(string(d.Body))
			id := string(d.Body)
			// helper.FailOnError(err, "Failed to get id from message")
			log.Printf(" [x] %s", id)
			PublishMsg(id, ch, ctx)

			err := ch.PublishWithContext(ctx,
				"logger", // exchange
				"debug",  // routing key
				false,    // mandatory
				false,    // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(id),
				})
			helper.FailOnError(err, "Failed to publish a message")

		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func PublishMsg(id string, ch *amqp.Channel, ctx context.Context) {

	body, statusCode, err := helper.GetPhone(id)
	helper.LogOnError(err, "Failed to get phone")

	switch statusCode {
	case http.StatusNotFound:
		{
			err = ch.PublishWithContext(ctx,
				"logger", // exchange
				"error",  // routing key
				false,    // mandatory
				false,    // immediate
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
			err := ch.PublishWithContext(ctx,
				"logger", // exchange
				"info",   // routing key
				false,    // mandatory
				false,    // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        body,
				})
			helper.FailOnError(err, "Failed to publish a message")
		}

	default:
		helper.LogOnError(nil, "Unexpected response")
	}
	log.Printf(" [x] Sent %s", string(body))
}
