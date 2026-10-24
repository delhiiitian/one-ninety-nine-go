// Chapter 5.4: Wrapping errors with context
// Run: go run ./ch05-error-handling/04-wrapping-errors
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Lowest level: reads a file
func readConfig(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %q: %w", path, err) // %w wraps
	}
	return data, nil
}

// Middle level: parses a port number
func loadPort(path string) (int, error) {
	data, err := readConfig(path)
	if err != nil {
		return 0, fmt.Errorf("load port: %w", err)
	}
	port, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, fmt.Errorf("load port: %w", err)
	}
	return port, nil
}

func main() {
	// Top level: one message that tells the whole story.
	_, err := loadPort("missing.conf")
	fmt.Println("error:", err)

	// The original error is still inside and can be checked.
	fmt.Println("file missing?", errors.Is(err, os.ErrNotExist))

	// Unwrap peels one layer at a time.
	inner := errors.Unwrap(err)
	fmt.Println("one layer down:", inner)

	// A file with bad contents
	f, _ := os.CreateTemp("", "port-*.conf")
	f.WriteString("eighty")
	f.Close()
	defer os.Remove(f.Name())

	_, err = loadPort(f.Name())
	var numErr *strconv.NumError
	if errors.As(err, &numErr) {
		fmt.Println("bad number in config:", numErr.Num)
	}
}
