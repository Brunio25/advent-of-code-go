package day01

import (
	"advent-of-code-go/util"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var testCasesPart1 = util.TestInputs{
	{"Example", ")())())", -3},
	{"Actual", input, 138},
}

func TestPart1(t *testing.T) {
	for _, tc := range testCasesPart1 {
		t.Run(tc.Name, func(t *testing.T) {
			if actual := part1(tc.Input); actual != tc.Expected {
				t.Errorf("part1() = %v, Expected %v", actual, tc.Expected)
			}
		})
	}
}

var testCasesPart2 = util.TestInputs{
	{"Example", "()())", 5},
	{"Actual", input, 1771},
}

func TestPart2(t *testing.T) {
	for _, tc := range testCasesPart2 {
		t.Run(tc.Name, func(t *testing.T) {
			if actual := part2(tc.Input); actual != tc.Expected {
				t.Errorf("part2() = %v, Expected %v", actual, tc.Expected)
			}
		})
	}
}
