package main

import "fmt"

// Book is one title in the catalogue. Copies counts how many the library owns.
type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	Genre    string
	Copies   int
	Borrowed map[string]int // member name -> copies they hold
}

// Available returns how many copies are on the shelf right now.
func (b *Book) Available() int {
	out := 0
	for _, n := range b.Borrowed {
		out += n
	}
	return b.Copies - out
}

// String makes fmt print a book as one tidy catalogue line.
func (b *Book) String() string {
	return fmt.Sprintf("#%-3d %-30.30s %-22.22s %4d  %-11s %d/%d on shelf",
		b.ID, b.Title, b.Author, b.Year, b.Genre, b.Available(), b.Copies)
}
