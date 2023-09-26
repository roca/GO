package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/honeycombio/honeycomb-opentelemetry-go"
    "github.com/honeycombio/otel-config-go/otelconfig"

    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Implement an HTTP Handler func to be instrumented
func httpHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World")
}

// Wrap the HTTP handler func with OTel HTTP instrumentation
func wrapHandler() {
    handler := http.HandlerFunc(httpHandler)
    wrappedHandler := otelhttp.NewHandler(handler, "hello")
    http.Handle("/hello", wrappedHandler)
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

    // Initialize HTTP handler instrumentation
    wrapHandler()
    log.Fatal(http.ListenAndServe(":3030", nil))
}