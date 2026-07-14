package main

import "testing"

func Test_alphabetSize(t *testing.T) {
	tests := []struct {
		name    string
		numBits int
		want    float64
	}{
		{"Two", 2, 4},
		{"Four", 4, 16},
		{"Seven", 7, 128},
		{"Eight", 8, 256},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alphabetSize(tt.numBits)
			if got != tt.want {
				t.Errorf("Wanted %v but got %v", tt.want, got)
			}
		})
	}
}
