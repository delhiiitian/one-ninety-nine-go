# Project 4: Library Inventory Manager

A command-line catalogue for a small library: list and sort books, search, add titles,
lend and take back copies, see who has a book, and print a report per genre.

Book: *Go in 199 Pages*, Chapter 4 project.

## Run it

```bash
go run ./projects/04-library-inventory
printf "list year\nborrow 2 meera\nwho 2\nreport\nquit\n" | go run ./projects/04-library-inventory
go test ./projects/04-library-inventory
```

## Files

| File | Job |
| --- | --- |
| `main.go` | Reads commands in a loop; adds sample books |
| `book.go` | The `Book` struct, `Available` and `String` methods |
| `library.go` | The `Library` store (map + slice), search, sort, borrow, return, report |
| `commands.go` | `Run` turns one command line into a result string |
| `library_test.go` | Tests for lending rules, search, sorting and commands |

## Ideas to extend

- Add due dates and list overdue books (see `time` in Chapter 9).
- Save the catalogue to a JSON file (Chapter 6).
- Add a `remove <id>` command that refuses if copies are still out.
