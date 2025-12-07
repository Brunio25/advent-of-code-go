package day06

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/slicy"
	"strings"
	"unicode"
)

func part1(input string) int {
	lines := strings.Split(input, "\n")
	splitLines := make([][]string, len(lines))
	for i, line := range lines {
		splitLines[i] = strings.Fields(line)
	}

	problemNumber := len(splitLines[0])
	problemNumsLength := len(splitLines) - 1

	res := 0
	for problemIndex := 0; problemIndex < problemNumber; problemIndex++ {
		nums := make([]int, problemNumsLength)
		for line := 0; line < problemNumsLength; line++ {
			nums[line] = cast.ToInt(splitLines[line][problemIndex])
		}

		res += slicy.Reduce(nums, getOperation(splitLines[problemNumsLength][problemIndex]))
	}

	return res
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	problemLastIndex := len(lines) - 1

	res := 0
	currProblemNums := make([]int, 0)
	for i := len(lines[0]) - 1; i >= 0; i-- {
		currNum := 0
		for lineIndex := 0; lineIndex < problemLastIndex; lineIndex++ {
			r := rune(lines[lineIndex][i])
			if unicode.IsSpace(r) {
				continue
			}

			currNum = currNum*10 + cast.ToInt(r)
		}
		currProblemNums = append(currProblemNums, currNum)

		if op := rune(lines[problemLastIndex][i]); !unicode.IsSpace(op) {
			res += slicy.Reduce(currProblemNums, getOperation(string(op)))
			currProblemNums = make([]int, 0)
			i-- // Skip blank problem separator
		}
	}

	return res
}

func getOperation(op string) func(int, int) int {
	operation := addFunc
	if op == "*" {
		operation = multFunc
	}
	return operation
}

var addFunc = func(a, b int) int {
	return a + b
}
var multFunc = func(a, b int) int {
	return a * b
}
