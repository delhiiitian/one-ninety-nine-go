// Chapter 3.7: Recursion (a function that calls itself)
// Run: go run ./ch03-functions/07-recursion
package main

import (
	"fmt"
	"strings"
)

// factorial: 5! = 5 * 4!  ...  1! = 1 (the base case stops the recursion)
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Folder is a tree: folders can contain folders.
type Folder struct {
	Name    string
	SizeKB  int
	Folders []Folder
}

// totalSize walks the whole tree: perfect for recursion.
func totalSize(f Folder) int {
	size := f.SizeKB
	for _, sub := range f.Folders {
		size += totalSize(sub)
	}
	return size
}

func printTree(f Folder, depth int) {
	fmt.Printf("%s%s (%d KB)\n", strings.Repeat("  ", depth), f.Name, f.SizeKB)
	for _, sub := range f.Folders {
		printTree(sub, depth+1)
	}
}

func main() {
	fmt.Println("5! =", factorial(5))
	fmt.Println("10! =", factorial(10))

	home := Folder{"home", 4, []Folder{
		{"photos", 2048, []Folder{{"2025", 900, nil}, {"2026", 1200, nil}}},
		{"docs", 300, []Folder{{"tax", 120, nil}}},
	}}
	printTree(home, 0)
	fmt.Println("Total:", totalSize(home), "KB")
}
