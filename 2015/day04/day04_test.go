package day04

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs{
		{"Example1", "abcdef", 609043},
		{"Example2", "pqrstuv", 1048970},
		{"Actual", input, 346386},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs{
		{"Actual", input, 9958218},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
