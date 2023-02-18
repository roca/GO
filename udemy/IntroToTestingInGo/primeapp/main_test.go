package main

import "testing"

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name string
		testNum int
		expected bool
		msg string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not prime, it is divisible by 2"},
		{"Zero", 0, false, "0 is not prime, by definition!"},
		{"One", 1, false, "1 is not prime, by definition!"},
		{"Negative", -1, false, "Negative numbers are not prime, by definition!"},
		{"Four", 4, false, "4 is not prime, it is divisible by 2"},
	}

	for _, e := range primeTests {
		t.Run(e.name, func(t *testing.T) {
			prime, msg := isPrime(e.testNum)
			if prime != e.expected {
				t.Errorf("Expected %t, got %t", e.expected, prime)
			}
			if msg != e.msg {
				t.Errorf("Expected %s, got %s", e.msg, msg)
			}
		})	
	}
}
