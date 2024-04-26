package main

import (
	"fmt"
	"net/http"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
