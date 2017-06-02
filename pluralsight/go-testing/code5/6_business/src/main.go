package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GOCODE/pluralsight/go-testing/code2/src/poms"
	"github.com/GOCODE/pluralsight/go-testing/code2/src/poms/ctrl"
)

func main() {

	ctrl.Setup()

	go http.ListenAndServe(":3000", new(poms.GZipServer))

	log.Println("Server started, press <ENTER> to exit")
	fmt.Scanln()
}
