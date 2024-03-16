package main

import (
	"context"
	"fmt"
	"log"
	"metrics/setup"
)

const Example = "counter_simple"

func main() {
	mp, r, err := setup.NewMetricProvider(Example)
	ctx := context.Background()
	if err != nil {
		log.Fatal(err)
	}

	meter := mp.Meter(fmt.Sprintf("pito.local/examples/metrics/%s", Example))

	counter , err := meter.Int64Counter("coffees")
	if err != nil {
		log.Fatal(err)
	}

	counter.Add(ctx, 1)
	counter.Add(ctx, 1)
	counter.Add(ctx, 1)

	if err := r.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
