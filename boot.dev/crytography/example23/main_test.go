package main

import "testing"

func Test_nonceStrength(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		strength int
	}{
		{"eight", 5, 1099511627776},
		{"four", 4, 4294967296},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nonce, err := generateIV(tt.length)
			if err != nil {
				t.Errorf("Could not create nonce. Error: %v\n", err)
			}
			got := nonceStrength(nonce)
			if  got != tt.strength {
				t.Errorf("Incorrect strength got: %d, want: %d \n", got, tt.strength)
			}
		})
	}
}
