package main

import (
	"testing"
)

func TestApp_getGoldText(t *testing.T) {
	open, _, _ := testApp.getGoldText()
	if open.Text != "Open: $2005.2650 USD" {
		t.Errorf("Expected %s, got %s", "Open: $2005.2650 USD", open.Text)
	}
}
