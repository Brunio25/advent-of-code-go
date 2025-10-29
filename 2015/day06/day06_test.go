package day06

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs{
		{"Example1", "turn on 0,0 through 999,999", 1000000},
		{"Example2", "toggle 0,0 through 999,0", 1000},
		{"Example3", "turn off 499,499 through 500,500", 0},
		{"Actual", input, 377891},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs{
		{"Example1", "turn on 0,0 through 0,0", 1},
		{"Example2", "toggle 0,0 through 999,999", 2000000},
		{"Actual", input, 14110788},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
