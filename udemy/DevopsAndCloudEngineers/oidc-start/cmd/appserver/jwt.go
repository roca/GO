package main

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
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
	// fmt.Print(tokenResponse.IDToken)

	claims := jwt.RegisteredClaims{}
	parsedToken, err := jwt.ParseWithClaims(tokenResponse.IDToken, &claims, func(token *jwt.Token) (interface{}, error) {
		kid, ok := token.Header["kid"]
		if !ok {
			return nil, fmt.Errorf("kid not found in token header")
		}
		publiucKey, err := getPublicKeyFromJwks(jwksUrl, kid.(string))
		if err != nil {
			return nil, fmt.Errorf("getPublicKeyFromJwks error: %s", err)
		}

		return publiucKey, nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("Token parsing failed: %s", err)
	}

	// return token, &claims, nil
	return parsedToken, &claims, nil
}

func getPublicKeyFromJwks(jwksUrl, kid string) (*rsa.PublicKey, error) {
	res, err := http.Get(jwksUrl)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode was not 200")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var jwks oidc.Jwks

	err = json.Unmarshal(body, &jwks)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal jwks error: %s", err)
	}

	for _, key := range jwks.Keys {
		if key.Kid == kid {
			nBytes, err := base64.StdEncoding.DecodeString(key.N)
			if err != nil {
				return nil, fmt.Errorf("base64 decode error: %s", err)
			}
			n := big.NewInt(0)
			n.SetBytes(nBytes)
			return &rsa.PublicKey{
				N: n,
				E: 65537,
			}, nil
		}
	}

	return nil, fmt.Errorf("No public key found for kid: %s", kid)
}
