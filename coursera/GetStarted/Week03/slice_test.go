package main

import (
	"testing"
)

func TestGetIntFromString(t *testing.T) {

	testCases := map[string]int{
		"1":   1,
		"2.1": 2,
		"x":   0,
		"X":   0,
	}

	for k, expected := range testCases {

		v, err := GetIntFromString(k)
		if v != expected {
			t.Errorf("Failed to parse %s string to %d", k, expected)
		}

		if k == "x" && err.Error() != "Done" {
			t.Error("Error message should read 'Done'")
		}

		if k == "X" && err.Error() != "Done" {
			t.Error("Error message should read 'Done'")
		}

	}
}
