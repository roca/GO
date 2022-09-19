package server

import (
	"fmt"
	"net/http"
)

func (s *server) authorization(w http.ResponseWriter, r *http.Request) {
	var (
		clientID     string
		clientSecret string
		redirectURI  string
		scope        string
		responseType string
		state        string
	)

	w.Header().Set("location", "https://localhost:8080/login")
	if clientID = r.URL.Query().Get("client_id"); clientID == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("client_id is missing"))
		return
	}
	if clientSecret = r.URL.Query().Get("client_secret"); clientSecret == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("client_secret is missing"))
		return
	}
	if redirectURI = r.URL.Query().Get("redirect_uri"); redirectURI == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("redirect_uri is missing"))
		return
	}
	if scope = r.URL.Query().Get("scope"); scope != "openid" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("scope is missing or invalid"))
		return
	}
	if responseType = r.URL.Query().Get("response_type"); responseType != "code" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("response_type is missing or invalid"))
		return
	}
	if state = r.URL.Query().Get("state"); state == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("state is missing"))
		return
	}

	w.WriteHeader(http.StatusFound)
}
