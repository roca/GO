package main

import (
	"testing"
)

func TestIsOdd(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want error
	}{
		{"odd", 1, nil},
		{"even", 2, &EvenNumberError{Number: 2}},
		{"odd", 3, nil},
		{"even", 4, &EvenNumberError{Number: 4}},
		{"odd", 5, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOdd(tt.n); got != nil && got.Error() != tt.want.Error()  {
				t.Errorf("isOdd() = %v, want %v", got, tt.want)
			} else if got == nil && tt.want != nil {
				t.Errorf("isOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
