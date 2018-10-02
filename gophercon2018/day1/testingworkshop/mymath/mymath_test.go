package mymath_test

import (
	"testing"

	"github.com/GOCODE/gophercon2018/day1/testingworkshop/mymath"
)

func TestRace(t *testing.T) {
	got := mymath.AddTwoNumber(3.0, 3.0)
	if got != 6.0 {
		t.Fatalf("failed got %v expected %v", got, 5.0)
	}
}
