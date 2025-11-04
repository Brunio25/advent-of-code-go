package day12

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example1", `[1,2,3]`, 6},
		{"Example2", `{"a":{"b":4},"c":-1}`, 3},
		{"Example3", `[]`, 0},
		{"Actual", input, 111754},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example1", "[1,2,3]", 6},
		{"Example2", `[1,{"c":"red","b":2},3]`, 4},
		{"Example3", `{"d":"red","e":[1,2,3,4],"f":5}`, 0},
		{"Example5", `[1,"red",5]`, 6},
		{"Actual", input, 65402},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
