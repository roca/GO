package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Signup(t *testing.T) {

	var expectedUser User

	aUser := User{
		Email:    "JoeSmith@testing.com",
		Password: "qwertyu",
	}

	jsonUser, _ := json.Marshal(&aUser)
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(signup)

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

	var expectedJWT JWT
	var userJWT JWT

	aUser := User{
		Email:    "JoeSmith@testing.com",
		Password: "qwertyu",
	}

	userJWTTokeString, _ := GenerateToken(aUser)
	userJWT.Token = userJWTTokeString

	jsonUser, _ := json.Marshal(&aUser)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login)

	handler.ServeHTTP(rr, req)

	// Should return a 200 status code
	assert.Equal(t, 200, rr.Code, "OK response is expected")

	// Should return a JWT token
	json.NewDecoder(rr.Body).Decode(&expectedJWT)
	assert.Equal(t, userJWT, expectedJWT, "Returned inncorrect JWT token")

}

func Test_Protected(t *testing.T) {

	var userJWT JWT

	aUser := User{
		Email:    "JoeSmith@testing.com",
		Password: "qwertyu",
	}

	userJWTTokeString, _ := GenerateToken(aUser)
	userJWT.Token = userJWTTokeString

	req, err := http.NewRequest("GET", "/protected", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", userJWT.Token))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TokenVerifyMiddleware(protectedEndpoint))

	handler.ServeHTTP(rr, req)
	assert.Equal(t, "successfully called protected", rr.Body.String(), "Bearer token dose not match")

}
