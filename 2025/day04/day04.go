package day04

import (
	"advent-of-code-go/util/data_structures/grid"
	"advent-of-code-go/util/geom"
	"advent-of-code-go/util/stringy"
)

func part1(input string) int {
	g := parseInput(input)
	return len(getAccessibleRolls(g))
}

func part2(input string) int {
	g := parseInput(input)

	removedRolls := 0
	for {
		accessibleRolls := getAccessibleRolls(g)
		if len(accessibleRolls) == 0 {
			break
		}

		for _, coord := range accessibleRolls {
			g.SetValueAtCoordinates(coord, false)
		}
		removedRolls += len(accessibleRolls)
	}

	return removedRolls
}

func getAccessibleRolls(g grid.Grid[bool]) []geom.Coordinates {
	accessibleRolls := make([]geom.Coordinates, 0)

	for coord, isRoll := range g.Iterator() {
		if !isRoll {
			continue
		}

		adjCoords := g.AdjacentCoords(coord.X, coord.Y)
		adjRolls := 0
		for _, adjCoord := range adjCoords {
			if g.ValueAtCoordinates(adjCoord) {
				adjRolls++
			}
		}
		if adjRolls < 4 {
			accessibleRolls = append(accessibleRolls, coord)
		}
	}

	return accessibleRolls
}

func parseInput(input string) grid.Grid[bool] {
	gridSize := getGridSize(input)
	g := make([][]bool, gridSize)

	y := 0
	for row := range stringy.SplitLines(input) {
		g[y] = make([]bool, gridSize)
		for x, r := range row {
			g[y][x] = r == '@'
		}
		y++
	}

	return g
}

func getGridSize(input string) int {
	for i, r := range input {
		if r == '\n' {
			return i
		}
	}

	return -1
}
