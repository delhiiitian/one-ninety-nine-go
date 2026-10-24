// Chapter 6.7: CSV files (spreadsheets as text)
// Run: go run ./ch06-files-io/07-csv
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	path := filepath.Join(os.TempDir(), "marks.csv")
	defer os.Remove(path)

	// Write
	rows := [][]string{
		{"name", "maths", "science", "english"},
		{"Aarav", "88", "92", "75"},
		{"Diya", "95", "89", "91"},
		{"Kabir, Jr.", "67", "72", "80"}, // a comma inside a value: csv quotes it for you
	}
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := csv.NewWriter(f)
	w.WriteAll(rows) // writes and flushes
	f.Close()
	if err := w.Error(); err != nil {
		fmt.Println(err)
		return
	}

	raw, _ := os.ReadFile(path)
	fmt.Print("file contents:\n", string(raw), "\n")

	// Read
	f, _ = os.Open(path)
	defer f.Close()
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	header := records[0]
	fmt.Printf("%-12s %s\n", "student", "average")
	for _, rec := range records[1:] {
		sum := 0
		for _, cell := range rec[1:] {
			n, err := strconv.Atoi(cell)
			if err != nil {
				fmt.Println("bad number:", cell)
				continue
			}
			sum += n
		}
		fmt.Printf("%-12s %.1f\n", rec[0], float64(sum)/float64(len(header)-1))
	}
}
