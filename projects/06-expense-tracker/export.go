package main

import (
	"encoding/csv"
	"io"
	"strconv"
)

// WriteCSV writes expenses to any io.Writer: a file, the screen, or a web response.
func WriteCSV(w io.Writer, list []Expense) error {
	cw := csv.NewWriter(w)
	cw.Write([]string{"id", "date", "amount", "category", "note"})
	for _, e := range list {
		cw.Write([]string{
			strconv.Itoa(e.ID),
			e.Date,
			strconv.FormatFloat(e.Amount, 'f', 2, 64),
			e.Category,
			e.Note,
		})
	}
	cw.Flush()
	return cw.Error()
}
