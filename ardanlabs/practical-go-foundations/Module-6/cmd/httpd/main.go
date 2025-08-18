package main

import (
	"encoding/json"
	"expvar"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/roca/GO/tree/staging/ardanlabs/practical-go-foundations/Module-5/nlp"
	"github.com/roca/GO/tree/staging/ardanlabs/practical-go-foundations/Module-5/nlp/stemmer"
)

var (
	stemCalls = expvar.NewInt("stemCalls")
)

/* Configuration Pattern

   Defaults < Configuration file < Environment variables < Command line

   Configuration: YAML, TOML ...

   Environment: os.Getenv("YOUR_VAR")

   Command line: flags package


   External:

   - viper + cobra
   - pkg.go.dev/github.com/ardanlabs/conf/v3

*/

var config struct {
	Addr string
}

func main() {

	config.Addr = os.Getenv("NLP_ADDR")
	if config.Addr == "" {
		config.Addr = ":8080"
	}
	flag.StringVar(&config.Addr, "addr", config.Addr, "address port (exp ':8080') to listen on")
	flag.Parse()

	// IMPORTANT: validate-config

	if err := health(); err != nil {
		fmt.Fprintf(os.Stderr, "error: health check - %s\n", err)
		os.Exit(1)
	}

	api := API{log: slog.Default().With("app", "nlp")}
	// Routing
	http.HandleFunc("GET /health", api.healthHandler)
	http.HandleFunc("GET /stem/{word}", api.stemHandler)
	http.HandleFunc("POST /tokenize", api.tokenizeHandler)

	api.log.Info("server staring", "address", config.Addr)
	if err := http.ListenAndServe(config.Addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

}

type API struct {
	log *slog.Logger
}

func (a *API) stemHandler(w http.ResponseWriter, r *http.Request) {
	stemCalls.Add(1)
	if r.Method != "GET" {
		http.Error(w, "/stem only handles GET requests", http.StatusBadRequest)
		return
	}

	word := r.PathValue("word")
	a.log.Info("Stem", "word", word)

	fmt.Fprintln(w, "word stem:", stemmer.Stem(word))
}

func (a *API) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
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
		a.log.Error("Tokenize", "error", err)
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

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := health(); err != nil {
		a.log.Error("Health", "error", err)
		http.Error(w, "health check failed", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "OK")
}

func health() error {
	// TODO: Actual health check
	return nil
}
