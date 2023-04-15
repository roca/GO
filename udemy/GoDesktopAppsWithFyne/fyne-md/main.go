package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {
	// create a fyne app
	a := app.New()

	// create a window for the app
	win := a.NewWindow("Markdown Editor")

	// get the user interface
	edit, preview := cfg.makeUI()
	cfg.createMenuItems(win)

	// set the content of the window
	win.SetContent(container.NewHSplit(edit, preview))
	

	// show window and run app
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	// create widgets
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")

	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	// return widgets
	return edit, preview
}

func (app *config) createMenuItems(win fyne.Window) {

	openMenutItem := fyne.NewMenuItem("Open...", func() {})
	saveMenutItem := fyne.NewMenuItem("Save", func() {})
	saveAsMenutItem := fyne.NewMenuItem("Save As...", func() {})

	menu := fyne.NewMenu("File", openMenutItem, saveMenutItem, saveAsMenutItem)

	win.SetMainMenu(fyne.NewMainMenu(menu))
	
}
