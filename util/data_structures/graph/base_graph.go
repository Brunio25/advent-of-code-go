package graph

import (
	"advent-of-code-go/util/data_structures/set"
	"fmt"
	"maps"
	"slices"
)

type NodeId interface {
	int | string | rune
}

type Interface[K NodeId] interface {
	AddNode(nodeId K)
	AddEdge(from, to K, weight int)
	HasEdge(from, to K) bool
	Neighbours(nodeId K) []K
	Weight(from, to K) int
}

type BaseGraph[K NodeId] map[K]set.Set[Edge[K]]

type Edge[K NodeId] struct {
	To     K
	Weight int
}

func (g BaseGraph[K]) AddNode(nodeId K) {
	g[nodeId] = set.Set[Edge[K]]{}
}

func (g BaseGraph[K]) AddNodeIfNotExists(nodeId K) {
	if _, ok := g[nodeId]; !ok {
		g.AddNode(nodeId)
	}
}

func (g BaseGraph[K]) HasEdge(from, to K) bool {
	for edge := range g[from].All() {
		if edge.To == to {
			return true
		}
	}
	return false
}

func (g BaseGraph[K]) Weight(from, to K) int {
	for edge := range g[from] {
		if edge.To == to {
			return edge.Weight
		}
	}
	panic(fmt.Sprintf("%v not an edge of %v", to, from))
}

func (g BaseGraph[K]) Neighbours(nodeId K) []K {
	edges := g[nodeId]
	res := make([]K, 0, len(edges))
	for edge := range edges.All() {
		res = append(res, edge.To)
	}
	return res
}

func (g BaseGraph[K]) Nodes() []K {
	return slices.Collect(maps.Keys(g))
}
