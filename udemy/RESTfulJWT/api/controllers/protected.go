package controllers

import (
	"net/http"

	"udemy.com/RESTfulJWT/api/utils"
)

// ProtectedEndpoint ...
func (c Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ResponseJSON(w, "successfully called protected")
	}
}
