package main

import "testing"

func Test_generateRandomKey(t *testing.T) {
	i := 32
	randomString , err := generateRandomKey(i)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(randomString) != 2 * i {
		t.Errorf("Error: Expected 16, got %v", len(randomString))
	}
}
