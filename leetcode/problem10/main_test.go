package main

import "testing"


func TestIsMatch(t *testing.T) {
	if isMatch("aa", "a") {
		t.Error("Expected false, got true")
	}

	if !isMatch("aa", "a*") {
		t.Error("Expected true, got false")
	}

	if !isMatch("aa", ".*") {
		t.Error("Expected true, got false")
	}

	if !isMatch("ab", ".*") {
		t.Error("Expected true, got false")
	}

	if !isMatch("abc", "a***abc") {
		t.Error("Expected true, got false")
	}
}
