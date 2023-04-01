package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
	"webapp/pkg/data"

	"github.com/go-chi/chi/v5"
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

	// get the user id from the claims
	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.DB.GetUser(userID)
	if err != nil {
		tools.ErrorJSON(w, errors.New("Unknown user"), http.StatusBadRequest)
		return
	}

	// generate a tokens
	tokenPairs, err := app.generateTokenPair(user)
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "__Host-refresh_token",
		Path:     "/",
		Value:    tokenPairs.RefreshToken,
		Expires:  time.Now().Add(refreshTokenExpiry),
		MaxAge:   int(refreshTokenExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   true,
	})

	_ = tools.WriteJSON(w, http.StatusOK, tokenPairs, nil)
}

func (app *application) allUser(w http.ResponseWriter, r *http.Request) {
	users, err := app.DB.AllUsers()
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = tools.WriteJSON(w, http.StatusOK, users, nil)
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.DB.GetUser(id)
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = tools.WriteJSON(w, http.StatusOK, user, nil)
}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := tools.ReadJSON(w, r, &user)
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.DB.UpdateUser(user)
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	
	err = app.DB.DeleteUser(id)
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func (app *application) insertUser(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := tools.ReadJSON(w, r, &user)
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	
	_, err = app.DB.InsertUser(user)
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
