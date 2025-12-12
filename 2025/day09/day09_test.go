package day09

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var exampleInput = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 50},
		{"Actual", input, 4763509452},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 25272},
		{"Actual", input, 107256172},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
