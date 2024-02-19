package main

import (
	"math/rand"
	"testing"
)

func Test_generateIV(t *testing.T) {
	rand.Seed(0)
	for i := 8; i < 17; i++ {
		iv, err := generateIV(i)
		if err != nil {
			t.Log(err)
			continue
		}
		if len(iv) != i {
			t.Errorf("expected length of iv: %d, got: %d", i, len(iv))
		}
	}
}
