package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		want    int
		wantErr bool
	}{
		{"F(0)", 0, 0, false},
		{"F(1)", 1, 1, false},
		{"F(2)", 2, 1, false},
		{"F(3)", 3, 2, false},
		{"F(4)", 4, 3, false},
		{"F(5)", 5, 5, false},
		{"F(10)", 10, 55, false},
		{"F(15)", 15, 610, false},
		{"F(20)", 20, 6765, false},
		{"negative input", -1, 0, true},
		{"negative input large", -100, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fibonacci(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("fibonacci(%d) error = %v, wantErr %v", tt.n, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
