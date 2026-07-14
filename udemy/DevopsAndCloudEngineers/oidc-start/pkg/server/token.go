package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/oidc"
	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/users"
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

	loginRequest, ok := s.Codes[code]
	if !ok {
		returnError(w, http.StatusBadRequest, fmt.Errorf("code is invalid"))
		return
	}

	if time.Now().After(loginRequest.CodeIssuedAt.Add(10*time.Minute)) == true {
		returnError(w, http.StatusBadRequest, fmt.Errorf("code is expired"))
		return
	}

	if loginRequest.ClientID != client_id {
		returnError(w, http.StatusBadRequest, fmt.Errorf("client_id does not match"))
		return
	}

	if loginRequest.AppConfig.ClientSecret != client_secret {
		returnError(w, http.StatusBadRequest, fmt.Errorf("client_secret does not match"))
		return
	}

	if loginRequest.RedirectURI != redirect_uri {
		returnError(w, http.StatusBadRequest, fmt.Errorf("redirect_uri does not match"))
		return
	}

	signedIDToken, err := generateIDJWT(loginRequest.User, loginRequest.AppConfig, s.PrivateKey)
	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("error generating signedIDToken JWT: %s", err))
		return
	}

	signedAccessToken, err := generateAccessJWT(loginRequest.User, loginRequest.AppConfig, s.PrivateKey, s.Config.Url)
	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("error generating signedAccessToken JWT: %s", err))
		return
	}

	tokenOutput := oidc.Token{
		IDToken:     signedIDToken,
		AccessToken: signedAccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   60,
	}

	delete(s.Codes, code)

	tokenBytes, err := json.Marshal(tokenOutput)
	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("json.Marshal(token) error: %s", err))
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	w.Write(tokenBytes)

}

func generateIDJWT(user users.User, appConfig AppConfig, privateKey []byte) (string, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": user.Sub,
		"iss": appConfig.Issuer,
		"aud": appConfig.ClientID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		// "authorized": true,
		// "email":      user.Email,
	})

	token.Header["kid"] = "0-0-0-1"

	signedIDToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedIDToken, nil

}

func generateAccessJWT(user users.User, appConfig AppConfig, privateKey []byte, appURL string) (string, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": user.Sub,
		"iss": appConfig.Issuer,
		"aud": []string{
			appURL + "/userinfo",
		},
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		// "authorized": true,
		// "email":      user.Email,
	})

	token.Header["kid"] = "0-0-0-1"

	signedIDToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedIDToken, nil

}
