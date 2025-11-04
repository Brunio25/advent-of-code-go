package day08

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var exampleInput = `"acb"
"aaa\"aaa"
"\x27"`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 21 - 11},
		{"Actual", input, 1371},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", exampleInput, 36 - 21},
		{"Actual", input, 2117},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
