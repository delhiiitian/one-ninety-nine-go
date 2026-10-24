// Chapter 6.3: Reading line by line with bufio.Scanner
// Run: go run ./ch06-files-io/03-read-line-by-line
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const serverLog = `2026-10-24 09:00:01 INFO  server started
2026-10-24 09:00:05 INFO  GET /home 200
2026-10-24 09:01:12 ERROR database timeout
2026-10-24 09:02:40 INFO  GET /cart 200
2026-10-24 09:03:03 WARN  slow response 2.1s
2026-10-24 09:04:59 ERROR payment service unreachable
`

func main() {
	path := filepath.Join(os.TempDir(), "server.log")
	os.WriteFile(path, []byte(serverLog), 0o644)
	defer os.Remove(path)

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Scanner reads one line at a time, so even a 10 GB log uses little memory.
	sc := bufio.NewScanner(f)
	counts := map[string]int{}
	lineNo := 0
	for sc.Scan() {
		lineNo++
		line := sc.Text()
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		level := fields[2]
		counts[level]++
		if level == "ERROR" {
			fmt.Printf("line %d: %s\n", lineNo, strings.Join(fields[3:], " "))
		}
	}
	if err := sc.Err(); err != nil { // check for read errors after the loop
		fmt.Println("scan error:", err)
	}
	fmt.Printf("INFO=%d WARN=%d ERROR=%d (total %d lines)\n",
		counts["INFO"], counts["WARN"], counts["ERROR"], lineNo)

	// Scanner can split by words too.
	ws := bufio.NewScanner(strings.NewReader("one two  three\nfour"))
	ws.Split(bufio.ScanWords)
	words := 0
	for ws.Scan() {
		words++
	}
	fmt.Println("words:", words)
}
