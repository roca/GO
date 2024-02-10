package main

import (
	"regexp"
	"testing"
)

func Test_generateRandomKey(t *testing.T) {
	i := 32
	randomString, err := generateRandomKey(i)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(randomString) != 2*i {
		t.Errorf("Error: Expected 16, got %v", len(randomString))
	}
	match, err := regexp.MatchString("[a-z,A_Z]+|[0-9]+", randomString )
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if !match {
		t.Errorf("Error: Expected true, got %v", match)
	}
}
