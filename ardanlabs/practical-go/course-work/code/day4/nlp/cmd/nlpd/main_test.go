package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	// Note: This bypasses the routing & middleware
	healthHandler(w, r)

	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
