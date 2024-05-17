package main

import (
	"go-breaders/configuration"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {

	testApp = application{
		App: configuration.New(nil),
	}

	os.Exit(m.Run())
}
