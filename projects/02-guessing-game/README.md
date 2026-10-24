# Project 2: Number Guessing Game

The computer picks a secret number and you guess it. After each guess you get a hint
(cold, warm, boiling hot) and a direction. Three difficulty levels, a limited number of
tries, play-again, and a best-score table.

Book: *Go in 199 Pages*, Chapter 2 project.

## Run it

```bash
go run ./projects/02-guessing-game            # random numbers
go run ./projects/02-guessing-game -seed 7    # same numbers every time (good for demos)
go test ./projects/02-guessing-game
```

## Files

| File | Job |
| --- | --- |
| `main.go` | The game loop: menus, reading guesses, play again, best scores |
| `game.go` | The rules: `Level`, `Game`, `Guess` hints, and the seeded number `Picker` |
| `game_test.go` | Tests for the hints, losing, and repeatable seeds |

## Ideas to extend

- Add a "Custom" level where the player chooses the range.
- Let two players take turns and keep separate scores.
- Show the remaining range ("between 35 and 50") after each guess.
