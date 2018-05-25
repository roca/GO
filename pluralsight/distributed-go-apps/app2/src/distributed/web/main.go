package main

import (
	"net/http"

	"github.com/GOCODE/pluralsight/distributed-go-apps/app2/src/distributed/web/controller"
)

func main() {
	controller.Initialize()

	http.ListenAndServe(":3000", nil)
}
