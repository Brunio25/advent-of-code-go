package day18

import (
	"advent-of-code-go/util/data_structures/grid"
	"advent-of-code-go/util/stringy"
)

type config struct {
	gridSide int
	steps    int
}

func getConfig() config {
	configs := struct {
		example config
		actual  config
	}{
		example: config{gridSide: 6, steps: 5},
		actual:  config{gridSide: 100, steps: 100},
	}

	return configs.actual
}

func part1(input string) int {
	conf := getConfig()
	g := parseInput(input, conf)
	return computeLightsOn(g, true, conf)
}

// Don't like having a boolean passed in to decide to do things differently
// but at this point this is good enough
func computeLightsOn(g grid.Grid[bool], part1 bool, conf config) int {
	for i := 0; i < conf.steps; i++ {
		tempG := g.Copy()

		g.ForEach(func(x, y int, on bool) {
			nAdjOn := 0
			for _, coords := range g.AdjacentCoords(x, y) {
				if g.ValueAtCoordinates(coords) {
					nAdjOn++
				}
			}

			tempG.SetValueAt(x, y, nAdjOn == 3 || (on && nAdjOn == 2))
		})

		if !part1 {
			tempG.SetValueAt(0, 0, true)
			tempG.SetValueAt(0, len(tempG)-1, true)
			tempG.SetValueAt(len(tempG)-1, 0, true)
			tempG.SetValueAt(len(tempG)-1, len(tempG)-1, true)
		}
		g = tempG
	}

	return g.CountFunc(func(on bool) bool { return on })
}

func part2(input string) int {
	conf := getConfig()
	g := parseInput(input, conf)
	g.SetValueAt(0, 0, true)
	g.SetValueAt(0, len(g)-1, true)
	g.SetValueAt(len(g)-1, 0, true)
	g.SetValueAt(len(g)-1, len(g)-1, true)
	return computeLightsOn(g, false, conf)
}

func parseInput(input string, conf config) grid.Grid[bool] {
	g := grid.NewGrid[bool](conf.gridSide, conf.gridSide)

	y := 0
	for line := range stringy.SplitLines(input) {
		for x, r := range line {
			if r == '#' {
				g[y][x] = true
			}
		}
		y++
	}

	return g
}
