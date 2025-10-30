package stringy

import (
	"iter"
	"strings"
)

// SplitLines Exactly like strings.Lines except that the lines yielded
// exclude their terminating newlines
func SplitLines(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for line := range strings.Lines(s) {
			line = strings.TrimRight(line, "\n")
			if !yield(line) {
				return
			}
		}
	}
}
