package server

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/oidc"
	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/users"
)

//go:embed templates/*
var templateFs embed.FS

func (s *server) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.loginGet(w, r)
		return
	}
	s.loginPost(w, r)
}

func (s *server) loginGet(w http.ResponseWriter, r *http.Request) {
	var sessionID string

	if sessionID = r.URL.Query().Get("sessionID"); sessionID == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("session_id is missing"))
		return
	}

	if _, ok := s.LoginRequests[sessionID]; !ok {
		returnError(w, http.StatusBadRequest, fmt.Errorf("session_id is invalid"))
		return
	}

	templateFile, err := templateFs.Open("templates/login.html")
	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("template error: %s", err))
		return
	}

	templatFileBytes, err := io.ReadAll(templateFile)
	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("ReadAll error: %s", err))
		return
	}

	templatFileStr := strings.Replace(string(templatFileBytes), "$SESSIONID", sessionID, -1)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(templatFileStr))

}

func (s *server) loginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	login := r.Form.Get("login")
	password := r.Form.Get("password")
	sessionID := r.Form.Get("sessionID")

	exists, user, err := users.Auth(login, password, "")
	if err != nil {
		returnError(w, http.StatusUnauthorized, fmt.Errorf("authentication error: %s", err))
		return
	}

	if !exists {
		returnError(w, http.StatusUnauthorized, fmt.Errorf("authentication error: user not found"))
		return
	}

	loginRequest, ok := s.LoginRequests[sessionID]
	if !ok {
		returnError(w, http.StatusBadRequest, fmt.Errorf("session_id is invalid"))
		return
	}

	redirectURI := loginRequest.RedirectURI
	code, err := oidc.GetRandomString(64)

	loginRequest.CodeIssuedAt = time.Now()
	loginRequest.User = user
	s.Codes[code] = loginRequest

	delete(s.LoginRequests, sessionID) // So the user can't login twice

	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("GetRandomString error: %s", err))
		return
	}

	newURI := fmt.Sprintf("%s?code=%s&state=%s", redirectURI, code, loginRequest.State)
	w.Header().Add("location", newURI)
	w.WriteHeader(http.StatusFound)
}
