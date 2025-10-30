package day08

import (
	"advent-of-code-go/util/stringy"
	"regexp"
	"strconv"
)

var escapeSeqs = []struct {
	regex       *regexp.Regexp
	nToSubtract int
}{
	{regexp.MustCompile(`\\x[0-9a-f]{2}`), 3},
	{regexp.MustCompile(`\\\\`), 1},
	{regexp.MustCompile(`\\"`), 1},
}

func part1(input string) int {
	codeCharCount, memCharCount := 0, 0

	for line := range stringy.SplitLines(input) {
		codeCharCount += len(line)
		line = line[1 : len(line)-1]

		memCharCount += len(line)
		for _, seq := range escapeSeqs {
			if matches := seq.regex.FindAllString(line, -1); matches != nil {
				memCharCount -= seq.nToSubtract * len(matches)
			}
		}
	}

	return codeCharCount - memCharCount
}

func part2(input string) int {
	originalCodeCount, encodedCodeCount := 0, 0
	for line := range stringy.SplitLines(input) {
		originalCodeCount += len(line)
		encodedCodeCount += len(strconv.Quote(line))
	}
	return encodedCodeCount - originalCodeCount
}
