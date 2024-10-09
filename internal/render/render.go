package render

import (
	"github.com/jdeeline/currency-monitor/internal/grabber"
	"github.com/rodaine/table"
)

// Displays a table with the specified currencies.
func Table(rows []grabber.Currency) {
	tbl := table.New("Code", "Name", "Nominal", "Rate", "Previous Rate")

	for _, c := range rows {
		tbl.AddRow(c.Code, c.Name, c.Nominal, c.Rate, c.PreviousRate)
	}

	tbl.Print()
}
