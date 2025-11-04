package day03

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var exampleInput = `^>v<`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 4},
		{"Actual", input, 2572},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 3},
		{"Actual", input, 2631},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
