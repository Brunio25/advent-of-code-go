package grid

import (
	"advent-of-code-go/util/stringy"
	"strings"
)

func FromInputToRuneGrid(input string) Grid[rune] {
	linesCount := strings.Count(input, "\n") + 1
	g := make([][]rune, linesCount)

	y := 0
	for line := range stringy.SplitLines(input) {
		g[y] = []rune(line)
		y++
	}

	return g
}
