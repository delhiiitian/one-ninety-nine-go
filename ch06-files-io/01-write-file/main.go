// Chapter 6.1: Writing a file
// Run: go run ./ch06-files-io/01-write-file
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := os.TempDir() // a safe scratch folder on every OS
	path := filepath.Join(dir, "shopping-list.txt")

	// 1. The quick way: write the whole thing at once.
	list := "rice\ndal\nmilk\n"
	if err := os.WriteFile(path, []byte(list), 0o644); err != nil {
		fmt.Println("could not write:", err)
		return
	}
	fmt.Println("wrote", len(list), "bytes to", path)

	// 2. The step-by-step way: create, write, close.
	receipt := filepath.Join(dir, "receipt.txt")
	f, err := os.Create(receipt) // creates or empties the file
	if err != nil {
		fmt.Println("could not create:", err)
		return
	}
	defer f.Close()

	// bufio.Writer collects small writes and sends them in big chunks (faster).
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, "%-10s %6s\n", "Item", "Price")
	for _, item := range []struct {
		name  string
		price float64
	}{{"Tea", 120}, {"Biscuits", 45.5}} {
		fmt.Fprintf(w, "%-10s %6.2f\n", item.name, item.price)
	}
	if err := w.Flush(); err != nil { // don't forget: Flush sends what is left
		fmt.Println("could not flush:", err)
		return
	}
	fmt.Println("wrote receipt to", receipt)

	data, _ := os.ReadFile(receipt)
	fmt.Print(string(data))
}
