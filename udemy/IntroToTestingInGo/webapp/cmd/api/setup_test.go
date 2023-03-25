package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/dbrepo"
)

var app application
var expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE2NzkzNzU1ODIsImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.lagC6EKn2gSSVGsKK3ptShPo3huq5k79h1QFXDAazU0"

func TestMain(m *testing.M) {

	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "some-very-secret-secret"

	code := m.Run()

	os.Exit(code)
}
