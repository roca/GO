package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"nlp"
	"nlp/stemmer"

	"github.com/gorilla/mux"
)

func main() {
	// Create server (dependency injection)
	logger := log.New(log.Writer(), "nlp ",log.LstdFlags|log.Lshortfile)
	server := Server{logger: logger}
	//routing
	// /health is a exact match
	// /health/ is a prefix match
	r := mux.NewRouter()

	r.HandleFunc("/health", server.healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", server.tokenizeHandler).Methods(http.MethodPost)
	r.HandleFunc("/stem/{word}", server.stemHandler).Methods(http.MethodGet)
	http.Handle("/", r)

	// start a web server
	addr := ":8080"
	server.logger.Printf("server starting on  %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

type Server struct {
	logger *log.Logger
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func (s *Server) stemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	if word == "" {
		http.Error(w, "missing word", http.StatusBadRequest)
		return
	}

	stem := stemmer.Stem(word)
	fmt.Fprintln(w, stem)
}

// exercise: Write a tokenizeHandler that will read  the text from the request
// body and return JSON in the format: {"tokens": ["word1", "word2", ...]}
func (s *Server) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
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
		s.logger.Printf("error reading request body: %s", err)
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
