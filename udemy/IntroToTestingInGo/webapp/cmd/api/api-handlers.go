package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/roca/go-toolkit/v2"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools

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

func (app *application) refresh(w http.ResponseWriter, r *http.Request) {}

func (app *application) allUser(w http.ResponseWriter, r *http.Request) {}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {}

func (app *application) insertUser(w http.ResponseWriter, r *http.Request) {}
