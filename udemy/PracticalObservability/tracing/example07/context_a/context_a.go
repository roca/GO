package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"tracing/setup"

	"go.opentelemetry.io/otel/trace"
)

func main() {
	tp, err := setup.NewTracerProvider()
	if err != nil {
		log.Fatal(err)
	}
	tracer := tp.Tracer("context_a.go")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, serverSpan := tracer.Start(
			context.Background(),
			"serve_a",
			trace.WithSpanKind(trace.SpanKindServer),
		)

		ctx, clientSpan := tracer.Start(
			ctx,
			"request_b",
			trace.WithSpanKind(trace.SpanKindClient),
		)

		res, err := http.Get("http://localhost:8083")
		if err != nil {
				// TODO: handle error
		} else {
			w.WriteHeader(res.StatusCode)
			io.Copy(w, res.Body)
		}

		w.Write([]byte("Hello, #PITO!"))

		clientSpan.End()
		serverSpan.End()
	})

	log.Fatal(http.ListenAndServe("localhost:8084", nil))

}
