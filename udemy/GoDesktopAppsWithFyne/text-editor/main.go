package main

import (
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var window fyne.Window
var entry *widget.Entry
var currentFile fyne.URI
var cursorRow *widget.Label
var cursorCol *widget.Label

func main() {
	a := app.New()
	window = a.NewWindow("TestEdit")

	ui := makeUI()
	window.SetContent(ui)

	window.Resize(fyne.NewSize(480, 360))
	window.ShowAndRun()
}

func makeUI() fyne.CanvasObject {
	entry = widget.NewMultiLineEntry()
	cursorRow = widget.NewLabel("1")
	cursorCol = widget.NewLabel("1")

	toolbar := buildToolbar()

	status := container.NewHBox(
		layout.NewSpacer(),
		widget.NewLabel("Cursor Row: "), cursorRow,
		widget.NewLabel("Col:"), cursorCol,
	)

	entry.OnCursorChanged = updateStatus

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

func open() {
	dialog.ShowFileOpen(func(r fyne.URIReadCloser, err error) {
		defer r.Close()
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		if r == nil {
			return
		}
		data, err := io.ReadAll(r)
		if err == nil {
			entry.SetText(string(data))
			currentFile = r.URI()
		} else {
			dialog.ShowError(err, window)
		}
	}, window)
}
func save() {
	if currentFile != nil {
		w, err := storage.Writer(currentFile)
		defer w.Close()
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		_, err = w.Write([]byte(entry.Text))
		if err != nil {
			dialog.ShowError(err, window)
		}
	}
}
func cut()   {}
func copy()  {}
func paste() {}
func updateStatus() {
	cursorRow.SetText(fmt.Sprintf("%d", entry.CursorRow+1))
	cursorCol.SetText(fmt.Sprintf("%d", entry.CursorColumn+1))
}
