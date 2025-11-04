package day01

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", ")())())", -3},
		{"Actual", input, 138},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", "()())", 5},
		{"Actual", input, 1771},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
