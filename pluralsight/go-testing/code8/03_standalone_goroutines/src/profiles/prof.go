package main

import (
	"log"
	"net/http"
	"os"
	"runtime/pprof"

	"github.com/GOCODE/pluralsight/go-testing/code2/src/poms"
	"github.com/GOCODE/pluralsight/go-testing/code2/src/poms/ctrl"
)

func main() {
	ctrl.Setup()

	go http.ListenAndServe(":3000", new(poms.GZipServer))

	f, err := os.Create("cpu.prof")

	if err != nil {
		log.Fatal(err.Error())
	}

	pprof.StartCPUProfile(f)
	for i := 0; i < 1000; i++ {
		http.Get("http://localhost:3000/api/purchaseOrders/1")
	}
	pprof.StopCPUProfile()
}
