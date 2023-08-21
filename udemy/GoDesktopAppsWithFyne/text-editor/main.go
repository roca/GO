package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TestEdit")

	ui := makeUI()
	w.SetContent(ui)

	w.Resize(fyne.NewSize(480, 360))
	w.ShowAndRun()
}

func makeUI() fyne.CanvasObject {
	entry := widget.NewMultiLineEntry()
	cursorRow := widget.NewLabel("1")
	cursorCol := widget.NewLabel("1")

	toolbar := buildToolbar()

	status := container.NewHBox(
		layout.NewSpacer(),
		widget.NewLabel("Cursor Row: "), cursorRow,
		widget.NewLabel("Col:"), cursorCol,
	)

	return container.NewBorder(
		toolbar,
		status,
		nil,
		nil,
		container.NewScroll(entry),
	)
}

func buildToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), open),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), save),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), cut),
		widget.NewToolbarAction(theme.ContentCopyIcon(), copy),
		widget.NewToolbarAction(theme.ContentPasteIcon(), paste),
	)
}

func open()  {}
func save()  {}
func cut()   {}
func copy()  {}
func paste() {}
