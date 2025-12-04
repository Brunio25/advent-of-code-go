package stringy

import (
	"iter"
	"strings"
)

// SplitLines Exactly like strings.Lines except that the lines yielded
// exclude their terminating newlines
func SplitLines(s string) iter.Seq[string] {
	return strings.SplitSeq(s, "\n")
}
