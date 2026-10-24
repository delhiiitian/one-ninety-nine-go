package main

import (
	"strings"
	"testing"
)

func TestGenerateUsesEverySet(t *testing.T) {
	for range 50 {
		pw, err := Generate(12, WithLower, WithUpper, WithDigits, WithSymbols)
		if err != nil {
			t.Fatal(err)
		}
		if len(pw) != 12 {
			t.Fatalf("length %d, want 12", len(pw))
		}
		for _, set := range []string{lower, upper, digits, symbols} {
			if !strings.ContainsAny(pw, set) {
				t.Fatalf("%q has nothing from %q", pw, set)
			}
		}
	}
}

func TestGenerateRejectsBadLength(t *testing.T) {
	if _, err := Generate(4, WithLower); err == nil {
		t.Fatal("expected an error for length 4")
	}
}

func TestCheck(t *testing.T) {
	cases := []struct {
		pw    string
		label string
	}{
		{"sunshine", "Weak"},
		{"Sunshine2026!", "Medium"},
		{"T!ger-Moon-42-river", "Strong"},
	}
	for _, c := range cases {
		score, _ := Check(c.pw)
		if got := Label(score); got != c.label {
			t.Errorf("%q: got %s, want %s", c.pw, got, c.label)
		}
	}
}
