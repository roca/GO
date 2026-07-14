package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"udemy.com/RESTfulJWT/api/controllers"
	"udemy.com/RESTfulJWT/api/driver"
	"udemy.com/RESTfulJWT/api/models"
	"udemy.com/RESTfulJWT/api/utils"
)

func Test_Signup(t *testing.T) {

	controller := controllers.Controller{}

	db = driver.ConnectDB()

	var expectedUser models.User

	aUser := models.User{
		Email:    "JoeSmith@testing.com",
		Password: "qwertyu",
	}

	jsonUser, _ := json.Marshal(&aUser)
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Signup(db))

	handler.ServeHTTP(rr, req)

	// Should return a 200 status code
	assert.Equal(t, 200, rr.Code, "OK response is expected")

	// Should return a json object of type User
	json.NewDecoder(rr.Body).Decode(&expectedUser)
	assert.Equal(t, aUser.Email, expectedUser.Email, "Returned Email should be blank")
	// Returned Password should be blank
	assert.Equal(t, "", expectedUser.Password, "Returned Password should be blank")

}

func Test_Login(t *testing.T) {

	controller := controllers.Controller{}

	db = driver.ConnectDB()

	var expectedJWT models.JWT
	var userJWT models.JWT

	aUser := models.User{
		Email:    "JoeSmith@testing.com",
		Password: "qwertyu",
	}

	userJWTTokeString, _ := utils.GenerateToken(aUser)
	userJWT.Token = userJWTTokeString

	jsonUser, _ := json.Marshal(&aUser)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)

	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Login(db))

	handler.ServeHTTP(rr, req)

	// Should return a 200 status code
	assert.Equal(t, 200, rr.Code, "OK response is expected")

	// Should return a JWT token
	json.NewDecoder(rr.Body).Decode(&expectedJWT)
	assert.Equal(t, userJWT, expectedJWT, "Returned inncorrect JWT token")

}

func Test_Protected(t *testing.T) {

	controller := controllers.Controller{}

	db = driver.ConnectDB()

	var userJWT models.JWT

	aUser := models.User{
		Email:    "JoeSmith@testing.com",
		Password: "qwertyu",
	}

	userJWTTokeString, _ := utils.GenerateToken(aUser)
	userJWT.Token = userJWTTokeString

	req, err := http.NewRequest("GET", "/protected", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", userJWT.Token))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.TokenVerifyMiddleware(controller.ProtectedEndpoint()))

	handler.ServeHTTP(rr, req)

	message := ""
	json.NewDecoder(rr.Body).Decode(&message)

	assert.Equal(t, "successfully called protected", message, "Bearer token dose not match")

}
