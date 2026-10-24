// Chapter 6.4: Appending to a file (a simple diary / log)
// Run: go run ./ch06-files-io/04-append-to-file
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// appendLine opens the file in append mode, creating it if needed.
func appendLine(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open diary: %w", err)
	}
	defer f.Close()

	stamp := time.Date(2026, 10, 24, 8, 0, 0, 0, time.UTC).Format("2006-01-02")
	if _, err := fmt.Fprintf(f, "%s  %s\n", stamp, text); err != nil {
		return fmt.Errorf("write diary: %w", err)
	}
	return nil
}

func main() {
	path := filepath.Join(os.TempDir(), "diary.txt")
	os.Remove(path) // start fresh for the demo
	defer os.Remove(path)

	for _, entry := range []string{"Started learning Go", "Wrote my first file", "Appended to it!"} {
		if err := appendLine(path, entry); err != nil {
			fmt.Println(err)
			return
		}
	}

	data, _ := os.ReadFile(path)
	fmt.Print(string(data))

	// The flags combine with | :
	//   os.O_RDONLY  read only        os.O_WRONLY  write only
	//   os.O_CREATE  create if absent  os.O_APPEND  add to the end
	//   os.O_TRUNC   empty it first    os.O_EXCL    fail if it exists
}
