package stringy_test

import (
	"advent-of-code-go/util/stringy"
	"slices"
	"testing"
)

const inputStr = `First
Second
Third`

var expected = []string{"First", "Second", "Third"}

func TestSplitLines(t *testing.T) {
	actual := slices.Collect(stringy.SplitLines(inputStr))
	if !slices.Equal(actual, expected) {
		t.Errorf("stringy.SplitLines() = %v, Expected %v", actual, expected)
	}
}
