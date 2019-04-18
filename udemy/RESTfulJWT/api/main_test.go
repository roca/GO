package main

import (
	"bytes"
	"encoding/json"
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
