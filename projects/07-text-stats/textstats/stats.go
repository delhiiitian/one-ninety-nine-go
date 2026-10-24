// Package textstats measures a piece of text: words, sentences, reading time,
// the most common words and a readability score.
//
// It has no idea where the text comes from (a file, a web form, a pipe):
// it reads from any io.Reader.
package textstats

import (
	"bufio"
	"cmp"
	"io"
	"slices"
	"strings"
	"unicode"
)

// WordsPerMinute is a typical adult silent-reading speed.
const WordsPerMinute = 230

// Stats is the result of analysing one text.
type Stats struct {
	Lines      int
	Words      int
	Characters int // letters and digits only
	Sentences  int
	Syllables  int
	counts     map[string]int
}

// WordCount is one entry of the "most common words" list.
type WordCount struct {
	Word  string
	Count int
}

// stopWords are too common to be interesting in a "top words" list.
var stopWords = map[string]bool{
	"the": true, "a": true, "an": true, "and": true, "or": true, "of": true,
	"to": true, "in": true, "is": true, "it": true, "that": true, "for": true,
	"on": true, "with": true, "as": true, "was": true, "are": true, "be": true,
	"this": true, "by": true, "at": true, "from": true, "i": true, "you": true,
	"their": true, "its": true, "but": true, "not": true, "can": true, "has": true, "have": true,
}

// Analyze reads all text from r and returns its statistics.
func Analyze(r io.Reader) (Stats, error) {
	s := Stats{counts: map[string]int{}}
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 64*1024), 1024*1024) // allow long lines (1 MB)
	for sc.Scan() {
		s.Lines++
		s.addLine(sc.Text())
	}
	return s, sc.Err()
}

func (s *Stats) addLine(line string) {
	for _, raw := range strings.Fields(line) {
		if strings.ContainsAny(raw, ".!?") {
			s.Sentences++
		}
		word := normalize(raw)
		if word == "" {
			continue
		}
		s.Words++
		s.Characters += len([]rune(word))
		s.Syllables += syllables(word)
		s.counts[word]++
	}
}

// normalize lowercases a word and strips punctuation around it: "Hello," -> "hello".
func normalize(w string) string {
	w = strings.TrimFunc(w, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) })
	return strings.ToLower(w)
}

// syllables is a simple English estimate: count groups of vowels.
func syllables(word string) int {
	n, prevVowel := 0, false
	for _, r := range word {
		v := strings.ContainsRune("aeiouy", r)
		if v && !prevVowel {
			n++
		}
		prevVowel = v
	}
	if strings.HasSuffix(word, "e") && !strings.HasSuffix(word, "le") && n > 1 {
		n-- // silent e: "make", "code" (but not "simple")
	}
	return max(n, 1)
}

// Top returns the n most frequent words, skipping stop words.
// Ties are broken alphabetically so the result is always the same.
func (s Stats) Top(n int) []WordCount {
	var list []WordCount
	for w, c := range s.counts {
		if !stopWords[w] {
			list = append(list, WordCount{w, c})
		}
	}
	slices.SortFunc(list, func(a, b WordCount) int {
		if c := cmp.Compare(b.Count, a.Count); c != 0 {
			return c
		}
		return cmp.Compare(a.Word, b.Word)
	})
	return list[:min(n, len(list))]
}

// ReadingMinutes estimates reading time, rounded up to at least one minute.
func (s Stats) ReadingMinutes() int {
	if s.Words == 0 {
		return 0
	}
	return (s.Words + WordsPerMinute - 1) / WordsPerMinute
}

// Readability is the Flesch Reading Ease score: higher is easier.
// 90-100 is easy for an 11-year-old, 60-70 is plain English, below 30 is very hard.
func (s Stats) Readability() float64 {
	if s.Words == 0 || s.Sentences == 0 {
		return 0
	}
	wps := float64(s.Words) / float64(s.Sentences)
	spw := float64(s.Syllables) / float64(s.Words)
	return 206.835 - 1.015*wps - 84.6*spw
}

// Level turns a readability score into words.
func Level(score float64) string {
	switch {
	case score >= 80:
		return "easy"
	case score >= 60:
		return "plain English"
	case score >= 40:
		return "fairly difficult"
	default:
		return "difficult"
	}
}
