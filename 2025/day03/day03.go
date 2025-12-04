package day03

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/stringy"
)

func part1(input string) int {
	return maxJoltage(input, 2)
}

func part2(input string) int {
	return maxJoltage(input, 12)
}

func maxJoltage(input string, nBatteries int) int {
	totalJoltage := 0
	for bank := range stringy.SplitLines(input) {
		bankLen := len(bank)
		currJoltage, maxIndex := 0, -1

		for pickedBatteries := 0; pickedBatteries < nBatteries; pickedBatteries++ {
			maxIndex++
			offset := nBatteries - pickedBatteries - 1
			maxIndex = maxIndexIn(bank, maxIndex, bankLen-offset)
			currJoltage = 10*currJoltage + cast.ToInt(bank[maxIndex])
		}
		totalJoltage += currJoltage
	}

	return totalJoltage
}

func maxIndexIn(batteries string, start, end int) int {
	maxIndex := start
	for i := start; i < end; i++ {
		if batteries[i] > batteries[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}
