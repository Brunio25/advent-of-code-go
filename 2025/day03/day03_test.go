package day03

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var exampleInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 357},
		{"Actual", input, 17403},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 3121910778619},
		{"Actual", input, 173416889848394},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
