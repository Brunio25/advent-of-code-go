package day05

import (
	"advent-of-code-go/util/tests"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

var examplePart1 = `ugknbfddgicrmopn
aaa
jchzalrnumimnmhp
haegwjzuvuyypxyu
dvszwmarrgswjxmb`

func TestPart1(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", examplePart1, 2},
		{"Actual", input, 236},
	}
	tests.RunWithTestCases(testCases, part1, t)
}

var examplePart2 = `qjhvhtzxzqqjkmpb
xxyxx
uurcxstgmygtbstg
ieodomkazucvgmuy`

func TestPart2(t *testing.T) {
	testCases := tests.TestInputs[int]{
		{"Example", examplePart2, 2},
		{"Actual", input, 51},
	}
	tests.RunWithTestCases(testCases, part2, t)
}
