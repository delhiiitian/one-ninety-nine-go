package main

import (
	"errors"
	"strings"
	"testing"
)

func TestBorrowAndReturn(t *testing.T) {
	lib := NewLibrary()
	b := lib.Add("Go", "Gopher", 2020, "programming", 1)

	if err := lib.Borrow(b.ID, "asha"); err != nil {
		t.Fatal(err)
	}
	if err := lib.Borrow(b.ID, "ravi"); !errors.Is(err, ErrNoCopies) {
		t.Fatalf("want ErrNoCopies, got %v", err)
	}
	if err := lib.Return(b.ID, "ravi"); err == nil {
		t.Fatal("ravi never borrowed it, want an error")
	}
	if err := lib.Return(b.ID, "asha"); err != nil {
		t.Fatal(err)
	}
	if b.Available() != 1 {
		t.Fatalf("available = %d, want 1", b.Available())
	}
}

func TestSearchAndSort(t *testing.T) {
	lib := NewLibrary()
	seed(lib)
	if got := len(lib.Search("PROGRAMMING")); got != 2 {
		t.Fatalf("search found %d, want 2", got)
	}
	books, err := lib.Sorted("year")
	if err != nil {
		t.Fatal(err)
	}
	if books[0].Title != "Malgudi Days" {
		t.Fatalf("oldest = %q, want Malgudi Days", books[0].Title)
	}
}

func TestRunCommands(t *testing.T) {
	lib := NewLibrary()
	seed(lib)
	if out := Run(lib, "add Clean Code;Robert Martin;2008;Programming;2"); out != "Added book #6." {
		t.Fatalf("add: %q", out)
	}
	if out := Run(lib, "borrow 6 kiran"); !strings.Contains(out, "1/2") {
		t.Fatalf("borrow: %q", out)
	}
	if out := Run(lib, "dance"); !strings.HasPrefix(out, "Unknown command") {
		t.Fatalf("unknown: %q", out)
	}
}
