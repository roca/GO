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
			got := getN(p,q)
			want := p.Mul(p, q)
			if got != want {
				t.Errorf("got: %v, want: %v", firstNDigits(*got, 10), firstNDigits(*want, 10))
			}
		})
	}
}

func Test_generatePrivateNums(t *testing.T) {

}
