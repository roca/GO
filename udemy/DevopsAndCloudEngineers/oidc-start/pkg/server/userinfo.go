package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/users"
)

func (s *server) userinfo(w http.ResponseWriter, r *http.Request) {

	bearerText := r.Header.Get("Authorization")
	if bearerText == "" {
		returnError(w, http.StatusUnauthorized, fmt.Errorf("Authorization header not set"))
		return
	}
	accessToken := bearerText[7:]

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		privateKeyParsed, err := jwt.ParseRSAPrivateKeyFromPEM(s.PrivateKey)
		if err != nil {
			return nil, err
		}
		return &privateKeyParsed.PublicKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		returnError(w, http.StatusUnauthorized, fmt.Errorf("Invalid token error: %s", err))
		return
	}

	found := claims.VerifyAudience(s.Config.Url + "/userinfo", true)
	if !found {
		returnError(w, http.StatusUnauthorized, fmt.Errorf("Invalid audience %s", claims["aud"]))
		return
	}

	for _, user := range users.GetAllUsers() {
		if user.Sub == claims["sub"].(string) {
			bytes, err := json.Marshal(user)
			if err != nil {
				returnError(w, http.StatusInternalServerError, fmt.Errorf("userinfo marshall error: %s", err))
				return
			}
			w.Write(bytes)
			return
		}
	}

	w.Write([]byte("User not found"))

}
