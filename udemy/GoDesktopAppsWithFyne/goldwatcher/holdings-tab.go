package main

import (
	"fmt"
	"goldwatcher/repository"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) holdingsTab() *fyne.Container {
	app.Holdings = app.getHoldingsSlice()

	app.HoldingsTable = app.getholdingsTable()

	holdingsContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, app.HoldingsTable),
	)

	return holdingsContainer
}

func (app *Config) getholdingsTable() *widget.Table {

	t := widget.NewTable(
		func() (int, int) {
			return len(app.Holdings), len(app.Holdings[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			if i.Col == (len(app.Holdings[0])-1) && i.Row != 0 {
				// last cell is a button
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete", "Are you sure you want to delete this holding?", func(deleted bool) {
						if deleted {
							id, _ := strconv.Atoi(app.Holdings[i.Row][0].(string))
							err := app.DB.DeleteHolding(int64(id))
							if err != nil {
								app.ErrorLog.Println(err)
							}
						}
						// refresh holdings table
						app.refreshHoldingsTable()
					}, app.MainWindow)
				})
				w.Importance = widget.HighImportance
				obj.(*fyne.Container).Objects = []fyne.CanvasObject{w}
			} else {
				// we're just putting i textual data
				obj.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(app.Holdings[i.Row][i.Col].(string)),
				}
			}
		})

	colWidths := []float32{50, 200, 200, 200, 110}
	for c, w := range colWidths {
		t.SetColumnWidth(c, w)
	}

	return t
}

func (app *Config) getHoldingsSlice() [][]interface{} {
	var slice [][]interface{}

	holdings, err := app.currentHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil
	}
	// Table headers
	slice = append(slice, []interface{}{"ID", "Amount", "Price", "Date", "Delete?"})

	for _, h := range holdings {
		var currentRow []interface{}
		currentRow = append(currentRow, strconv.FormatInt(h.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%.2f", h.Amount))
		currentRow = append(currentRow, fmt.Sprintf("%.2f", float32(h.PurchasePrice)/100))
		currentRow = append(currentRow, h.PurchaseDate.Format("2006-01-02"))
		// Last column is a button to delete the row
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}

	return slice
}

func (app *Config) currentHoldings() ([]repository.Holdings, error) {
	holdings, err := app.DB.AllHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	return holdings, nil
}
