package main

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/golang-jwt/jwt/v4"
)

// gets token from tokenUrl validating token with jwksUrl and returning token & claims
func getTokenFromCode(tokenUrl, jwksUrl, redirectUri, clientID, clientSecret, code string) (*jwt.Token, *jwt.StandardClaims, error) {

	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("redirect_uri", redirectUri)
	form.Add("client_id", clientID)
	form.Add("client_secret", clientSecret)
	form.Add("code", code)

	res, err := http.Post(tokenUrl, "application/x-www-form-urlencoded", bytes.NewBufferString(form.Encode()))
	if err != nil {
		return nil, nil, err
	}
	
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	var tokenResponse oidc.Token

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return nil, nil, err
	}

	if tokenResponse.IDToken == "" {
		return nil, nil, err
	}

	claims := jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(tokenResponse.IDToken, &claims, func(token *jwt.Token) (interface{}, error) {
		privateKeyParsed, err := jwt.ParseRSAPrivateKeyFromPEM(s.PrivateKey)
		if err != nil {
			return nil, nil, err
		}
		return &privateKeyParsed.PublicKey, nil
	})

	return nil, nil, nil
}
