package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_healthHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	api := API{log: slog.Default()}
	api.healthHandler(w, r)

	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
