package main

import (
	"math/big"
	"testing"
)

func Test_getN(t *testing.T) {
	tests := []struct {
		name string
		int1 int64
		int2 int64
	}{
		{"test1", 512, 512},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := big.NewInt(tt.int1)
			q := big.NewInt(tt.int2)
			got := getN(p, q)
			want := p.Mul(p, q)
			if got != want {
				t.Errorf("got: %v, want: %v", firstNDigits(*got, 10), firstNDigits(*want, 10))
			}
		})
	}
}

func Test_generatePrivateNums(t *testing.T) {
	tests := []struct {
		name                   string
		keySize                int
		expectedLength         int
		expectedLengthOfProduct int
	}{
		{"test_5", 5, 2, 3},
		{"test_512", 512, 155, 309},
		{"test_1024", 1024, 309, 617},
		{"test_2048", 2048, 617, 1233},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, q := generatePrivateNums(tt.keySize)
			z := *p
			product := getN(&z, q)
			if len(p.String()) != tt.expectedLength ||
				len(q.String()) != tt.expectedLength ||
				len(product.String()) != tt.expectedLengthOfProduct {
				t.Errorf("Lengths are incorrect got: %d, %d, %d want: %d, %d", len(p.String()), len(q.String()), len(product.String()), tt.expectedLength, tt.expectedLengthOfProduct)
			}
			if len(p.String()) != tt.expectedLength || len(q.String()) != tt.expectedLength {
				t.Errorf("Lengths are incorrect got: %d, %d, want: %d", len(p.String()), len(q.String()), tt.expectedLength)
			}

		})
	}
}
