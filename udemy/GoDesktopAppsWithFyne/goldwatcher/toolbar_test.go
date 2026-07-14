package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestApp_getToolBar(t *testing.T) {
	toolBar := testApp.getToolBar()
	if toolBar == nil {
		t.Error("Expected a toolbar, got nil")
	}
	if len(toolBar.Items) != 4 {
		t.Errorf("Expected %d items, got %d", 4, len(toolBar.Items))
	}
	
}

func TestApp_addHoldingsDialog(t *testing.T) {
	testApp.addHoldingsDialog()
	
	test.Type(testApp.AddHoldingsPurchaseAmountEntry, "1")
	test.Type(testApp.AddHoldingsPurchasePriceEntry, "1000")
	test.Type(testApp.AddHoldingsPurchaseDateEntry, "2021-01-01")

	if testApp.AddHoldingsPurchaseDateEntry.Text != "2021-01-01" {
		t.Errorf("Expected date %s, got %s", "2021-01-01", testApp.AddHoldingsPurchaseDateEntry.Text)
	}

	if testApp.AddHoldingsPurchaseAmountEntry.Text != "1" {
		t.Errorf("Expected amount %s, got %s", "1", testApp.AddHoldingsPurchaseAmountEntry.Text)
	}

	if testApp.AddHoldingsPurchasePriceEntry.Text != "1000" {
		t.Errorf("Expected price %s, got %s", "1000", testApp.AddHoldingsPurchasePriceEntry.Text)
	}

}
