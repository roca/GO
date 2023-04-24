package main

import "testing"

func TestApp_getToolBar(t *testing.T) {
	toolBar := testApp.getToolBar()
	if toolBar == nil {
		t.Error("Expected a toolbar, got nil")
	}
	if len(toolBar.Items) != 4 {
		t.Errorf("Expected %d items, got %d", 4, len(toolBar.Items))
	}
	
}
