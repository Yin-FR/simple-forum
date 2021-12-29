package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

type post struct {
	username string
	content  string
	title    string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func queue_publish(p post, r string, ch *amqp.Channel, err error) post {
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

	body := p
	err = ch.Publish(
		"logs_direct", // exchange
		r,             // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%v", body)),
		})
	failOnError(err, "Failed to publish a message")

	return body
}
func get_info_from_queue(q amqp.Queue) string {
	res := ""
	msgs, err := Ch.Consume(
		Q_post.Name, // queue
		"",          // consumer
		true,        // auto ack
		false,       // exclusive
		false,       // no local
		false,       // no wait
		nil,         // args
	)

	failOnError(err, "Failed to register a consumer")
	go func() {
		for d := range msgs {
			fmt.Printf(string(d.Body))
		}
	}()
	return res
}
func hello_server(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "POST":
		// add item to the queue
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		plaintext := post{title: r.FormValue("title"), username: r.FormValue("username"), content: r.FormValue("content")}
		queue_publish(plaintext, "post", Ch, Error)
		fmt.Fprintf(w, "%v", plaintext)

	case "GET":
		fmt.Fprintf(w, get_info_from_queue(Q_post))
		fmt.Fprintf(w, "got!!")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

	http.HandleFunc("/", hello_server)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	http.ListenAndServe(":8080", nil)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
	// queue_consume("post")
}
