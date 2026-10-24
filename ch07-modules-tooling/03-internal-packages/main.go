// Chapter 7.3: internal packages and project layout
// Run: go run ./ch07-modules-tooling/03-internal-packages
//
// A typical Go project layout:
//
//	myshop/
//	├── go.mod
//	├── main.go            (or cmd/myshop/main.go when there are several programs)
//	├── internal/          private code: cannot be imported by other modules
//	│   ├── pricing/
//	│   └── storage/
//	└── pkg/ or top-level  public packages others may import (optional)
package main

import (
	"fmt"

	"github.com/delhiiitian/one-ninety-nine-go/ch07-modules-tooling/03-internal-packages/internal/pricing"
)

func main() {
	cart := []struct {
		item, category string
		price          float64
	}{
		{"Rice 5kg", "food", 450},
		{"Earphones", "electronics", 1299},
		{"Go book", "books", 199},
	}
	total := 0.0
	for _, c := range cart {
		p := pricing.WithTax(c.price, c.category)
		total += p
		fmt.Printf("%-10s %8.2f\n", c.item, p)
	}
	fmt.Printf("%-10s %8.2f\n", "TOTAL", total)
}
