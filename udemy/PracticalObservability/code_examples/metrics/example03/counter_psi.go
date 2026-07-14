package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"metrics/setup"
	"os"
	"strconv"
	"strings"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const Example = "counter_gc"

var flimit = flag.String("limit", "1m", "How long to run the program for")

func main() {
	flag.Parse()
	dur, err := time.ParseDuration(*flimit)
	if err != nil {
		log.Fatal(err)
	}

	mp, r, err := setup.NewMetricProvider(Example)
	ctx := context.Background()
	if err != nil {
		log.Fatal(err)
	}

	meter := mp.Meter(fmt.Sprintf("pito.local/examples/metrics/%s", Example))

	counter, err := meter.Int64ObservableCounter(
		"system.psi.time",
		metric.WithDescription("How long the computer waited on a resource to complete a task"),
		metric.WithUnit("us"),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range []string{"cpu", "io", "memory"} {
		_, err = meter.RegisterCallback(PSI(r, counter), counter)
	}
	if err != nil {
		log.Fatal(err)
	}

	<-time.After(dur)

	if err := r.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

var ErrFailedToReadPSI string = "Failed to read PSI"

func PSI(resource string, counter metric.Int64ObservableCounter) func(context.Context, metric.Observer) error {
	noop := func(context.Context, metric.Observer) error { return nil }
	handle, err := os.Open("/proc/pressure/" + resource)

	if err != nil {
		fmt.Println(ErrFailedToReadPSI, err)
		return noop
	}

	return func(ctx context.Context, o metric.Observer) error {
		if _, err := handle.Seek(0, 0); err != nil {
			fmt.Printf("%w: %s", ErrFailedToReadPSI, err)
			return nil
		}

		sc := bufio.NewScanner(handle)
		sc.Split(bufio.ScanWords)
		delay := "unknown"

		for sc.Scan() {
			if sc.Text() == "some" || sc.Text() == "full" {
				delay = sc.Text()
			}

			if after, isPresent := strings.CutPrefix(sc.Text(), "total="); isPresent {
				total, err := strconv.ParseInt(after, 10, 64)
				if err != nil {
					delay = "unknown"
					return fmt.Errorf("%w: %s", ErrFailedToReadPSI, err)
				}

				// Recordthe metric
				o.ObserveInt64(counter, total, metric.WithAttributes(
					attribute.String("resource", resource),
					attribute.String("delay", delay),
				))

				delay = "unknown"
			}
		}

		return nil
	}
}
