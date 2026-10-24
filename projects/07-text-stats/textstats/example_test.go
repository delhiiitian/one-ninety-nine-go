package textstats_test

import (
	"fmt"
	"strings"

	"github.com/delhiiitian/one-ninety-nine-go/projects/07-text-stats/textstats"
)

// This example appears in the package documentation (go doc) and is checked by go test.
func ExampleAnalyze() {
	s, _ := textstats.Analyze(strings.NewReader("Go is simple. Go is fast. Simple is good."))
	fmt.Println("words:", s.Words, "sentences:", s.Sentences)
	fmt.Println("top:", s.Top(2))
	// Output:
	// words: 9 sentences: 3
	// top: [{go 2} {simple 2}]
}
