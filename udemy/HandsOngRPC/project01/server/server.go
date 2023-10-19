package main

import (
	"log"
	"net/http"

	pb "project01/hello"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	req := pb.Hello{
		Name: name,
	}
	w.Write([]byte(req.Name))
}

func main() {
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}
