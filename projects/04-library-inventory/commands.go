package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

const helpText = `Commands:
  list [title|author|year]        show all books, optionally sorted
  search <words>                  find by title, author or genre
  add <title>;<author>;<year>;<genre>;<copies>
  borrow <id> <member>            lend a copy
  return <id> <member>            take a copy back
  who <id>                        who has this book
  report                          titles and copies per genre
  quit`

// Run executes one command line and returns the text to print.
// It never prints itself, which keeps it easy to test.
func Run(lib *Library, line string) string {
	cmd, rest, _ := strings.Cut(line, " ")
	rest = strings.TrimSpace(rest)

	switch strings.ToLower(cmd) {
	case "help":
		return helpText
	case "list":
		books := lib.All()
		if rest != "" {
			sorted, err := lib.Sorted(rest)
			if err != nil {
				return "Error: " + err.Error()
			}
			books = sorted
		}
		return table(books)
	case "search":
		if rest == "" {
			return "Usage: search <words>"
		}
		found := lib.Search(rest)
		if len(found) == 0 {
			return "No matches for " + strconv.Quote(rest)
		}
		return table(found)
	case "add":
		return add(lib, rest)
	case "borrow", "return":
		idText, member, ok := strings.Cut(rest, " ")
		id, err := strconv.Atoi(idText)
		if !ok || err != nil || strings.TrimSpace(member) == "" {
			return "Usage: " + cmd + " <id> <member>"
		}
		member = strings.TrimSpace(member)
		if cmd == "borrow" {
			err = lib.Borrow(id, member)
		} else {
			err = lib.Return(id, member)
		}
		if err != nil {
			return "Error: " + err.Error()
		}
		b, _ := lib.Get(id)
		return fmt.Sprintf("OK. %q now has %d/%d on the shelf.", b.Title, b.Available(), b.Copies)
	case "who":
		id, err := strconv.Atoi(rest)
		if err != nil {
			return "Usage: who <id>"
		}
		b, err := lib.Get(id)
		if err != nil {
			return "Error: " + err.Error()
		}
		if len(b.Borrowed) == 0 {
			return "Nobody has borrowed " + strconv.Quote(b.Title)
		}
		var lines []string
		for _, m := range slices.Sorted(maps.Keys(b.Borrowed)) {
			lines = append(lines, fmt.Sprintf("  %s (%d)", m, b.Borrowed[m]))
		}
		return strings.Join(lines, "\n")
	case "report":
		report := lib.GenreReport()
		var sb strings.Builder
		fmt.Fprintf(&sb, "%-12s %6s %6s %4s\n", "GENRE", "TITLES", "COPIES", "OUT")
		for _, g := range slices.Sorted(maps.Keys(report)) {
			s := report[g]
			fmt.Fprintf(&sb, "%-12s %6d %6d %4d\n", g, s.Titles, s.Copies, s.Out)
		}
		return strings.TrimRight(sb.String(), "\n")
	default:
		return "Unknown command " + strconv.Quote(cmd) + ". Type 'help'."
	}
}

func add(lib *Library, rest string) string {
	parts := strings.Split(rest, ";")
	if len(parts) != 5 {
		return "Usage: add <title>;<author>;<year>;<genre>;<copies>"
	}
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	year, err1 := strconv.Atoi(parts[2])
	copies, err2 := strconv.Atoi(parts[4])
	if err1 != nil || err2 != nil || copies < 1 {
		return "Year and copies must be numbers, and copies at least 1."
	}
	b := lib.Add(parts[0], parts[1], year, parts[3], copies)
	return fmt.Sprintf("Added book #%d.", b.ID)
}

func table(books []*Book) string {
	lines := make([]string, len(books))
	for i, b := range books {
		lines[i] = b.String()
	}
	return strings.Join(lines, "\n")
}
