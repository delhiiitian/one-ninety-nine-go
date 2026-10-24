package main

import (
	"fmt"
	"math/rand/v2"
)

// Level describes one difficulty setting.
type Level struct {
	Name  string
	Max   int
	Tries int
}

// Levels are listed in menu order.
var Levels = []Level{
	{"Easy", 50, 10},
	{"Medium", 100, 7},
	{"Hard", 500, 9},
}

// Game holds the state of one round.
type Game struct {
	Level  Level
	Secret int
	Used   int
	Won    bool
}

func NewGame(level Level, secret int) *Game {
	return &Game{Level: level, Secret: secret}
}

// Over reports whether the round has finished.
func (g *Game) Over() bool {
	return g.Won || g.Used >= g.Level.Tries
}

// Guess records one try and returns a hint for the player.
func (g *Game) Guess(n int) string {
	g.Used++
	diff := n - g.Secret
	if diff < 0 {
		diff = -diff
	}

	direction := "higher"
	if n > g.Secret {
		direction = "lower"
	}

	switch {
	case diff == 0:
		g.Won = true
		return fmt.Sprintf("Correct! You got it in %d tries.", g.Used)
	case diff <= 3:
		return "Boiling hot! Go " + direction + "."
	case diff <= 10:
		return "Warm. Go " + direction + "."
	default:
		return "Cold. Go " + direction + "."
	}
}

// Picker chooses secret numbers. A fixed seed makes games repeatable.
type Picker struct{ r *rand.Rand }

func NewPicker(seed uint64) Picker {
	if seed == 0 {
		seed = rand.Uint64()
	}
	return Picker{rand.New(rand.NewPCG(seed, seed))}
}

func (p Picker) Pick(max int) int { return p.r.IntN(max) + 1 }
