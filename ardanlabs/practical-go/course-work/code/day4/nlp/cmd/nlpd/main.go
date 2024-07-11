package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"nlp"

	"github.com/gorilla/mux"
)

func main() {
	//routing
	// /health is a exact match
	// /health/ is a prefix match
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", tokenizeHandler).Methods(http.MethodPost)
	r.HandleFunc("/stem/{word}", stemHandler).Methods(http.MethodGet)
	http.Handle("/", r)

	// start a web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func stemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
}

// exercise: Write a tokenizeHandler that will read  the text from the request
// body and return JSON in the format: {"tokens": ["word1", "word2", ...]}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	/* Before gorilla/mux
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	*/

	// bytes, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	rdr := io.LimitReader(r.Body, 1_000_000)
	bytes, err := io.ReadAll(rdr)
	if err != nil {
		http.Error(w, "error reading request body", http.StatusBadRequest)
		return
	}

	if len(bytes) == 0 {
		http.Error(w, "empty request body", http.StatusBadRequest)
		return
	}

	resp := struct {
		Tokens []string `json:"tokens"`
		Ok     bool     `json:"ok"`
	}{
		Tokens: nlp.Tokenize(string(bytes)),
		Ok:     true,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
}
