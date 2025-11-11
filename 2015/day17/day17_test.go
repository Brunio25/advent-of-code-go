package day17

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var example = `20
15
10
5
5`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		//{"Example", example, 4}, // Must set target to 25
		{"Actual", input, 4372},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		//{"Example", example, 3}, // Must set target to 25
		{"Actual", input, 4},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
