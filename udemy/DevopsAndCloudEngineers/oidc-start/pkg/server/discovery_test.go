package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/oidc"
)

func TestDiscovery(t *testing.T) {
	s := newServer(privkeyPem, testConfig)

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

	discovery := oidc.Discovery{}
	err = json.Unmarshal(body, &discovery)
	if err != nil {
		t.Errorf("json.Unmarshal error: %s", err)
	}

	if discovery.Issuer != s.Config.Url {
		t.Errorf("Expected issuer %s, got %s", s.Config.Url, discovery.Issuer)
	}

	t.Log(string(body), res.StatusCode)

}
