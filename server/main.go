package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

var Post_current []Post

type Post struct {
	Postid      []uint8   `json:"postId"`
	Author      string    `json:"author"`
	Content     string    `json:"content"`
	Title       string    `json:"title"`
	Comment_len int       `json:"commentNumber"`
	Comment     []Comment `json:"comment"`
}

type Comment struct {
	Post_id         []uint8 `json:"postId"`
	Author          string  `json:"author"`
	Post_title      string  `json:"postTitle"`
	Comment_content string  `json:"commentContent`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func queue_publish(p Post, r string, ch *amqp.Channel, err error) []byte {
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

	h := sha256.New()
	p.Postid = h.Sum([]byte(fmt.Sprintf("%v", p.Title)))
	p.Comment_len = len(p.Comment)

	byteBody, err := json.MarshalIndent(p, "", "  ")
	err = ch.Publish(
		"logs_direct", // exchange
		r,             // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        byteBody,
		})
	failOnError(err, "Failed to publish a message")

	return byteBody
}

func hello_server(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/post" {
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

		plaintext := Post{
			Title:   r.FormValue("title"),
			Author:  r.FormValue("author"),
			Content: r.FormValue("content")}

		queue_publish(plaintext, "post", Ch, Error)

		fmt.Fprintf(w, "ack")

	case "GET":

		byteArray, err := json.MarshalIndent(Post_current, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, string(byteArray))
		// fmt.Fprintf(w, "got!!")

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func hello_server_comment(w http.ResponseWriter, r *http.Request) {

	// if r.URL.Path != "/comment" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	// switch r.Method {

	// case "POST":
	// 	// add item to the queue
	// 	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	// 	if err := r.ParseForm(); err != nil {
	// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
	// 		return
	// 	}

	// 	commenttext := Comment{ Post_title: r.FormValue("title"),
	// 							Username: r.FormValue("username"),
	// 							Comment_content: r.FormValue("content")}
	// 	queue_publish(commenttext, "comment", Ch, Error)

	// 	fmt.Fprintf(w, "%v", plaintext)

	// case "GET":
	// 	fmt.Fprintf(w, get_info_from_queue(Q_comment))
	// 	fmt.Fprintf(w, "got!!")
	// default:
	// 	fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	// }
}

func main() {

	http.HandleFunc("/post", hello_server)
	http.HandleFunc("/comment", hello_server_comment)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	http.ListenAndServe(":8080", nil)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
	// queue_consume("post")
}
