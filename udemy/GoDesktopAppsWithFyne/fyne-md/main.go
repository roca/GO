package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

// makeUI creates two widgets, assigns them to the app config, and
// adds a listener on the edit widget that updates the preview widget
// with parsed markdown whenever the user types in the edit widget.
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

	openMenuItem := fyne.NewMenuItem("Open...", app.openFunc(win))
	saveMenuItem := fyne.NewMenuItem("Save", app.saveFunc(win))
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true
	saveAsMenuItem := fyne.NewMenuItem("Save...", app.saveAsFunc(win))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)

}

func (app *config) openFunc(win fyne.Window) func() {
	return func() {}
}
func (app *config) saveFunc(win fyne.Window) func() {
	return func() {}
}
func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDiaglog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if write == nil {
				// user cancelled
				return
			}
			// save the file
			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + "_" + app.CurrentFile.Name())
			app.SaveMenuItem.Disabled = false
		}, win)
		saveDiaglog.Show()
	}
}
