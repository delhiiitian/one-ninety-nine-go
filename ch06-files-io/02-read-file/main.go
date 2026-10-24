// Chapter 6.2: Reading a whole file
// Run: go run ./ch06-files-io/02-read-file
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := filepath.Join(os.TempDir(), "quote.txt")
	os.WriteFile(path, []byte("Simplicity is complicated.\n- Rob Pike\n"), 0o644)
	defer os.Remove(path)

	// ReadFile returns all the bytes at once: perfect for small files.
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("read failed:", err)
		return
	}
	text := string(data)
	fmt.Print(text)
	fmt.Println("lines:", strings.Count(text, "\n"), "bytes:", len(data))

	// File information without reading it
	info, err := os.Stat(path)
	if err == nil {
		fmt.Println("name:", info.Name(), "size:", info.Size(), "is folder:", info.IsDir())
	}

	// A missing file is a normal error you can recognise.
	_, err = os.ReadFile("does-not-exist.txt")
	if errors.Is(err, fs.ErrNotExist) {
		fmt.Println("no such file; creating a default instead would go here")
	}
}
