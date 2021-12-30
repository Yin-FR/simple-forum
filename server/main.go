package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/rs/xid"
	"github.com/streadway/amqp"
)

var Post_current []Post

type Post struct {
	Postid      xid.ID    `json:"postId"`
	Author      string    `json:"author"`
	Content     string    `json:"content"`
	Title       string    `json:"title"`
	Comment_len int       `json:"commentNumber"`
	Comment     []Comment `json:"comment"`
}

type Comment struct {
	Post_id         xid.ID `json:"postId"`
	Author          string `json:"author"`
	Post_title      string `json:"postTitle"`
	Comment_content string `json:"commentContent`
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

	p.Postid = xid.New()
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
	setupCORS(&w, r)

	w.Header().Set("Content-Type", "application/json")

	if (*r).Method == "OPTIONS" {
		return
	}

	// if r.URL.Path != "/post" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	switch r.Method {

	case "POST":
		// add item to the queue
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		decoder := json.NewDecoder(r.Body)

		// 用于存放参数key=value数据
		var params map[string]string

		// 解析参数 存入map
		decoder.Decode(&params)

		fmt.Printf("Post from website! r.PostFrom = %v\n", params["author"])

		plaintext := Post{
			Title:   params["title"],
			Author:  params["author"],
			Content: params["content"]}

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
func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/post", hello_server)
	mux.HandleFunc("/comment", hello_server_comment)
	fmt.Printf("Starting server for testing HTTP POST...\n")

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8000", handler)

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
	// queue_consume("post")
}
