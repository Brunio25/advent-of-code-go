package day15

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var example = `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", example, 62842880},
		{"Actual", input, 13882464},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", example, 57600000},
		{"Actual", input, 11171160},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
