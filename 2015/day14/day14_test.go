package day14

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var example = `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", example, 2660},
		{"Actual", input, 2655},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", example, 689},
		{"Actual", input, 640},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
