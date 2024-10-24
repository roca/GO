package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GOCODE/pluralsight/go-testing/code2/src/poms"
	"github.com/GOCODE/pluralsight/go-testing/code2/src/poms/ctrl"
)

func main() {

	ctrl.Setup()

	port := os.Getenv("APP_PORT")

	go http.ListenAndServe(":"+port, new(poms.GZipServer))

	log.Printf("Server started on port: %v, press <ENTER> to exit", port)
	fmt.Scanln()
}
