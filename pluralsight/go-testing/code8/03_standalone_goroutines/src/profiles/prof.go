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

	f, err := os.Create("goroutines.prof")

	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < 5; i++ {
		http.Get("http://localhost:3000/api/purchaseOrders/1")
	}
	pprof.Lookup("goroutine").WriteTo(f, 1)
	f.Close()
}
