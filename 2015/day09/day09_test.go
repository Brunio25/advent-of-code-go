package day09

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var exampleInput = `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs{
		{"Example", exampleInput, 605},
		{"Actual", input, 141},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs{
		{"Example", exampleInput, 982},
		{"Actual", input, 736},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
