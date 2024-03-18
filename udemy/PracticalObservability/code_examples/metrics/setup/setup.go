package setup

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/metric"
	sdk "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var (
	ErrFailMeterSetup = errors.New("failed to setup meter provider")
)

func NewMetricProvider(applicationName string) (metric.MeterProvider, sdk.Reader, error) {
	ctx := context.Background()
	hostname, _ := os.Hostname()

	res, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceName(applicationName),
		semconv.ServiceInstanceID(hostname),
	))
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %s", ErrFailMeterSetup, err)
	}

	ctx, cxl := context.WithTimeout(ctx, time.Second)
	defer cxl()

	//var c otlpmetrichttp.Compression = otlpmetrichttp.

	exp, err := otlpmetrichttp.New(
		ctx,
		otlpmetrichttp.WithInsecure(),
		otlpmetrichttp.WithEndpoint("prometheus:9090"),
		otlpmetrichttp.WithURLPath("/api/v1/otlp/v1/metrics"),
		otlpmetrichttp.WithCompression(otlpmetrichttp.GzipCompression),
		otlpmetrichttp.WithHeaders(map[string]string{
			"Content-Encoding": "gzip",
		}),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %s", ErrFailMeterSetup, err)
	}

	r := sdk.NewPeriodicReader(exp, sdk.WithInterval(time.Second*15))

	mp := sdk.NewMeterProvider(
		sdk.WithReader(r),
		sdk.WithResource(res),
	)

	return mp, r, nil
}
