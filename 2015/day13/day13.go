package day13

import (
	"advent-of-code-go/util/algorithms"
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/data_structures/graph"
	"advent-of-code-go/util/mathy"
	"advent-of-code-go/util/stringy"
	"math"
	"regexp"
)

func part1(input string) int {
	g := parseInput(input)
	return calcMaxHappiness(g)
}

func part2(input string) int {
	g := parseInput(input)
	meNode := "Me"
	for _, node := range g.Nodes() {
		g.AddEdge(meNode, node, 0)
		g.AddEdge(node, meNode, 0)
	}

	return calcMaxHappiness(g)
}

func calcMaxHappiness(g graph.UnidirectionalGraph[string]) int {
	maxHappiness := math.MinInt
	for _, arrangement := range algorithms.Permutations[string](g.Nodes()) {
		currentHappiness := 0
		lastNode := arrangement[0]
		for _, node := range append(arrangement[1:], arrangement[0]) {
			currentHappiness += g.Weight(lastNode, node) + g.Weight(node, lastNode)
			lastNode = node
		}
		maxHappiness = mathy.Max(maxHappiness, currentHappiness)
	}

	return maxHappiness
}

func parseInput(input string) graph.UnidirectionalGraph[string] {
	inRegex := regexp.MustCompile(`^(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+)\.$`)
	g := graph.NewUnidirectionalGraph[string]()

	for line := range stringy.SplitLines(input) {
		matches := inRegex.FindAllStringSubmatch(line, -1)

		amount := cast.ToInt(matches[0][3])
		if matches[0][2] == "lose" {
			amount = -1 * amount
		}

		g.AddEdge(matches[0][1], matches[0][4], amount)
	}

	return g
}
