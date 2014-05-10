package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type handler struct {
	Word1 string
	Word2 string
}


func (h *handler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
    if err := enc.Encode(&h); nil != err {
        fmt.Fprintf(w, `{"error":"%s"}`, err)
    }
}

func main() {
	myHandler := handler{"Hello","World"}
	http.HandleFunc(
		"/", myHandler.ServeHTTP,
	)
	http.ListenAndServe(":8080", nil)
}
