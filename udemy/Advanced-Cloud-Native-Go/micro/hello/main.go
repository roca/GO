package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"hello/handler"
	"hello/subscriber"

	hello "hello/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.hello"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.foo.service.hello", service.Server(), new(subscriber.Hello))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
