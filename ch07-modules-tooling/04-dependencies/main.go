// Chapter 7.4: Using third-party packages
//
// This folder is its own small module (it has its own go.mod) so the main
// book code stays dependency-free. To run it:
//
//	cd ch07-modules-tooling/04-dependencies
//	go run .
//
// How the dependency was added:
//
//	go mod init example.com/ticketing
//	go get github.com/google/uuid@v1.6.0
//	go mod tidy
package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Ticket struct {
	ID    uuid.UUID
	Event string
	Seat  string
}

func main() {
	// Random (version 4) UUIDs: practically unique without a central counter.
	for _, seat := range []string{"A1", "A2", "A3"} {
		t := Ticket{ID: uuid.New(), Event: "Go Conference 2026", Seat: seat}
		fmt.Printf("%s  %s  seat %s\n", t.ID, t.Event, t.Seat)
	}

	// Parse checks the format of an ID someone gives you.
	if _, err := uuid.Parse("not-a-uuid"); err != nil {
		fmt.Println("invalid ticket ID:", err)
	}
}
