package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	go client()
	go server()

	// to keep main go routine alive while the others are busily publishing and receiving messages
	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever


}

func client() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume (
		q.Name, // name of queue
		"", // consumer
		true, // autoAck
		false, // exclusive
		false, // nolocal
		false, // nowait
		nil, // header
	)

	failOnError(err, "Failed to register a consumer")

	for msg := range msgs {
		log.Printf("Received message with message: %s", msg.Body)
	}
}

func server() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte("Hello RabbitMQ"),
	}

	// publish with direct exchange type (default): guarantee that only one of client get the message and only one time
	ch.Publish("", q.Name, false, false, msg)
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err,"Failed to open a channel")

	q, err := ch.QueueDeclare (
		"hello", // name of queue
		false, // durable
		false, // autodelete
		false, // exclusive
		false, // nowait
		nil, // header
	) // direct exchange
	failOnError(err, "Failed to declare a queue")

	return conn, ch, &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}