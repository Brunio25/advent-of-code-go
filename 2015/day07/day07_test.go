package day07

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Actual", input, 956},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Actual", input, 40149},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
