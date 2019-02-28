package main

import (
	"log"
	"net/http"
)

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := NewRouter()

	router.Handle("channel add", addChannel)

	http.Handle("/", router)
	http.ListenAndServe(":3001", nil)
}

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

}
