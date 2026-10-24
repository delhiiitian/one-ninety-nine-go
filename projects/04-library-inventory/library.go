package main

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"strings"
)

// Library stores books two ways:
//   - byID, a map, for instant lookups ("give me book 3")
//   - order, a slice of IDs, to remember the order books were added
type Library struct {
	byID   map[int]*Book
	order  []int
	nextID int
}

func NewLibrary() *Library {
	return &Library{byID: map[int]*Book{}, nextID: 1}
}

// Add puts a new title in the catalogue and returns it.
func (l *Library) Add(title, author string, year int, genre string, copies int) *Book {
	b := &Book{
		ID: l.nextID, Title: title, Author: author, Year: year,
		Genre: strings.ToLower(genre), Copies: copies, Borrowed: map[string]int{},
	}
	l.byID[b.ID] = b
	l.order = append(l.order, b.ID)
	l.nextID++
	return b
}

// Get finds a book by ID.
func (l *Library) Get(id int) (*Book, error) {
	b, ok := l.byID[id]
	if !ok {
		return nil, fmt.Errorf("no book with ID %d", id)
	}
	return b, nil
}

// All returns every book in the order they were added.
func (l *Library) All() []*Book {
	books := make([]*Book, 0, len(l.order))
	for _, id := range l.order {
		books = append(books, l.byID[id])
	}
	return books
}

// Search matches the query against title, author and genre (case-insensitive).
func (l *Library) Search(query string) []*Book {
	q := strings.ToLower(query)
	var found []*Book
	for _, b := range l.All() {
		text := strings.ToLower(b.Title + " " + b.Author + " " + b.Genre)
		if strings.Contains(text, q) {
			found = append(found, b)
		}
	}
	return found
}

// Sorted returns a copy of the catalogue ordered by the given field.
func (l *Library) Sorted(field string) ([]*Book, error) {
	books := l.All()
	var compare func(a, b *Book) int
	switch field {
	case "title":
		compare = func(a, b *Book) int { return cmp.Compare(a.Title, b.Title) }
	case "author":
		compare = func(a, b *Book) int { return cmp.Compare(a.Author, b.Author) }
	case "year":
		compare = func(a, b *Book) int { return cmp.Compare(a.Year, b.Year) }
	default:
		return nil, fmt.Errorf("cannot sort by %q (use title, author or year)", field)
	}
	slices.SortStableFunc(books, compare)
	return books, nil
}

var ErrNoCopies = errors.New("no copies left on the shelf")

// Borrow lends one copy to a member.
func (l *Library) Borrow(id int, member string) error {
	b, err := l.Get(id)
	if err != nil {
		return err
	}
	if b.Available() == 0 {
		return ErrNoCopies
	}
	b.Borrowed[member]++
	return nil
}

// Return takes one copy back from a member.
func (l *Library) Return(id int, member string) error {
	b, err := l.Get(id)
	if err != nil {
		return err
	}
	if b.Borrowed[member] == 0 {
		return fmt.Errorf("%s has not borrowed %q", member, b.Title)
	}
	b.Borrowed[member]--
	if b.Borrowed[member] == 0 {
		delete(b.Borrowed, member)
	}
	return nil
}

// GenreReport counts titles and copies per genre.
type GenreStats struct{ Titles, Copies, Out int }

func (l *Library) GenreReport() map[string]GenreStats {
	report := map[string]GenreStats{}
	for _, b := range l.All() {
		s := report[b.Genre]
		s.Titles++
		s.Copies += b.Copies
		s.Out += b.Copies - b.Available()
		report[b.Genre] = s
	}
	return report
}
