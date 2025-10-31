package day09

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/data_structures/graph"
	"advent-of-code-go/util/data_structures/set"
	"advent-of-code-go/util/mathy"
	"advent-of-code-go/util/stringy"
	"math"
	"strings"
)

func part1(input string) int {
	g := parseInput(input)

	minDist := math.MaxInt
	for _, startNode := range g.Nodes() {
		dfsRes, _ := dfs(g, startNode, map[string]struct{}{startNode: {}})
		minDist = mathy.Min(minDist, dfsRes)
	}
	return minDist
}

func part2(input string) int {
	g := parseInput(input)

	maxDist := math.MinInt
	for _, startNode := range g.Nodes() {
		_, dfsRes := dfs(g, startNode, map[string]struct{}{startNode: {}})
		maxDist = mathy.Max(maxDist, dfsRes)
	}
	return maxDist
}

func dfs(g graph.BidirectionalGraph[string], start string, visits set.Set[string]) (min, max int) {
	if len(visits) == len(g.Nodes()) {
		return 0, 0
	}

	minDist := math.MaxInt
	maxDist := math.MinInt
	for _, destNode := range g.Neighbours(start) {
		if visits.Contains(destNode) {
			continue
		}

		visits.Add(destNode)
		weight := g.Weight(start, destNode)
		minDfs, maxDfs := dfs(g, destNode, visits)
		minDist = mathy.Min(minDist, weight+minDfs)
		maxDist = mathy.Max(maxDist, weight+maxDfs)

		delete(visits, destNode)
	}

	return minDist, maxDist
}

func parseInput(input string) graph.BidirectionalGraph[string] {
	g := graph.NewBidirectionalGraph[string]()
	for line := range stringy.SplitLines(input) {
		edgeAndWeight := strings.Split(line, " = ")
		fromAndTo := strings.Split(edgeAndWeight[0], " to ")
		g.AddEdge(fromAndTo[0], fromAndTo[1], cast.ToInt(edgeAndWeight[1]))
	}
	return g
}
