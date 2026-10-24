// Chapter 2.2: if with a short statement
// Run: go run ./ch02-control-flow/02-if-short-statement
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// The variable lives only inside the if/else: tidy and safe.
	if pin, err := strconv.Atoi("4821"); err != nil {
		fmt.Println("PIN must be digits:", err)
	} else {
		fmt.Println("PIN accepted, last digit:", pin%10)
	}

	if _, err := strconv.Atoi("48a1"); err != nil {
		fmt.Println("Rejected:", err)
	}

	// Very common with maps (Chapter 4): look up and test in one line.
	stock := map[string]int{"apple": 12, "mango": 0}
	if qty, found := stock["mango"]; !found {
		fmt.Println("We don't sell mangoes.")
	} else if qty == 0 {
		fmt.Println("Mangoes are out of stock.")
	} else {
		fmt.Println("Mangoes left:", qty)
	}

	if qty, found := stock["kiwi"]; found {
		fmt.Println("Kiwis left:", qty)
	} else {
		fmt.Println("We don't sell kiwis.")
	}
	// fmt.Println(qty) // would not compile: qty only exists inside the if
}
