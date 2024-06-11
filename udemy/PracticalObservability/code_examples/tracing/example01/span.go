package main

import (
	"context"
	"log"
	"time"
	"tracing/setup"
)

func main() {
	tp, err := setup.NewTracerProvider("example01")

	if err != nil {
		log.Fatal(err)
	}

	tracer := tp.Tracer("span.go")

	_, span := tracer.Start(context.Background(), "span")
	time.Sleep(time.Millisecond * 300)
	span.End()
}
