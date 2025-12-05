package day05

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var exampleInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 3},
		{"Actual", input, 613},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 14},
		{"Actual", input, 336495597913098},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
