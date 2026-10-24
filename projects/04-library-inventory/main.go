// Project 4: Library Inventory Manager
// Run: go run ./projects/04-library-inventory
// Type "help" to see the commands.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lib := NewLibrary()
	seed(lib)

	fmt.Println("=== CITY LIBRARY ===  (type 'help')")
	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !in.Scan() {
			fmt.Println()
			return
		}
		line := strings.TrimSpace(in.Text())
		if line == "" {
			continue
		}
		if line == "quit" || line == "exit" {
			fmt.Println("Goodbye!")
			return
		}
		fmt.Println(Run(lib, line))
	}
}

// seed fills the library with a few books so there is something to try.
func seed(lib *Library) {
	lib.Add("The Pragmatic Programmer", "David Thomas", 1999, "programming", 2)
	lib.Add("Wings of Fire", "A. P. J. Abdul Kalam", 1999, "biography", 3)
	lib.Add("The Go Programming Language", "Alan Donovan", 2015, "programming", 1)
	lib.Add("Malgudi Days", "R. K. Narayan", 1943, "fiction", 2)
	lib.Add("Sapiens", "Yuval Noah Harari", 2011, "history", 1)
}
