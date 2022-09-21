package server

import (
	"fmt"
	"net/http"

	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/oidc"
)

func (s *server) authorization(w http.ResponseWriter, r *http.Request) {
	var (
		clientID     string
		redirectURI  string
		scope        string
		responseType string
		state        string
	)

	if clientID = r.URL.Query().Get("client_id"); clientID == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("client_id is missing"))
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

	appConfig := AppConfig{}
	for _, app := range s.Config.Apps {
		if app.ClientID == clientID {
			appConfig = app
		}
	}
	if appConfig.ClientID == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("client_id not found"))
		return
	}
	found := false
	for _, uri := range appConfig.RedirectURIs {
		if uri == redirectURI {
			found = true
		}
	}
	if !found {
		returnError(w, http.StatusBadRequest, fmt.Errorf("redirect_uri not whitelisted"))
		return
	}

	sessionID, err := oidc.GetRandomString(128)
	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("GetRandomString error: %s", err))
		return
	}

	s.LoginRequests[sessionID] = LoginRequest{
		ClientID:     clientID,
		RedirectURI:  redirectURI,
		Scope:        scope,
		ResponseType: responseType,
		State:        state,
	}

	w.Header().Add("location", fmt.Sprintf("%s/login?session_id=%s", s.Config.Url, sessionID))
	w.WriteHeader(http.StatusFound)
}
