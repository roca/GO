package main

import "testing"

func TestGetIntFromString(t *testing.T) {

	testCases := map[string]int{
		"1":   1,
		"2.1": 2,
	}

	for k, expected := range testCases {

		v, _ := GetIntFromString(k)
		if v != expected {
			t.Errorf("Failed to parse %s string to %d", k, expected)
		}

	}
}
