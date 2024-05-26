package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ToDo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// No adapter
	todo := getRemoteData()

	fmt.Println("TODO without adapter:\t", todo.ID, todo.Title)
}

func getRemoteData() *ToDo {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var todo ToDo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Fatalln(err)
	}

	return &todo
}
