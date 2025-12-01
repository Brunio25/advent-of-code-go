package day01

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/mathy"
	"advent-of-code-go/util/stringy"
)

func part1(input string) int {
	currPos := 50
	timesAtZero := 0
	for line := range stringy.SplitLines(input) {
		inst := cast.ToInt(line[1:])
		if line[0] == 'L' {
			currPos = mathy.Mod(currPos-inst, 100)
		} else {
			currPos = mathy.Mod(currPos+inst, 100)
		}

		if currPos == 0 {
			timesAtZero++
		}
	}
	return timesAtZero
}

func part2(input string) int {
	currPos := 50
	timesAtZero := 0
	for line := range stringy.SplitLines(input) {
		inst := cast.ToInt(line[1:])

		if line[0] == 'L' {
			timesAtZero += countZeros(currPos-inst, currPos-1)
			inst = inst * -1
		} else {
			timesAtZero += countZeros(currPos+1, currPos+inst)
		}
		currPos = mathy.Mod(currPos+inst, 100)
	}
	return timesAtZero
}

func countZeros(smaller, larger int) int {
	if smaller > larger {
		return 0
	}
	return mathy.FloorDiv(larger, 100) - mathy.FloorDiv(smaller-1, 100)
}
