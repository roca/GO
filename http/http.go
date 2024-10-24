package main

import (
	"log"
	"net/http"
	"os"
)

const resp = `<html>
    <head>
        <title>Simple Web App</title>
    </head>
    <body>
        <h1>Simple Web App</h1>
        <p>Hello World!</p>
    </body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(resp))
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
