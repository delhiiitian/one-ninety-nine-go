// Project 3: Password Generator and Strength Checker
// Run: go run ./projects/03-password-tool
// Or:  printf "1\n16\n2\nsunshine\n2\nT!ger-Moon-42-river\n3\n" | go run ./projects/03-password-tool
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	ask := func(q string) (string, bool) { // a closure that captures "in"
		fmt.Print(q)
		if !in.Scan() {
			return "", false
		}
		return strings.TrimSpace(in.Text()), true
	}

	fmt.Println("=== PASSWORD TOOL ===")
	for {
		fmt.Println("\n1) Generate a password  2) Check a password  3) Quit")
		choice, ok := ask("Choose: ")
		if !ok || choice == "3" {
			fmt.Println("Stay safe!")
			return
		}
		switch choice {
		case "1":
			text, _ := ask("Length (8-64, Enter for 16): ")
			length, err := strconv.Atoi(text)
			if err != nil {
				length = 16
			}
			pw, err := Generate(length, WithLower, WithUpper, WithDigits, WithSymbols)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Your password:", pw)
			report(pw)
		case "2":
			pw, _ := ask("Password to check: ")
			report(pw)
		default:
			fmt.Println("Please choose 1, 2 or 3.")
		}
	}
}

func report(pw string) {
	score, tips := Check(pw)
	fmt.Printf("Strength: %s (%d/%d)\n", Label(score), score, len(Rules))
	for _, t := range tips {
		fmt.Println("  tip:", t)
	}
}
