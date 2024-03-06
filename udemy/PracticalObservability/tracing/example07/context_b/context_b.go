package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"tracing/setup"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	tp, err := setup.NewTracerProvider()
	if err != nil {
		log.Fatal(err)
	}
	tracer := tp.Tracer("context_b.go")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, span := tracer.Start(
			context.Background(),
			"serve_b",
			trace.WithSpanKind(trace.SpanKindServer),
		)

		if i := rand.Intn(2); i == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			span.SetStatus(codes.Error, "Internal Server Error: random number generated a failure")
			span.End()
			return
		}

		w.Write([]byte("Hello, #PITO!"))

		span.End()
	})

	log.Fatal(http.ListenAndServe("localhost:8083", nil))

}
