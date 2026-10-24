// Chapter 1.8: Formatted output with fmt
// Run: go run ./ch01-basics/08-formatted-output
package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func main() {
	name := "Notebook"
	qty := 3
	price := 45.5

	// Printf uses "verbs" that start with %.
	fmt.Printf("%s x %d = %.2f\n", name, qty, float64(qty)*price)

	// Width and alignment: a neat receipt.
	fmt.Printf("%-10s|%5s|%8s\n", "Item", "Qty", "Total")
	fmt.Printf("%-10s|%5d|%8.2f\n", "Pen", 10, 50.0)
	fmt.Printf("%-10s|%5d|%8.2f\n", "Notebook", 3, 136.5)

	// Handy verbs
	p := Product{"Lamp", 899}
	fmt.Printf("%%v  -> %v\n", p)  // default format
	fmt.Printf("%%+v -> %+v\n", p) // with field names
	fmt.Printf("%%T  -> %T\n", p)  // the type
	fmt.Printf("%%q  -> %q\n", name)
	fmt.Printf("%%t  -> %t\n", qty > 2)
	fmt.Printf("%%05d -> %05d\n", 42) // pad with zeros
	fmt.Printf("%%x  -> %x\n", 255)   // hexadecimal
	fmt.Printf("%%b  -> %b\n", 5)     // binary

	// Sprintf returns the text instead of printing it.
	label := fmt.Sprintf("%s (Rs %.0f)", p.Name, p.Price)
	fmt.Println("label:", label)
}
