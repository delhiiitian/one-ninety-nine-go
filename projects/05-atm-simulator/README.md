# Project 5: ATM Simulator

A cash machine in your terminal: insert a card, enter a PIN, check your balance,
withdraw in notes, deposit, and print a mini statement. Every rule a real ATM enforces
becomes an error the screen explains: wrong PIN, blocked card, insufficient funds,
daily limit, amounts that are not multiples of 100, and an ATM that is low on notes.

Book: *Go in 199 Pages*, Chapter 5 project.

## Run it

```bash
go run ./projects/05-atm-simulator
# demo cards: 1111 (PIN 1234), 2222 (PIN 0000)
printf "1111\n1234\n2\n2300\n2\n250\n2\n90000\n4\n5\noff\n" | go run ./projects/05-atm-simulator
go test ./projects/05-atm-simulator
```

## Files

| File | Job |
| --- | --- |
| `errors.go` | Sentinel errors and two custom error types that carry numbers |
| `bank.go` | Accounts, PIN checks with blocking, withdraw and deposit rules |
| `atm.go` | The machine's own note stock and the note-planning algorithm |
| `main.go` | Screens and menus, `explain` (error to message) and `safely` (panic to error) |
| `atm_test.go` | Tests for each rule |

## Ideas to extend

- Add a "change PIN" option that needs the old PIN.
- Reset `WithdrawnDay` when the date changes.
- Add a 2000-rupee note cassette and see the planner use it automatically.
