# Project 3: Password Generator and Strength Checker

Generate strong random passwords and check how strong an existing password is, with
specific tips to improve it.

Book: *Go in 199 Pages*, Chapter 3 project.

## Run it

```bash
go run ./projects/03-password-tool
printf "1\n16\n2\nsunshine\n3\n" | go run ./projects/03-password-tool
go test ./projects/03-password-tool
```

## Files

| File | Job |
| --- | --- |
| `main.go` | The menu loop and the `report` printer |
| `generator.go` | `Generate` with variadic `Option` functions, secure random picking and shuffling |
| `strength.go` | `Rule` list built from closures, `Check` and `Label` |
| `password_test.go` | Tests for generation and checking |

## Ideas to extend

- Add an option to generate passphrases made of four random words.
- Add a rule that rejects three repeated characters in a row ("aaa").
- Estimate the time to crack the password from its length and character sets.
