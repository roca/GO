package main

import (
	"fmt"

	"github.com/GOCODE/pluralsight/distributed-go-apps/app2/src/distributed/coordinator"
)

var dc *coordinator.DatabaseConsumer
var wc *coordinator.WebappConsumer

func main() {
	ea := coordinator.NewEventAggregator()
	dc = coordinator.NewDatabaseComsumer(ea)
	wc = coordinator.NewWebappComsumer(ea)
	ql := coordinator.NewQueueListener(ea)
	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}
