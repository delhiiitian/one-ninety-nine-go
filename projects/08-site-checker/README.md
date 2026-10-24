# Project 8: Concurrent Website Checker

Checks many websites at the same time and reports which are up, which are down, their
HTTP status codes and how long each took. A pool of worker goroutines does the checks,
every request has a timeout, Ctrl+C cancels cleanly, and results come back in the order
you gave them. A built-in demo site lets you try it without the internet.

Book: *Go in 199 Pages*, Chapter 8 project.

## Run it

```bash
go run ./projects/08-site-checker -demo
go run ./projects/08-site-checker https://go.dev https://example.com
go run ./projects/08-site-checker -file urls.txt -workers 20 -timeout 3s
go test ./projects/08-site-checker/...
```

The program exits with code 1 if any site is down, so you can use it in scripts or CI.

## Layout

| Path | Job |
| --- | --- |
| `main.go` | Flags, collecting URLs, Ctrl+C handling, the printed report |
| `demo.go` | A local website with fast, slow, broken, redirecting and missing pages |
| `checker/checker.go` | `Check` (one URL with a timeout) and `CheckAll` (worker pool, ordered results) |
| `checker/checker_test.go` | Tests against an in-process `httptest` server, including a parallelism test |

## Ideas to extend

- Repeat the checks every minute and print only changes ("shop went DOWN").
- Check that the page body contains an expected word, not just the status code.
- Send an alert to a chat webhook when a site goes down (Chapter 9 shows HTTP clients).
