// Chapter 8.1: Goroutines (doing several things at once)
// Run: go run ./ch08-concurrency-generics/01-goroutines
package main

import (
	"fmt"
	"time"
)

// cook pretends to prepare a dish that takes some time.
func cook(dish string, d time.Duration) {
	fmt.Println("start", dish)
	time.Sleep(d)
	fmt.Println("done ", dish)
}

func main() {
	// One cook, one dish after another: 300ms + 200ms + 100ms
	start := time.Now()
	cook("rice", 300*time.Millisecond)
	cook("dal", 200*time.Millisecond)
	cook("salad", 100*time.Millisecond)
	fmt.Printf("one by one: %v\n\n", time.Since(start).Round(100*time.Millisecond))

	// Three cooks at once: the "go" keyword starts a goroutine.
	start = time.Now()
	go cook("rice", 300*time.Millisecond)
	go cook("dal", 200*time.Millisecond)
	go cook("salad", 100*time.Millisecond)

	// main must wait, or the program ends before the goroutines finish.
	// (Sleeping is a crude way to wait; lesson 8.2 shows the right way.)
	time.Sleep(350 * time.Millisecond)
	fmt.Printf("all at once: %v\n", time.Since(start).Round(100*time.Millisecond))

	// Goroutines are cheap: start 10,000 of them easily.
	done := make(chan bool)
	for i := 0; i < 10_000; i++ {
		go func() { done <- true }()
	}
	for i := 0; i < 10_000; i++ {
		<-done
	}
	fmt.Println("10,000 goroutines started and finished")
}
