package main

import (
	"log"
	"math/rand"
	"net/http"
	"tracing/setup"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	tp, err := setup.NewTracerProvider()
	if err != nil {
		log.Fatal(err)
	}

	handleRoot := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if i := rand.Intn(2); i == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		w.Write([]byte("Hello, #PITO!"))
	})

	http.Handle("/", otelhttp.NewHandler(handleRoot, "server_root", otelhttp.WithTracerProvider(tp)))

	log.Fatal(http.ListenAndServe("localhost:8083", nil))

}
