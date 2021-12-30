package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// MyStruct is an example structure for this program.

func read_json() {

	filename := "forum.json"

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
	}
	var posts []Post
	json.Unmarshal([]byte(file), &posts)
}
