package main

import "testing"

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as test parameter, expected %t, got %t", 0, false, result)
	}

	if msg != "0 is not prime, by definition!" {
		t.Errorf("with %d as test parameter, expected %s, got %s", 0, "0 is not prime, by definition!", msg)
	}
}
