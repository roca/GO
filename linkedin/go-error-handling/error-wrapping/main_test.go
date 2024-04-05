package main

import (
	"fmt"
	"testing"
)

func TestIsOdd(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want error
	}{
		{"odd", 1, nil},
		{"even", 2, fmt.Errorf("Number %d is even", 2)},
		{"odd", 3, nil},
		{"even", 4, fmt.Errorf("Number %d is even", 4)},
		{"odd", 5, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOdd(tt.n); got != tt.want {
				t.Errorf("isOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
