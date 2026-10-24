// Chapter 6.6: File paths that work on every operating system
// Run: go run ./ch06-files-io/06-file-paths
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Join uses / on macOS and Linux and \ on Windows automatically.
	p := filepath.Join("projects", "invoices", "2026", "oct.pdf")
	fmt.Println("joined:", filepath.ToSlash(p))

	fmt.Println("dir: ", filepath.ToSlash(filepath.Dir(p)))
	fmt.Println("base:", filepath.Base(p))
	fmt.Println("ext: ", filepath.Ext(p))
	fmt.Println("name without ext:", strings.TrimSuffix(filepath.Base(p), filepath.Ext(p)))

	// Clean tidies messy paths.
	fmt.Println("clean:", filepath.ToSlash(filepath.Clean("projects//invoices/../reports/./q3.xlsx")))

	// Abs turns a relative path into a full one.
	abs, _ := filepath.Abs("go.mod")
	fmt.Println("absolute path is absolute?", filepath.IsAbs(abs))

	// Useful places
	home, _ := os.UserHomeDir()
	fmt.Println("home folder found?", home != "")
	cfg, _ := os.UserConfigDir()
	fmt.Println("config folder found?", cfg != "")
	fmt.Println("separator on this OS:", string(filepath.Separator))

	// Match with wildcards
	for _, name := range []string{"photo.jpg", "photo.png", "notes.txt"} {
		ok, _ := filepath.Match("*.jpg", name)
		fmt.Printf("%-10s matches *.jpg? %v\n", name, ok)
	}
}
