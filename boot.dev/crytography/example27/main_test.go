package main

import (
	"math/big"
	"testing"
)

func Test_getTot(t *testing.T) {
	p := big.NewInt(11)
	q := big.NewInt(7)
	expected := "60"

	tot := getTot(p, q)
	if tot.String() != expected {
		t.Errorf("Expected %s, got %s", expected, tot.String())
	}
}

func Test_getE(t *testing.T) {
	tot := big.NewInt(11)
	expected := "3"

	e := getE(tot)
	if e.String() != expected {
		t.Errorf("Expected %s, got %s", expected, e.String())
	}
}
