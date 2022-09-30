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
func getTokenFromCode(tokenUrl, jwksUrl, redirectUri, clientID, clientSecret, code string) (*jwt.Token, *jwt.StandardClaims, error) {

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
		return nil, nil, err
	}

	if tokenResponse.IDToken == "" {
		return nil, nil, err
	}
	fmt.Print(tokenResponse.IDToken)

	return nil, nil, fmt.Errorf("Not implemented")

	// claims := jwt.StandardClaims{}
	// publicKey, err := ioutil.ReadFile("server.pub")
	// if err != nil {
	// 	return nil, nil, err
	// }

	// publicKeyParsed, _, _, _, err := ssh.ParseAuthorizedKey(publicKey)
	// if err != nil {
	// 	return nil, nil, err
	// }

	// token, err := jwt.ParseWithClaims(tokenResponse.IDToken, &claims, func(token *jwt.Token) (interface{}, error) {

	// 	return publicKeyParsed, nil
	// })
	// if err != nil {
	// 	return nil, nil, err
	// }

	// return token, &claims, nil
}
