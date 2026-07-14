package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/honeycombio/honeycomb-opentelemetry-go"
	"github.com/honeycombio/otel-config-go/otelconfig"
	"github.com/roca/celeritas"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	// enable multi-span attributes
	bsp := honeycomb.NewBaggageSpanProcessor()

	// use honeycomb distro to setup OpenTelemetry SDK
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry(
		otelconfig.WithSpanProcessor(bsp),
	)
	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}
	defer otelShutdown()

	c := initApplication()
	c.App.ListenAndServe()
}
