package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"tracing/setup"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	tp, err := setup.NewTracerProvider("example08_a")
	if err != nil {
		log.Fatal(err)
	}
	tracer := tp.Tracer("baggage_a.go")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, serverSpan := tracer.Start(
			context.Background(),
			"serve_a",
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer serverSpan.End()

		b := baggage.Baggage{}

		tfSrc := "Unknown"
		if n := r.Header.Get("X-Traffic-Source"); n != "" {
			tfSrc = n
		}

		m,err := baggage.NewMember("http.traffic.source", tfSrc)
		if err != nil {
			serverSpan.AddEvent("failed to create baggage member: " + err.Error())
		}

		b,err = b.SetMember(m)
		if err != nil {
			serverSpan.AddEvent("failed to set baggage member: " + err.Error())
		}

		
		ctx = baggage.ContextWithBaggage(ctx, b)

		serverSpan.SetAttributes(attribute.String("http.traffic.source", tfSrc))

		res, err := otelhttp.Get(ctx, "http://localhost:8083")
		if err != nil {
			serverSpan.SetStatus(codes.Error, err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		} else {
			w.WriteHeader(res.StatusCode)
			io.Copy(w, res.Body)
		}

		w.Write([]byte("Hello, #PITO! server_a"))

		// req, err := http.NewRequest("GET", "http://localhost:8083", nil)
		// if err != nil {
		// 	serverSpan.SetStatus(codes.Error, err.Error())
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		// ctx, clientSpan := tracer.Start(
		// 	ctx,
		// 	"request_b",
		// 	trace.WithSpanKind(trace.SpanKindClient),
		// )

		// propagator := propagation.NewCompositeTextMapPropagator(
		// 	propagation.TraceContext{},
		// 	propagation.Baggage{},
		// )
		// propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))

		// res, err := http.DefaultClient.Do(req)
		// if err != nil {
		// 	clientSpan.SetStatus(codes.Error, err.Error())
		// } else {
		// 	if res.StatusCode > 499 {
		// 		clientSpan.SetStatus(codes.Error, "status code above 499")
		// 	}
		// 	w.WriteHeader(res.StatusCode)
		// 	io.Copy(w, res.Body)
		// }

		// w.Write([]byte("Hello, #PITO! server_a"))

		// clientSpan.End()

	})

	log.Fatal(http.ListenAndServe("localhost:8084", nil))

}
