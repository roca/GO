package main

import (
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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

	// Set Theme
	a.Settings().SetTheme(&myTheme{})

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
	saveAsMenuItem := fyne.NewMenuItem("Save as...", app.saveAsFunc(win))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)

}

var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func (app *config) openFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				// user cancelled
				return
			}
			// read the file
			defer reader.Close()
			data, err := ioutil.ReadAll(reader)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			app.EditWidget.SetText(string(data))

			app.CurrentFile = reader.URI()
			win.SetTitle(win.Title() + ": " + app.CurrentFile.Name())
			app.SaveMenuItem.Disabled = false
		}, win)

		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (app *config) saveFunc(win fyne.Window) func() {
	return func() {
		if app.CurrentFile != nil {
			writer, err := storage.Writer(app.CurrentFile)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			defer writer.Close()
			writer.Write([]byte(app.EditWidget.Text))
		}
	}
}

func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveAsDiaglog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if writer == nil {
				// user cancelled
				return
			}

			if !strings.HasSuffix(strings.ToLower(writer.URI().String()), ".md") {
				dialog.ShowInformation("Invalid file extension", "Please use a .md file extension", win)
				return
			}

			// save the file
			defer writer.Close()
			writer.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = writer.URI()

			win.SetTitle(win.Title() + ": " + app.CurrentFile.Name())
			app.SaveMenuItem.Disabled = false
		}, win)
		saveAsDiaglog.SetFileName("Untitled.md")
		saveAsDiaglog.SetFilter(filter)
		saveAsDiaglog.Show()
	}
}
