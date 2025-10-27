package day05

import (
	"regexp"
	"strings"

	"github.com/dlclark/regexp2"
)

func part1(input string) int {
	vowelPattern := regexp.MustCompile(`[aeiou]`)
	twiceInRowPattern := regexp2.MustCompile(`(\w)\1`, regexp2.None)
	invalidPattern := regexp.MustCompile("(ab|cd|pq|xy)")

	niceStrings := 0
	for line := range strings.Lines(input) {
		twiceInRowMatch, _ := twiceInRowPattern.MatchString(line)
		if len(vowelPattern.FindAllString(line, -1)) >= 3 &&
			twiceInRowMatch &&
			!invalidPattern.MatchString(line) {
			niceStrings++
		}
	}

	return niceStrings
}

func part2(input string) int {
	doubleTwiceInRowPattern := regexp2.MustCompile(`(\w\w)\w*\1`, regexp2.None)
	interleavedSamePattern := regexp2.MustCompile(`(\w)\w\1`, regexp2.None)

	niceStrings := 0
	for line := range strings.Lines(input) {
		doubleTwiceMatch, _ := doubleTwiceInRowPattern.MatchString(line)
		interleavedSameMatch, _ := interleavedSamePattern.MatchString(line)

		if doubleTwiceMatch && interleavedSameMatch {
			niceStrings++
		}
	}

	return niceStrings
}
