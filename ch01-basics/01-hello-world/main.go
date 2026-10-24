// Chapter 1.1: Hello, World
// Run: go run ./ch01-basics/01-hello-world
package main // every runnable Go program lives in package main

import "fmt" // fmt = "format": printing and formatting text

// main is where the program starts, like the front door of a house.
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to Go in 199 Pages.")

	// Print without a newline, then add one yourself.
	fmt.Print("Go was created at Google ")
	fmt.Print("in 2007.\n")
}
