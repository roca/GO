package chatsess

import "testing"

func TestPass(t *testing.T) {
	h := NewPassword("hello")

	if !CheckPassword("hello", h) {
		t.Error("hello no match")
	}

	if CheckPassword("goodbye", h) {
		t.Error("goodby macthes")
	}
}
