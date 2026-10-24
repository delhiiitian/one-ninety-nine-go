// Chapter 2.8: break, continue and labels
// Run: go run ./ch02-control-flow/08-break-continue
package main

import "fmt"

func main() {
	// continue: skip the rest of this turn, go to the next
	orders := []int{250, -1, 400, 0, 120}
	for _, amount := range orders {
		if amount <= 0 {
			fmt.Println("skipping invalid order", amount)
			continue
		}
		fmt.Println("processing order of Rs", amount)
	}

	// break: leave the loop completely
	seats := []string{"taken", "taken", "free", "free"}
	for i, s := range seats {
		if s == "free" {
			fmt.Println("first free seat:", i+1)
			break
		}
	}

	// A label lets break leave an outer loop from inside an inner one.
	shelves := [][]string{
		{"soap", "shampoo"},
		{"biscuits", "tea", "sugar"},
		{"pens"},
	}
	target := "tea"
search:
	for shelf, items := range shelves {
		for slot, item := range items {
			if item == target {
				fmt.Printf("found %s on shelf %d slot %d\n", target, shelf+1, slot+1)
				break search
			}
		}
	}
}
