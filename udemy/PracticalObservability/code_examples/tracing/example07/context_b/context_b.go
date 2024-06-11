package main

import (
	"log"
	"math/rand"
	"net/http"
	"tracing/setup"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	tp, err := setup.NewTracerProvider("example07_b")
	if err != nil {
		log.Fatal(err)
	}

	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)

		if i := rand.Intn(2); i == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		w.Write([]byte("Hello, #PITO! server_b"))
	})

	http.Handle("/", otelhttp.NewHandler(hf, "serve_b", otelhttp.WithTracerProvider(tp)))

	log.Fatal(http.ListenAndServe("localhost:8083", nil))

}
