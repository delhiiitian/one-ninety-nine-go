// Chapter 4.4: Maps (look things up by key)
// Run: go run ./ch04-data-structures/04-maps
package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func main() {
	// A phone book: name -> number
	phone := map[string]string{
		"Mum":    "98450 11111",
		"Office": "080 2222 3333",
	}
	phone["Doctor"] = "98860 44444" // add
	phone["Mum"] = "98450 12345"    // update
	fmt.Println("Mum:", phone["Mum"])

	// Missing keys return the zero value; use ", ok" to tell the difference.
	if num, ok := phone["Plumber"]; ok {
		fmt.Println("Plumber:", num)
	} else {
		fmt.Println("No plumber saved")
	}

	delete(phone, "Office")
	fmt.Println("contacts:", len(phone))

	// Counting words: the classic map job
	text := "the cat sat on the mat the end"
	counts := map[string]int{}
	for _, w := range strings.Fields(text) {
		counts[w]++ // missing key starts at 0
	}

	// Maps have no order. Sort the keys to print predictably.
	keys := slices.Sorted(maps.Keys(counts))
	for _, k := range keys {
		fmt.Printf("%-4s %d\n", k, counts[k])
	}

	// Map of slices: students per class
	classes := map[string][]string{}
	classes["5A"] = append(classes["5A"], "Ira", "Kabir")
	classes["5B"] = append(classes["5B"], "Zoya")
	fmt.Println("5A:", classes["5A"], "5B:", classes["5B"])
}
