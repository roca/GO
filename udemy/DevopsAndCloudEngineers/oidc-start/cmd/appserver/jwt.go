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

	return nil, nil, nil
}
