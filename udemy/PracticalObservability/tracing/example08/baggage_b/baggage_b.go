package main

import (
	"log"
	"math/rand"
	"net/http"
	"tracing/setup"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	tp, err := setup.NewTracerProvider()
	if err != nil {
		log.Fatal(err)
	}

	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)

		b := baggage.FromContext(r.Context())
		span := trace.SpanFromContext(r.Context())
		if b.Member("http.traffic.source").Value() != "" {
			span.SetAttributes(
				attribute.String(
					"traffic.source",
					b.Member("http.traffic.source").Value(),
				),
			)
		}

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
