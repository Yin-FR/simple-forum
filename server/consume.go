package main

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

var Conn, Err = amqp.Dial("amqp://guest:guest@localhost:5672/")
var Ch, Error = Conn.Channel()
var Q_post = queue_built("post", Ch, Error)

// var Q_comment = queue_built("comment", Ch, Error)

func queue_built(r string, ch *amqp.Channel, err error) amqp.Queue {

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,        // queue name
		r,             // routing key
		"logs_direct", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := Ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)

	failOnError(err, "Failed to register a consumer")
	go func() {
		for d := range msgs {
			post_temp := &Post{}
			json.Unmarshal(d.Body, &post_temp)
			Post_current = append(Post_current, *post_temp)
		}
	}()

	fmt.Printf("Queue Built\n")
	return q
}
