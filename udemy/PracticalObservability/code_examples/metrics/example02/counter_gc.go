package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"metrics/setup"
	"runtime"
	"strconv"
	"time"

	"go.opentelemetry.io/otel/metric"
)

// while true; do ./example02 -limit 30m ; done

const Example = "counter_gc"

var flimit = flag.String("limit", "1m", "How long to run the program for")

func main() {
	flag.Parse()
	dur , err := time.ParseDuration(*flimit)
	if err != nil {
		log.Fatal(err)
	}

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

	ti := time.NewTicker(1 * time.Second)
	i := 0

	for {
		next := <-ti.C

		if i >= int(dur.Seconds()) {
			break
		}
		i++

		runtime.GC()

		fmt.Println(strconv.Itoa(i) + ": Time is : " + next.String())
	}

	if err := r.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
