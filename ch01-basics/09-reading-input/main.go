// Chapter 1.9: Reading input from the keyboard
// Run: go run ./ch01-basics/09-reading-input
// Or feed input: printf "Meera\n28\n" | go run ./ch01-basics/09-reading-input
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // remove the Enter key (newline)
	if name == "" {
		name = "friend"
	}

	fmt.Print("How old are you? ")
	ageText, _ := reader.ReadString('\n')
	age, err := strconv.Atoi(strings.TrimSpace(ageText))
	if err != nil {
		fmt.Println("\nThat was not a number, so I'll guess 18.")
		age = 18
	}

	fmt.Printf("\nHello, %s! In 10 years you will be %d.\n", name, age+10)
}
