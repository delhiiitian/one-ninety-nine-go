// Chapter 4.9: Sorting
// Run: go run ./ch04-data-structures/09-sorting
package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

type Product struct {
	Name   string
	Price  float64
	Rating float64
}

func main() {
	// Simple values
	nums := []int{42, 7, 19, 3, 88}
	slices.Sort(nums)
	fmt.Println("ascending:", nums)

	names := []string{"Zara", "amit", "Bina", "Chirag"}
	slices.Sort(names) // uppercase sorts before lowercase!
	fmt.Println("names:", names)

	// Structs: say how to compare two items
	shop := []Product{
		{"Headphones", 1999, 4.3},
		{"Charger", 499, 4.6},
		{"Speaker", 2999, 4.3},
		{"Cable", 199, 3.9},
	}

	// Cheapest first
	slices.SortFunc(shop, func(a, b Product) int {
		return cmp.Compare(a.Price, b.Price)
	})
	fmt.Println("\nby price:")
	for _, p := range shop {
		fmt.Printf("  %-10s %6.0f\n", p.Name, p.Price)
	}

	// Best rated first; ties broken by lower price
	slices.SortFunc(shop, func(a, b Product) int {
		if c := cmp.Compare(b.Rating, a.Rating); c != 0 {
			return c
		}
		return cmp.Compare(a.Price, b.Price)
	})
	fmt.Println("\nby rating, then price:")
	for _, p := range shop {
		fmt.Printf("  %-10s %.1f %6.0f\n", p.Name, p.Rating, p.Price)
	}

	// Binary search on sorted data is very fast
	i, found := slices.BinarySearch(nums, 19)
	fmt.Println("\n19 found:", found, "at index", i)

	// The older sort package still appears in a lot of code
	sort.Slice(shop, func(i, j int) bool { return shop[i].Name < shop[j].Name })
	fmt.Println("by name:", shop[0].Name, shop[1].Name, shop[2].Name, shop[3].Name)
}
