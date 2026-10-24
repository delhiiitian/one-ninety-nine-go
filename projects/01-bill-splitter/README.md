# Project 1: Restaurant Bill Splitter

A small command-line tool that asks for a bill amount, a tip percent and the number of
people, then prints a neat receipt with tax, tip, total and each person's share.

Book: *Go in 199 Pages*, Chapter 1 project.

## Run it

```bash
go run ./projects/01-bill-splitter
# or without typing:
printf "2450\n10\n4\n" | go run ./projects/01-bill-splitter
```

## Test it

```bash
go test ./projects/01-bill-splitter
```

## Files

| File | Job |
| --- | --- |
| `main.go` | Talks to the user: reads input, prints the receipt |
| `bill.go` | Pure maths: constants, the `Bill` type and `Calculate` |
| `bill_test.go` | Checks the maths, including rounding so nobody underpays |

## Ideas to extend

- Ask for a service charge as well as a tip.
- Let each person pay a different share (for example, by what they ordered).
- Print the receipt in another currency.
