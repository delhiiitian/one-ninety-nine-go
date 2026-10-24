package main

import "strings"

// JoinPlus builds a string with +, creating a new string every time.
func JoinPlus(words []string) string {
	s := ""
	for i, w := range words {
		if i > 0 {
			s += ","
		}
		s += w
	}
	return s
}

// JoinBuilder uses strings.Builder, which grows one buffer.
func JoinBuilder(words []string) string {
	var b strings.Builder
	for i, w := range words {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(w)
	}
	return b.String()
}
