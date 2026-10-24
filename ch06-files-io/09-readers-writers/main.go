// Chapter 6.9: io.Reader and io.Writer (the pipes of Go)
// Run: go run ./ch06-files-io/09-readers-writers
package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

// countWriter is our own Writer: it only counts the bytes that flow through it.
type countWriter struct{ n int }

func (c *countWriter) Write(p []byte) (int, error) {
	c.n += len(p)
	return len(p), nil
}

// shout works with ANY reader and ANY writer: files, network, memory...
func shout(dst io.Writer, src io.Reader) error {
	data, err := io.ReadAll(src)
	if err != nil {
		return err
	}
	_, err = dst.Write(bytes.ToUpper(data))
	return err
}

func main() {
	// A string can be a Reader; os.Stdout is a Writer.
	shout(os.Stdout, strings.NewReader("pipes connect everything\n"))

	// io.Copy streams from a reader to a writer in chunks.
	text := strings.Repeat("Go is fun. ", 1000)
	var counter countWriter
	io.Copy(&counter, strings.NewReader(text))
	fmt.Println("bytes counted:", counter.n)

	// Chain pipes: text -> gzip -> memory buffer
	var compressed bytes.Buffer
	zw := gzip.NewWriter(&compressed)
	io.Copy(zw, strings.NewReader(text))
	zw.Close()
	fmt.Printf("gzip: %d bytes -> %d bytes\n", len(text), compressed.Len())

	// ...and back: buffer -> gunzip -> hash
	zr, _ := gzip.NewReader(&compressed)
	h := sha256.New() // a hash is a Writer too
	io.Copy(h, zr)
	fmt.Printf("sha256 of original: %x...\n", h.Sum(nil)[:8])

	// MultiWriter sends the same data to several places at once.
	var log bytes.Buffer
	both := io.MultiWriter(os.Stdout, &log)
	fmt.Fprintln(both, "saved to screen and to the log buffer")
	fmt.Println("log buffer holds", log.Len(), "bytes")
}
