package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/roca/go-toolkit/v2"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

var tools toolkit.Tools

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {

	// read a json payload
	var creds Credentials
	err := tools.ReadJSON(w, r, &creds)
	log.Println(creds)
	if err != nil {
		tools.ErrorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// look up the user by email address
	user, err := app.DB.GetUserByEmail(creds.Username)
	if err != nil {
		tools.ErrorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// check the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		tools.ErrorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// generate a tokens
	tokenPairs, err := app.generateTokenPair(user)
	if err != nil {
		tools.ErrorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// send token to user
	_ = tools.WriteJSON(w, http.StatusOK, tokenPairs, nil)
}

func (app *application) refresh(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	refreshToken := r.Form.Get("refresh_token")
	claims := &Claims{}

	_, err = jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.JWTSecret), nil
	})
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) > 30*time.Second {
		tools.ErrorJSON(w, errors.New("refresh token does not need renewal yet"), http.StatusTooEarly)
		return
	}
}

func (app *application) allUser(w http.ResponseWriter, r *http.Request) {}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {}

func (app *application) insertUser(w http.ResponseWriter, r *http.Request) {}
