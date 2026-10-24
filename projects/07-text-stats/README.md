# Project 7: Text Statistics Tool

A command-line tool and a reusable Go package that measure any text: lines, words,
sentences, reading time, a readability score, and the most common words. It shows how a
real Go project is organised into a library package plus a small command, and how it is
tested: table-driven tests, a test data file, a benchmark and a documented example.

Book: *Go in 199 Pages*, Chapter 7 project.

## Run it

```bash
go run ./projects/07-text-stats projects/07-text-stats/testdata/sample.txt
go run ./projects/07-text-stats -top 3 README.md
echo "Go is fun. Go is fast." | go run ./projects/07-text-stats
```

## Test, benchmark, document

```bash
go test -v ./projects/07-text-stats/...
go test -bench . -benchmem ./projects/07-text-stats/textstats
go test -cover ./projects/07-text-stats/textstats
go doc ./projects/07-text-stats/textstats
go build -o textstats ./projects/07-text-stats
```

## Layout

| Path | Job |
| --- | --- |
| `main.go` | The command: flags, files or stdin, printing the report |
| `textstats/stats.go` | The library: `Analyze`, `Top`, `ReadingMinutes`, `Readability`, `Level` |
| `textstats/stats_test.go` | Table-driven tests, a file-based test and a benchmark |
| `textstats/example_test.go` | A runnable example that also appears in `go doc` |
| `testdata/sample.txt` | Input used by the tests (`go` ignores `testdata` folders when building) |

## Ideas to extend

- Add a `-json` flag that prints the stats as JSON.
- Count the longest sentence and the longest word.
- Support Hindi by counting syllables differently for Devanagari text.
