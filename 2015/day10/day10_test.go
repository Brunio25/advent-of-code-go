package day10

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		//{"Example1", "1", 605},
		//{"Example2", "11", 605},
		//{"Example3", "21", 605},
		//{"Example4", "1211", 605},
		//{"Example5", "111221", 605},
		{"Actual", input, 360154},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		//{"Example1", "1", 605},
		//{"Example2", "11", 605},
		//{"Example3", "21", 605},
		//{"Example4", "1211", 605},
		//{"Example5", "111221", 605},
		{"Actual", input, 360154},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
