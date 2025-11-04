package day02

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var exampleInput = `2x3x4
1x1x10`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 101},
		{"Actual", input, 1586300},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 48},
		{"Actual", input, 3737498},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
