package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/golang-jwt/jwt/v4"
	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/oidc"
)

// gets token from tokenUrl validating token with jwksUrl and returning token & claims
func getTokenFromCode(tokenUrl, jwksUrl, redirectUri, clientID, clientSecret, code string) (*jwt.Token, *jwt.RegisteredClaims, error) {

	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("redirect_uri", redirectUri)
	form.Add("client_id", clientID)
	form.Add("client_secret", clientSecret)
	form.Add("code", code)

	res, err := http.PostForm(tokenUrl, form)
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("StatusCode was not 200")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	var tokenResponse oidc.Token

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return nil, nil, fmt.Errorf("Unmarshal token error: %s", err)
	}

	if tokenResponse.IDToken == "" {
		return nil, nil, err
	}
	fmt.Print(tokenResponse.IDToken)

	claims := jwt.RegisteredClaims{}
	parsedToken, err := jwt.ParseWithClaims(tokenResponse.IDToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("Token parsing failed: %s", err)
	}

	// return token, &claims, nil
	return parsedToken, &claims, nil
}
