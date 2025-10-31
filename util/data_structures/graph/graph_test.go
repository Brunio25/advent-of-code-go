package graph_test

import (
	"advent-of-code-go/util/data_structures/graph"
	"advent-of-code-go/util/data_structures/set"
	"testing"
)

func TestUnidirectionalGraph(t *testing.T) {
	ug := graph.NewUnidirectionalGraph[int]()
	ug.AddNode(1)
	ug.AddNode(2)
	ug.AddEdge(1, 2, 10)

	expected := map[int]set.Set[graph.Edge[int]]{1: map[graph.Edge[int]]struct{}{graph.Edge[int]{To: 2, Weight: 10}: {}}}
	if !ug.HasEdge(1, 2) {
		t.Errorf("UnidirectionalGraph.AddEdge() = %v, Expected %v", ug, expected)
	}

	if ug.HasEdge(2, 1) {
		t.Errorf("UnidirectionalGraph.AddEdge() = %v, Expected %v", ug, expected)
	}
}

func TestBidirectionalGraph(t *testing.T) {
	ug := graph.NewBidirectionalGraph[int]()
	ug.AddNode(1)
	ug.AddNode(2)
	ug.AddEdge(1, 2, 10)

	expected := map[int]set.Set[graph.Edge[int]]{1: map[graph.Edge[int]]struct{}{graph.Edge[int]{To: 2, Weight: 10}: {}}}
	if !ug.HasEdge(1, 2) {
		t.Errorf("UnidirectionalGraph.AddEdge() = %v, Expected %v", ug, expected)
	}

	if !ug.HasEdge(2, 1) {
		t.Errorf("UnidirectionalGraph.AddEdge() = %v, Expected %v", ug, expected)
	}
}
