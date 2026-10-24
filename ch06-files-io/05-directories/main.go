// Chapter 6.5: Working with directories (folders)
// Run: go run ./ch06-files-io/05-directories
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// MkdirTemp makes a fresh folder with a unique name.
	root, err := os.MkdirTemp("", "photos-")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(root) // clean up everything inside

	// MkdirAll creates nested folders in one go (like mkdir -p).
	for _, dir := range []string{"2025/goa", "2026/kerala", "2026/ladakh"} {
		os.MkdirAll(filepath.Join(root, dir), 0o755)
	}
	files := map[string]string{
		"2025/goa/beach.jpg":        "....",
		"2026/kerala/houseboat.jpg": "......",
		"2026/ladakh/pangong.jpg":   "........",
		"2026/ladakh/notes.txt":     "cold!",
	}
	for name, content := range files {
		os.WriteFile(filepath.Join(root, name), []byte(content), 0o644)
	}

	// ReadDir lists one folder.
	entries, _ := os.ReadDir(filepath.Join(root, "2026"))
	fmt.Println("folders in 2026:")
	for _, e := range entries {
		fmt.Println(" -", e.Name(), "dir:", e.IsDir())
	}

	// WalkDir visits every file and folder underneath, recursively.
	fmt.Println("all photos:")
	var total int64
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".jpg") {
			info, _ := d.Info()
			total += info.Size()
			rel, _ := filepath.Rel(root, path)
			fmt.Printf(" - %-28s %d bytes\n", filepath.ToSlash(rel), info.Size())
		}
		return nil
	})
	fmt.Println("total photo bytes:", total)

	// Rename moves a file.
	old := filepath.Join(root, "2026/ladakh/notes.txt")
	os.Rename(old, filepath.Join(root, "notes.txt"))
	_, err = os.Stat(old)
	fmt.Println("notes moved:", os.IsNotExist(err))
}
