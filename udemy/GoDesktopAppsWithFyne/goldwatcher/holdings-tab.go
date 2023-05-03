package main

import (
	"fmt"
	"goldwatcher/repository"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) holdingsTab() *fyne.Container {
	return nil
}

func (app *Config) getholdingsTable() *widget.Table {
	return nil
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