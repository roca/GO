package api

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello Cloud Native Go. (docker image: rcampbell/cloud-native-go:1.0.1")
}
