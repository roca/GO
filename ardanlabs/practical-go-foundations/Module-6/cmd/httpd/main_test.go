package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_healthHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	healthHandler(w, r)

	resp := w.Result()

	fmt.Println(resp)

}
