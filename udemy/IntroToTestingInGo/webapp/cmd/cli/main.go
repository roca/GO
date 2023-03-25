package main

import (
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type application struct {
	JWTSecret string
	Action    string
}

// This is used to generate a token, so that we can test our api. Run this with go run ./cmd/cli and copy
// the token that is printed out.
// go run ./cmd/cli -action=valid     // will produce a valid token
// go run ./cmd/cli -action=expired   // will produce an expired token

func main() {
	var app application
	flag.StringVar(&app.JWTSecret, "jwt-secret", "some-very-secret-secret", "secret")
	flag.StringVar(&app.Action, "action", "valid", "action: valid|expired")
	flag.Parse()

	// generate a token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["sub"] = "1"
	claims["admin"] = true
	claims["aud"] = "example.com"
	claims["iss"] = "example.com"
	// leave this to 3 days, for easy manual testing
	if app.Action == "valid" {
		expires := time.Now().UTC().Add(time.Hour * 72)
		claims["exp"] = expires.Unix()
	} else {
		expires := time.Now().UTC().Add(time.Hour * 100 * -1)
		claims["exp"] = expires.Unix()
	}

	// create the token as a slice of bytes
	if app.Action == "valid" {
		fmt.Println("VALID Token:")
	} else {
		fmt.Println("EXPIRED Token:")
	}
	signedAccessToken, err := token.SignedString([]byte(app.JWTSecret))
	if err != nil {
		log.Fatal(err)
	}
	// print to console
	fmt.Println(string(signedAccessToken))
}
