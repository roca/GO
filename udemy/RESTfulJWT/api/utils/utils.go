package utils

import (
	"encoding/json"
	"net/http"

	"udemy.com/RESTfulJWT/api/models"
)

// RespondWithError ...
func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

// ResponseJSON ...
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
