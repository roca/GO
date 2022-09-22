package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *server) discovery(w http.ResponseWriter, r *http.Request) {
	var (
		clientID string
	)

	if clientID = r.URL.Query().Get("client_id"); clientID == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("client_id is missing"))
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
	body, err := json.MarshalIndent(appConfig, "", "  ")
	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("json.MarshalIndent error: %s", err))
		return
	}

	w.WriteHeader(http.StatusFound)
	w.Write(body)

}
