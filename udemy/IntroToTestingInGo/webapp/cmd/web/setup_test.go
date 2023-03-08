package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	app.Session = getSession()

	app.DB = &dbrepo.TestDBRepo{}

	code := m.Run()

	os.Exit(code)
}
