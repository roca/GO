package main

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
)

func Test_makeUI(t *testing.T) {
	var testCfg config

	edit, preview := testCfg.makeUI()

	test.Type(edit, "Hello, World!")

	if preview.String() != "Hello, World!" {
		t.Errorf("preview Text was '%s'; want 'Hello, World!'", preview.String())
	}
}

func Test_RunApp(t *testing.T) {
	var testCfg config
	testApp := test.NewApp()
	testWin := testApp.NewWindow("Test MarkDown")

	edit, preview := testCfg.makeUI()

	testCfg.createMenuItems(testWin)

	testWin.SetContent(container.NewHSplit(edit, preview))

	testApp.Run()

	test.Type(edit, "Some text")

	if preview.String() != "Some text" {
		t.Errorf("preview Text was '%s'; want 'Some text'", preview.String())
	}

}
