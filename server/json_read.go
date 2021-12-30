package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// MyStruct is an example structure for this program.
// type Post struct {
// 	Postid      string    `json:"postId"`
// 	Author      string    `json:"author"`
// 	Content     string    `json:"content"`
// 	Title       string    `json:"title"`
// 	Comment_len int       `json:"commentNumber"`
// 	Comment     []Comment `json:"comment"`
// }

// type Comment struct {
// 	Postid         string `json:"postId"`
// 	Author         string `json:"author"`
// 	CommentContent string `json:"commentContent`
// }

func Read_json() []Post {

	filename := "forum.json"

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
	}
	var posts []Post
	json.Unmarshal([]byte(file), &posts)
	return posts
}
