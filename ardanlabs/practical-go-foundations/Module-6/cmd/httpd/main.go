package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/roca/GO/tree/staging/ardanlabs/practical-go-foundations/Module-5/nlp"
	"github.com/roca/GO/tree/staging/ardanlabs/practical-go-foundations/Module-5/nlp/stemmer"
)

func main() {
	// Routing
	http.HandleFunc("GET /health", healthHandler)
	http.HandleFunc("GET /stem/{word}", stemHandler)
	http.HandleFunc("POST /tokenize", tokenizeHandler)

	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

}

func stemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "/stem only handles GET requests", http.StatusBadRequest)
		return
	}

	word := r.PathValue("word")

	fmt.Fprintln(w, "word stem:", stemmer.Stem(word))
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "/tokenize only handles POST requests", http.StatusBadRequest)
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

	if &text == nil || text.Text == "" {
		http.Error(w, "text data can't be null", http.StatusBadRequest)
		return
	}

	tokens := nlp.Tokenize(text.Text)

	w.Header().Set("content-type", "application/json")
	resp := map[string][]string{
		"tokens": tokens,
	}
	json.NewEncoder(w).Encode(resp)

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
