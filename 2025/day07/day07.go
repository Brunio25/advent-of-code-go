package day07

import (
	"advent-of-code-go/util/data_structures/grid"
	"strings"
)

func part1(input string) int {
	splitCount, _ := simulateBeams(input)
	return splitCount
}

func part2(input string) int {
	_, timelines := simulateBeams(input)
	return timelines
}

func simulateBeams(input string) (int, int) {
	g := grid.FromInputToRuneGrid(input)
	rowLength := len(g[0])

	beams := make(map[int]int, rowLength)
	beams[strings.IndexRune(string(g[0]), 'S')] = 1

	splitCount := 0
	for y := 0; y < len(g); y++ {
		nextBeams := make(map[int]int, rowLength)
		for x, beamCount := range beams {
			if rune(g[y][x]) != '^' {
				nextBeams[x] += beamCount
				continue
			}

			nextBeams[x-1] += beamCount
			nextBeams[x+1] += beamCount
			splitCount++
		}
		beams = nextBeams
	}

	return splitCount, mapSum(beams)
}

func mapSum(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}
