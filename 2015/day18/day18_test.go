package day18

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var example = `.#.#.#
...##.
#....#
..#...
#.#..#
####..`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		//{"Example", example, 19}, // Return config.example from getConfig
		{"Actual", input, 1061},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		//{"Example", example, 3}, // Return config.example from getConfig
		{"Actual", input, 4},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
