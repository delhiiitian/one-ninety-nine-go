// Chapter 5.1: Errors are values
// Run: go run ./ch05-error-handling/01-errors-are-values
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Many standard functions return (result, error).
	inputs := []string{"250", "12.5", "abc", ""}
	for _, in := range inputs {
		n, err := strconv.Atoi(in)
		if err != nil {
			fmt.Printf("%-6q -> problem: %v\n", in, err)
			continue
		}
		fmt.Printf("%-6q -> ok: %d\n", in, n)
	}

	// Opening a file that does not exist
	_, err := os.Open("no-such-file.txt")
	if err != nil {
		fmt.Println("open failed:", err)
	}

	// error is just an interface with one method: Error() string
	var e error = err
	fmt.Println("as text:", e.Error())

	// nil means "no error"
	_, err = strconv.ParseBool("true")
	fmt.Println("error is nil?", err == nil)
}
