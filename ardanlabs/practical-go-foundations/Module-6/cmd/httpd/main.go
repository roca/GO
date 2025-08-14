package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/roca/GO/tree/staging/ardanlabs/practical-go-foundations/Module-5/nlp"
)

func main() {
	// Routing
	http.HandleFunc("GET /health", healthHandler)
	http.HandleFunc("POST /tokenize", tokenizeHandler)

	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "tokenizeHandler only handles POST requests", http.StatusBadRequest)
		return
	}

	var text struct {
		Text string `json:"text"`
	}

	d := json.NewDecoder(r.Body)

	err := d.Decode(&text)
	if err != nil {
		http.Error(w, "Could not get text data", http.StatusBadRequest)
		return
	}

	a := nlp.Tokenize(text.Text)

	fmt.Fprintln(w, a)

	// fmt.Fprintf(os.Stdout, "text: %v\n", a)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := health(); err != nil {
		http.Error(w, "health check failed", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "OK")
}

func health() error {
	// TODO: Actual health check
	return nil
}
