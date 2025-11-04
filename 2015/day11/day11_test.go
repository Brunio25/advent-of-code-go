package day11

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[string]{
		{"Example1", "abcdefgh", "abcdffaa"},
		{"Example2", "ghijklmn", "ghjaabcc"},
		{"Actual", input, "hxbxxyzz"},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[string]{
		{"Example1", "abcdefgh", "abcdffbb"},
		{"Example2", "ghijklmn", "ghjbbcdd"},
		{"Actual", input, "hxcaabcc"},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
