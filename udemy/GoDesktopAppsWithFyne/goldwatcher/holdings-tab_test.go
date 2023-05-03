package main

import "testing"

func TestConfig_currentHoldings(t *testing.T) {
	all, err := testApp.currentHoldings()
	if err != nil {
		t.Errorf("currentHoldings() error = %v", err)
	}

	if len(all) != 2 {
		t.Errorf("currentHoldings() len = %v, want %v", len(all), 2)
	}
}

func TestConfig_getHoldingsSlice(t *testing.T) {
	slice := testApp.getHoldingsSlice()

	if len(slice) != 3 {
		t.Errorf("getHoldingsSlice() len = %v, want %v", len(slice), 3)
	}
}
