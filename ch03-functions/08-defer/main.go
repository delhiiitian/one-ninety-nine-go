// Chapter 3.8: defer (do this when the function ends)
// Run: go run ./ch03-functions/08-defer
package main

import "fmt"

func cookDinner() {
	fmt.Println("Turn on the stove")
	defer fmt.Println("Turn off the stove") // runs last, no matter how we leave

	fmt.Println("Cook rice")
	fmt.Println("Cook dal")
}

// Deferred calls run in reverse order: last in, first out (like a stack of plates).
func stackOfPlates() {
	for i := 1; i <= 3; i++ {
		defer fmt.Println("wash plate", i)
	}
	fmt.Println("dinner served")
}

// defer is ideal for cleanup, even when we return early.
func useMeetingRoom(booked bool) {
	fmt.Println("unlock room")
	defer fmt.Println("lock room")
	if !booked {
		fmt.Println("not booked: leaving early")
		return
	}
	fmt.Println("hold meeting")
}

func main() {
	cookDinner()
	fmt.Println()
	stackOfPlates()
	fmt.Println()
	useMeetingRoom(true)
	useMeetingRoom(false)
}
