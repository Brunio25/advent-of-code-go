package day16

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Actual", input, 40},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Actual", input, 241},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
