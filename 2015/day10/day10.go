package day10

import (
	"advent-of-code-go/util/cast"
	"fmt"
	"strings"
)

func part1(input string) int {
	return lookAndSay(input, 40)
}

func part2(input string) int {
	return lookAndSay(input, 50)
}

func lookAndSay(input string, rounds int) int {
	for i := 0; i < rounds; i++ {
		said := strings.Builder{}
		groupLen := 0

		for j := 0; j < len(input); j++ {
			if j == len(input)-1 || input[j] != input[j+1] {
				said.WriteString(fmt.Sprintf("%d%s", groupLen+1, cast.ToString(input[j])))
				groupLen = 0
			} else {
				groupLen++
			}
		}
		input = said.String()
	}

	return len(input)
}
