package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/oidc"
)

func (s *server) token(w http.ResponseWriter, r *http.Request) {

	var (
		grant_type    string
		client_id     string
		client_secret string
		redirect_uri  string
		code          string
	)

	if r.Method != http.MethodPost {
		returnError(w, http.StatusMethodNotAllowed, fmt.Errorf("Method not allowed"))
		return
	}

	if grant_type = r.FormValue("grant_type"); grant_type == "" || grant_type != "authorization_code" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("grant_type is missing or invalid"))
		return
	}

	if client_id = r.FormValue("client_id"); client_id == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("client_id is required"))
		return
	}

	if client_secret = r.FormValue("client_secret"); client_secret == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("client_secret is required"))
		return
	}

	if redirect_uri = r.FormValue("redirect_uri"); redirect_uri == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("redirect_uri is required"))
		return
	}

	if code = r.FormValue("code"); code == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("code is missing"))
		return
	}

	loginRequest,ok := s.Codes[code]
	if !ok {
		returnError(w, http.StatusBadRequest, fmt.Errorf("code is invalid"))
		return
	}

	if loginRequest.ClientID != client_id {
		returnError(w, http.StatusBadRequest, fmt.Errorf("client_id does not match"))
		return
	}

	if s.Config.Apps["app1"].ClientSecret != client_secret {
		returnError(w, http.StatusBadRequest, fmt.Errorf("client_secret does not match"))
		return
	}

	found := false

	for _, uri := range s.Config.Apps["app1"].RedirectURIs {
		if uri == redirect_uri {
			found = true
			break
		}
	}

	if !found {
		returnError(w, http.StatusBadRequest, fmt.Errorf("redirect_uri does not match"))
		return
	}

	token := oidc.Token{}
	tokenBytes, err := json.Marshal(token)
	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("json.Marshal(token) error: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tokenBytes)

}
