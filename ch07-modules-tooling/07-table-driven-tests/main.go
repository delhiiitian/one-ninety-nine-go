// Chapter 7.7: Table-driven tests
// Run the program: go run ./ch07-modules-tooling/07-table-driven-tests
// Run the tests:   go test -v ./ch07-modules-tooling/07-table-driven-tests
// Run one case:    go test -v -run 'TestShippingCost/heavy' ./ch07-modules-tooling/07-table-driven-tests
package main

import "fmt"

func main() {
	for _, w := range []float64{0.5, 4, 12, 40} {
		cost, err := ShippingCost(w, 500)
		if err != nil {
			fmt.Printf("%5.1f kg: %v\n", w, err)
			continue
		}
		fmt.Printf("%5.1f kg: Rs %.0f\n", w, cost)
	}
}
