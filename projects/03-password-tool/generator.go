package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// Character sets a password can draw from.
const (
	lower   = "abcdefghijkmnopqrstuvwxyz" // no 'l' (looks like 1)
	upper   = "ABCDEFGHJKLMNPQRSTUVWXYZ"  // no 'I' or 'O'
	digits  = "23456789"                  // no 0 or 1
	symbols = "!@#$%^&*-_=+?"
)

// An Option adds one character set. Options are plain functions,
// so callers pick exactly what they want: Generate(12, WithLower, WithDigits).
type Option func() string

func WithLower() string   { return lower }
func WithUpper() string   { return upper }
func WithDigits() string  { return digits }
func WithSymbols() string { return symbols }

// Generate builds a random password that contains at least one character
// from every chosen set.
func Generate(length int, options ...Option) (string, error) {
	if length < 8 || length > 64 {
		return "", errors.New("length must be between 8 and 64")
	}
	if len(options) == 0 {
		return "", errors.New("choose at least one character set")
	}

	var sets []string
	all := ""
	for _, opt := range options {
		set := opt()
		sets = append(sets, set)
		all += set
	}

	pw := make([]byte, length)
	// One guaranteed character from each set...
	for i, set := range sets {
		pw[i] = pick(set)
	}
	// ...and the rest from everything.
	for i := len(sets); i < length; i++ {
		pw[i] = pick(all)
	}
	shuffle(pw)
	return string(pw), nil
}

// pick returns one random byte from s using a secure random source.
func pick(s string) byte {
	return s[randomInt(len(s))]
}

// shuffle mixes the bytes so the guaranteed characters are not always first.
func shuffle(b []byte) {
	for i := len(b) - 1; i > 0; i-- {
		j := randomInt(i + 1)
		b[i], b[j] = b[j], b[i]
	}
}

func randomInt(n int) int {
	v, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		panic(err) // the operating system's random source failed: nothing sensible to do
	}
	return int(v.Int64())
}
