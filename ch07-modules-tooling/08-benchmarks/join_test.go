package main

import (
	"fmt"
	"strings"
	"testing"
)

var words = strings.Fields(strings.Repeat("go is simple and fast ", 200)) // 1000 words

func TestBothAgree(t *testing.T) {
	if JoinPlus(words) != JoinBuilder(words) {
		t.Fatal("the two implementations disagree")
	}
}

// Benchmarks start with Benchmark and run the code b.N times.
// Run: go test -bench . -benchmem ./ch07-modules-tooling/08-benchmarks
func BenchmarkJoinPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinPlus(words)
	}
}

func BenchmarkJoinBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinBuilder(words)
	}
}

// Examples are tests AND documentation: go test checks the Output comment.
func ExampleJoinBuilder() {
	fmt.Println(JoinBuilder([]string{"tea", "coffee", "juice"}))
	// Output: tea,coffee,juice
}
