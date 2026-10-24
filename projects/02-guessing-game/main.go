// Project 2: Number Guessing Game
// Run: go run ./projects/02-guessing-game
// Repeatable game for demos and tests: go run ./projects/02-guessing-game -seed 7
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	seed := flag.Uint64("seed", 0, "fix the secret numbers (0 = random each time)")
	flag.Parse()

	in := bufio.NewScanner(os.Stdin)
	picker := NewPicker(*seed)
	best := map[string]int{} // fewest tries per level

	fmt.Println("=== GUESS THE NUMBER ===")
	for {
		level, ok := chooseLevel(in)
		if !ok {
			break
		}
		game := NewGame(level, picker.Pick(level.Max))
		fmt.Printf("\nI'm thinking of a number from 1 to %d. You have %d tries.\n", level.Max, level.Tries)

		for !game.Over() {
			fmt.Printf("Try %d/%d: ", game.Used+1, level.Tries)
			if !in.Scan() {
				fmt.Println("\nBye!")
				return
			}
			guess, err := strconv.Atoi(strings.TrimSpace(in.Text()))
			if err != nil || guess < 1 || guess > level.Max {
				fmt.Printf("Please type a whole number from 1 to %d.\n", level.Max)
				continue // a bad entry does not cost a try
			}
			fmt.Println(game.Guess(guess))
		}

		if game.Won {
			if b, seen := best[level.Name]; !seen || game.Used < b {
				best[level.Name] = game.Used
				fmt.Println("New best score for", level.Name+"!")
			}
		} else {
			fmt.Println("Out of tries! The number was", game.Secret)
		}

		fmt.Print("\nPlay again? (y/n): ")
		if !in.Scan() || strings.ToLower(strings.TrimSpace(in.Text())) != "y" {
			break
		}
	}

	fmt.Println("\nBest scores:")
	for _, l := range Levels {
		if b, seen := best[l.Name]; seen {
			fmt.Printf("  %-6s %d tries\n", l.Name, b)
		}
	}
	fmt.Println("Thanks for playing!")
}

// chooseLevel keeps asking until the player picks 1, 2 or 3, or input ends.
func chooseLevel(in *bufio.Scanner) (Level, bool) {
	for {
		fmt.Println("\nChoose a level:")
		for i, l := range Levels {
			fmt.Printf("  %d) %-6s 1-%d, %d tries\n", i+1, l.Name, l.Max, l.Tries)
		}
		fmt.Print("Level: ")
		if !in.Scan() {
			return Level{}, false
		}
		switch strings.TrimSpace(in.Text()) {
		case "1":
			return Levels[0], true
		case "2":
			return Levels[1], true
		case "3":
			return Levels[2], true
		default:
			fmt.Println("Please type 1, 2 or 3.")
		}
	}
}
