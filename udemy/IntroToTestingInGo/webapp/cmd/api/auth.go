package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const jwtTokenExpiry = 15 * time.Minute
const refreshTokenExpiry = 24 * time.Hour

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	UserName string `json:"name"`
	jwt.RegisteredClaims
}

func (app *application) getTokenFromHeaderAndVerify(w http.ResponseWriter, r *http.Request) (string, *Claims, error) {
	// we expect the token to be in the Authorization header like this:
	// Bearer <token>
	// add a header
	w.Header().Add("Vary", "Authorization")

	// get the authorization header
	authHeader := r.Header.Get("Authorization")

	// sanity check
	if authHeader == "" {
		return "", nil, errors.New("missing authorization header")
	}

	// split the header on spaces
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", nil, errors.New("invalid authorization header")
	}

	// check to see if we have the word "Bearer" in the first part
	if headerParts[0] != "Bearer" {
		return "", nil, errors.New("invalid authorization header, no Bearer token")
	}

	token := headerParts[1]

	// declare an empty claims variable
	claims := &Claims{}

	// parse the token with our claims ( we read into claims), using our secret (from the receiver)
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// validate the signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(app.JWTSecret), nil
	})

	// check for error; not that this catches both expired and invalid tokens as well
	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired by") {
			return "", nil, errors.New("token is expired")
		}
		return "", nil, err
	}

	// make sure that we issued this token
	if claims.Issuer != app.Domain {
		return "", nil, errors.New("incorrect issuer")
	}

	// valodation passed
	return token, claims, nil
}
