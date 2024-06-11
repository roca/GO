package main

import (
	"context"
	"log"
	"time"
	"tracing/setup"
)

func main() {
	tp, err := setup.NewTracerProvider("example02")

	if err != nil {
		log.Fatal(err)
	}

	tracer := tp.Tracer("tracer.go")

	ctx, firstSpan := tracer.Start(context.Background(), "first")
	time.Sleep(time.Millisecond * 150)

	_, secondSpan := tracer.Start(ctx, "second")
	time.Sleep(time.Millisecond * 200)

	secondSpan.End()
	firstSpan.End()
}
