package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
	"tracing/setup"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	tp, err := setup.NewTracerProvider()
	if err != nil {
		log.Fatal(err)
	}

	handleRoot := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span := trace.SpanFromContext(r.Context())

		span.AddEvent(
			"start sleep",
			trace.WithAttributes(attribute.String("sleep.duration", "300ms")),
		)
		time.Sleep(time.Millisecond * 300)
		span.AddEvent("end sleep")

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
