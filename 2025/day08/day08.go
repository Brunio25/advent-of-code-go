package day08

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/data_structures/union_find"
	"advent-of-code-go/util/geom"
	"advent-of-code-go/util/slicy"
	"advent-of-code-go/util/stringy"
	"slices"
	"strings"
)

var nPairs = 1000 // For Example input, set this to 10

func part1(input string) int {
	sortedPointPairs := parseInputSorted(input)

	coordsToUf := createCoordsToUnionFind(sortedPointPairs)
	circuits := union_find.NewUnionFind(len(coordsToUf))

	for _, pair := range sortedPointPairs[:nPairs] {
		if circuits.SameSet(coordsToUf[pair.p1], coordsToUf[pair.p2]) {
			continue
		}

		circuits.Union(coordsToUf[pair.p1], coordsToUf[pair.p2])
	}

	circuitSizes := circuits.GetSizes()
	slices.Sort(circuitSizes)
	return slicy.Reduce(circuitSizes[len(circuitSizes)-3:], func(a int, b int) int {
		return a * b
	})
}

func part2(input string) int {
	sortedPointPairs := parseInputSorted(input)

	coordsToUf := createCoordsToUnionFind(sortedPointPairs)
	circuits := union_find.NewUnionFind(len(coordsToUf))

	for _, pair := range sortedPointPairs {
		if circuits.SameSet(coordsToUf[pair.p1], coordsToUf[pair.p2]) {
			continue
		}

		circuits.Union(coordsToUf[pair.p1], coordsToUf[pair.p2])
		if len(circuits.GetSizes()) == 1 {
			return pair.p1.X * pair.p2.X
		}
	}

	return -1
}

func createCoordsToUnionFind(sortedPairs []pointPair) map[geom.Coordinates3D]int {
	coordsToUf := make(map[geom.Coordinates3D]int)
	n := 0
	for _, pair := range sortedPairs {
		if _, inMap := coordsToUf[pair.p1]; !inMap {
			coordsToUf[pair.p1] = n
			n++
		}
		if _, inMap := coordsToUf[pair.p2]; !inMap {
			coordsToUf[pair.p2] = n
			n++
		}
	}
	return coordsToUf
}

func parseInputSorted(input string) []pointPair {
	points := make([]geom.Coordinates3D, 0)
	for line := range stringy.SplitLines(input) {
		splitLine := strings.Split(line, ",")
		points = append(points, geom.Coordinates3D{
			X: cast.ToInt(splitLine[0]),
			Y: cast.ToInt(splitLine[1]),
			Z: cast.ToInt(splitLine[2]),
		})
	}

	n := len(points)
	pairs := make([]pointPair, 0, (n*(n-1))/2)
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			if p1 == p2 {
				continue
			}

			pairs = append(pairs, pointPair{
				p1:   p1,
				p2:   p2,
				dist: p1.EuclideanDistance(p2),
			})
		}
	}

	slices.SortFunc(pairs, func(a, b pointPair) int {
		return int(a.dist - b.dist)
	})

	return pairs
}

type pointPair struct {
	p1, p2 geom.Coordinates3D
	dist   float64
}
