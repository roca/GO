package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDiscovery(t *testing.T) {

	endpoint := "/.well-known/openid-configuration"
	req := httptest.NewRequest(http.MethodGet, endpoint, nil)
	w := httptest.NewRecorder()
	s.discovery(w, req)
	res := w.Result()

	if res.StatusCode != http.StatusFound {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Readall error: %s", err)
	}

	t.Log(string(body), res.StatusCode)

}
