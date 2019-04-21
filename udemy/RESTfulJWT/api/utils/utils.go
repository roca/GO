package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"udemy.com/RESTfulJWT/api/models"
)

// RespondWithError ...
func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

// ResponseJSON ...
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

// GenerateToken ..
func GenerateToken(user models.User) (string, error) {
	var err error

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}
