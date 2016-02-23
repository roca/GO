package main

import (
	"fmt"

	"github.com/GOCODE/pluralsight/distributed-go-apps/app2/src/distributed/coordinator"
)

func main() {
	ql := coordinator.NewQueueListener()
	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}
