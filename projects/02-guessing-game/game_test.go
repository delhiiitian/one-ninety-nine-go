package main

import (
	"strings"
	"testing"
)

func TestGuessHints(t *testing.T) {
	g := NewGame(Levels[1], 40)
	cases := []struct {
		guess int
		want  string
	}{
		{10, "Cold. Go higher."},
		{90, "Cold. Go lower."},
		{35, "Warm. Go higher."},
		{42, "Boiling hot! Go lower."},
	}
	for _, c := range cases {
		if got := g.Guess(c.guess); got != c.want {
			t.Errorf("Guess(%d) = %q, want %q", c.guess, got, c.want)
		}
	}
	if got := g.Guess(40); !strings.HasPrefix(got, "Correct") || !g.Won || !g.Over() {
		t.Fatalf("winning guess not recognised: %q", got)
	}
}

func TestRunOutOfTries(t *testing.T) {
	g := NewGame(Level{"Tiny", 10, 2}, 5)
	g.Guess(1)
	g.Guess(2)
	if !g.Over() || g.Won {
		t.Fatal("game should be over and lost")
	}
}

func TestSeedIsRepeatable(t *testing.T) {
	a, b := NewPicker(7), NewPicker(7)
	for range 5 {
		if a.Pick(100) != b.Pick(100) {
			t.Fatal("same seed gave different numbers")
		}
	}
}
