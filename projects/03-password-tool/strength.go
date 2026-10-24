package main

import (
	"strings"
	"unicode"
)

// A Rule checks one thing about a password and gives a tip if it fails.
type Rule struct {
	Tip   string
	Check func(pw string) bool
}

// hasAny returns a check function: a closure that remembers which test to use.
func hasAny(test func(rune) bool) func(string) bool {
	return func(pw string) bool {
		for _, r := range pw {
			if test(r) {
				return true
			}
		}
		return false
	}
}

func minLength(n int) func(string) bool {
	return func(pw string) bool { return len([]rune(pw)) >= n }
}

var commonPasswords = []string{"password", "123456", "qwerty", "sunshine", "iloveyou", "admin", "welcome"}

func notCommon(pw string) bool {
	lowerPw := strings.ToLower(pw)
	for _, c := range commonPasswords {
		if strings.Contains(lowerPw, c) {
			return false
		}
	}
	return true
}

// Rules is the full checklist, in the order tips are shown.
var Rules = []Rule{
	{"use at least 12 characters", minLength(12)},
	{"add a lowercase letter", hasAny(unicode.IsLower)},
	{"add an uppercase letter", hasAny(unicode.IsUpper)},
	{"add a digit", hasAny(unicode.IsDigit)},
	{"add a symbol such as ! or #", hasAny(func(r rune) bool { return unicode.IsPunct(r) || unicode.IsSymbol(r) })},
	{"avoid common words like 'password'", notCommon},
}

// Check returns how many rules pass and a tip for each one that fails.
func Check(pw string) (score int, tips []string) {
	for _, rule := range Rules {
		if rule.Check(pw) {
			score++
		} else {
			tips = append(tips, rule.Tip)
		}
	}
	return score, tips
}

// Label turns a score into a word.
func Label(score int) string {
	switch {
	case score == len(Rules):
		return "Strong"
	case score >= 4:
		return "Medium"
	default:
		return "Weak"
	}
}
