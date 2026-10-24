// Chapter 1.6: Strings
// Run: go run ./ch01-basics/06-strings
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	first := "Ravi"
	last := "Kumar"

	// Join with +
	full := first + " " + last
	fmt.Println("full name:", full)

	// Raw strings use backticks: no escapes, can span lines.
	address := `12 MG Road
Pune 411001`
	fmt.Println(address)

	// The strings package is your toolbox.
	email := "  Ravi.Kumar@Example.com  "
	clean := strings.ToLower(strings.TrimSpace(email))
	fmt.Println("clean email:", clean)
	fmt.Println("contains @:", strings.Contains(clean, "@"))
	fmt.Println("domain:", clean[strings.Index(clean, "@")+1:])
	fmt.Println("parts:", strings.Split("tea,milk,bread", ","))
	fmt.Println("shout:", strings.ToUpper("sale today"))
	fmt.Println("replace:", strings.ReplaceAll("2026-10-24", "-", "/"))
	fmt.Println("repeat:", strings.Repeat("=", 20))

	// Length in bytes vs characters
	word := "café"
	fmt.Println("bytes:", len(word), "characters:", utf8.RuneCountInString(word))

	// Loop over characters (runes), not bytes.
	for i, r := range "Go♥" {
		fmt.Printf("index %d: %c\n", i, r)
	}

	// Strings are immutable: build new ones efficiently with a Builder.
	var sb strings.Builder
	for i := 1; i <= 3; i++ {
		fmt.Fprintf(&sb, "item%d ", i)
	}
	fmt.Println("built:", sb.String())
}
