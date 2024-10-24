package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.DialConfig("amqp://challenge:k64sMKmWEyg85VZs@rabbitmq-1.8696e293-bmoyles0117.node.tutum.io:5672", amqp.Config{Vhost: "Default"})
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// err = ch.ExchangeDeclare(
	// 	"amq",   // name
	// 	"topic", // type
	// 	false,   // durable
	// 	false,   // auto-deleted
	// 	false,   // internal
	// 	false,   // no-wait
	// 	nil,     // arguments
	// )
	// failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"roca", // name
		false,  // durable
		false,  // delete when usused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,      // queue name
		"formula.*", // routing key
		"amq.topic", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for formulas. To exit press CTRL+C")
	<-forever
}
