// Project 7: Text Statistics Tool
//
//	go run ./projects/07-text-stats projects/07-text-stats/testdata/sample.txt
//	go run ./projects/07-text-stats -top 5 file1.txt file2.txt
//	cat essay.txt | go run ./projects/07-text-stats
//	go test -v ./projects/07-text-stats/...
//	go test -bench . ./projects/07-text-stats/textstats
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/delhiiitian/one-ninety-nine-go/projects/07-text-stats/textstats"
)

func main() {
	top := flag.Int("top", 5, "how many common words to show")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: textstats [-top N] [file ...]   (reads standard input if no files)")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		if err := report(os.Stdout, "standard input", os.Stdin, *top); err != nil {
			fail(err)
		}
		return
	}
	for _, path := range flag.Args() {
		f, err := os.Open(path)
		if err != nil {
			fail(err)
		}
		err = report(os.Stdout, path, f, *top)
		f.Close()
		if err != nil {
			fail(err)
		}
	}
}

func report(w io.Writer, name string, r io.Reader, top int) error {
	s, err := textstats.Analyze(r)
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	score := s.Readability()
	fmt.Fprintf(w, "== %s ==\n", name)
	fmt.Fprintf(w, "lines %d, words %d, sentences %d, characters %d\n", s.Lines, s.Words, s.Sentences, s.Characters)
	fmt.Fprintf(w, "reading time: about %d min\n", s.ReadingMinutes())
	fmt.Fprintf(w, "readability: %.1f (%s)\n", score, textstats.Level(score))
	fmt.Fprintln(w, "most common words:")
	for i, wc := range s.Top(top) {
		fmt.Fprintf(w, "  %d. %-12s %d\n", i+1, wc.Word, wc.Count)
	}
	return nil
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
