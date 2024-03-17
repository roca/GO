package main

import (
	"context"
	"fmt"
	"log"
	"metrics/setup"
	"runtime"

	"go.opentelemetry.io/otel/metric"
)

const Example = "counter_simple"

func main() {
	var ms runtime.MemStats

	mp, r, err := setup.NewMetricProvider(Example)
	ctx := context.Background()
	if err != nil {
		log.Fatal(err)
	}

	meter := mp.Meter(fmt.Sprintf("pito.local/examples/metrics/%s", Example))

	counter, err := meter.Int64ObservableCounter(
		"process.runtime.go.gc.count",
		metric.WithDescription("The number of garbage collection cycles performed by the Go runtime."),
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = meter.RegisterCallback(func(ctx context.Context, o metric.Observer) error {
		log.Println("Metrics were collected")

		runtime.ReadMemStats(&ms)

		o.ObserveInt64(counter, int64(ms.NumGC))

		return nil
	}, counter)
	if err != nil {
		log.Fatal(err)
	}

	if err := r.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
