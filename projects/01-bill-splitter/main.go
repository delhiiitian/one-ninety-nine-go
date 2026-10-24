// Project 1: Restaurant Bill Splitter
// Run: go run ./projects/01-bill-splitter
// Or:  printf "2450\n10\n4\n" | go run ./projects/01-bill-splitter
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println(strings.Repeat("=", 32))
	fmt.Println("      RESTAURANT BILL SPLITTER")
	fmt.Println(strings.Repeat("=", 32))

	amount := askFloat(in, "Bill amount before tax (Rs): ", 0)
	tipPercent := askFloat(in, "Tip percent (e.g. 10): ", DefaultTipPercent)
	people := int(askFloat(in, "Number of people: ", 1))
	if people < 1 {
		people = 1
	}

	b := Calculate(amount, tipPercent, people)
	printReceipt(b)
}

// askFloat prints a question, reads one line and turns it into a number.
// Empty or bad input falls back to the given default.
func askFloat(in *bufio.Reader, question string, fallback float64) float64 {
	fmt.Print(question)
	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(line)
	value, err := strconv.ParseFloat(line, 64)
	if err != nil || value < 0 {
		fmt.Printf("  (using %.0f)\n", fallback)
		return fallback
	}
	return value
}

func printReceipt(b Bill) {
	fmt.Println()
	fmt.Println(strings.Repeat("-", 32))
	fmt.Printf("%-18s %12.2f\n", "Subtotal", b.Subtotal)
	fmt.Printf("%-18s %12.2f\n", fmt.Sprintf("GST (%.0f%%)", TaxRate*100), b.Tax)
	fmt.Printf("%-18s %12.2f\n", fmt.Sprintf("Tip (%.0f%%)", b.TipPercent), b.Tip)
	fmt.Println(strings.Repeat("-", 32))
	fmt.Printf("%-18s %12.2f\n", "TOTAL", b.Total)
	fmt.Printf("%-18s %12d\n", "People", b.People)
	fmt.Printf("%-18s %12.2f\n", "Each person pays", b.PerPerson)
	fmt.Println(strings.Repeat("=", 32))
}
