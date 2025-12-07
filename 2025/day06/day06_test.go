package day06

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var exampleInput = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 4277556},
		{"Actual", input, 5782351442566},
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
