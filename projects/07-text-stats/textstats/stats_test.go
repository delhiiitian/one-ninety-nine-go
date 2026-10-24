package textstats

import (
	"os"
	"strings"
	"testing"
)

func TestAnalyzeCounts(t *testing.T) {
	tests := []struct {
		name                    string
		text                    string
		lines, words, sentences int
	}{
		{"empty", "", 0, 0, 0},
		{"one sentence", "Go is fun.", 1, 3, 1},
		{"punctuation only", "... !!!", 1, 0, 2},
		{"two lines", "Hello there!\nHow are you?", 2, 5, 2},
		{"extra spaces", "  lots   of\tspace  ", 1, 3, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s, err := Analyze(strings.NewReader(tc.text))
			if err != nil {
				t.Fatal(err)
			}
			if s.Lines != tc.lines || s.Words != tc.words || s.Sentences != tc.sentences {
				t.Errorf("got lines=%d words=%d sentences=%d, want %d %d %d",
					s.Lines, s.Words, s.Sentences, tc.lines, tc.words, tc.sentences)
			}
		})
	}
}

func TestTopSkipsStopWordsAndIsStable(t *testing.T) {
	s, _ := Analyze(strings.NewReader("The cat and the dog. The dog barked at the cat. Bird."))
	top := s.Top(3)
	want := []WordCount{{"cat", 2}, {"dog", 2}, {"barked", 1}}
	for i := range want {
		if top[i] != want[i] {
			t.Fatalf("Top(3) = %v, want %v", top, want)
		}
	}
}

func TestSyllables(t *testing.T) {
	for word, want := range map[string]int{"go": 1, "code": 1, "simple": 2, "language": 2, "rhythm": 1} {
		if got := syllables(word); got != want {
			t.Errorf("syllables(%q) = %d, want %d", word, got, want)
		}
	}
}

func TestSampleFile(t *testing.T) {
	f, err := os.Open("../testdata/sample.txt") // tests run inside the package folder
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	s, err := Analyze(f)
	if err != nil {
		t.Fatal(err)
	}
	if s.Words < 100 || s.ReadingMinutes() != 1 {
		t.Fatalf("unexpected stats for sample: %+v", s)
	}
}

func BenchmarkAnalyze(b *testing.B) {
	text := strings.Repeat("Go makes it easy to build simple, reliable, and efficient software. ", 2000)
	for i := 0; i < b.N; i++ {
		Analyze(strings.NewReader(text))
	}
}
