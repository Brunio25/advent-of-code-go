package day09

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/geom"
	"advent-of-code-go/util/mathy"
	"advent-of-code-go/util/stringy"
	"slices"
	"strings"
)

func part1(input string) int {
	redTiles := parseInput(input)

	maxArea := 0
	for i, redTile1 := range redTiles {
		for _, redTile2 := range redTiles[i+1:] {
			if area := areaBetweenCoords(redTile1, redTile2); area > maxArea {
				maxArea = area
			}

		}
	}

	return maxArea
}

func areaBetweenCoords(a, b geom.Coordinates2D) int {
	return mathy.Abs(a.X-b.X+1) * mathy.Abs(a.Y-b.Y+1)
}

func part2(input string) int {
	redTiles := parseInput(input)
	polygon := slices.Concat(redTiles, redTiles[:1])
	//g := grid.NewGridDefault()

	return len(polygon)
}

func parseInput(input string) []geom.Coordinates2D {
	redTiles := make([]geom.Coordinates2D, 0)
	for line := range stringy.SplitLines(input) {
		splitLine := strings.Split(line, ",")
		redTiles = append(redTiles, geom.Coordinates2D{
			X: cast.ToInt(splitLine[0]),
			Y: cast.ToInt(splitLine[1]),
		})
	}
	return redTiles
}
