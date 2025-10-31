package graph_test

import (
	"advent-of-code-go/util/data_structures/graph"
	"advent-of-code-go/util/data_structures/set"
	"slices"
	"testing"
)

type testGraph struct {
	graph.BaseGraph[int]
}

func (tg testGraph) AddEdge(from, to int, weight int) {
	tg.AddNodeIfNotExists(from)
	edges := tg.BaseGraph[from]
	edges.Add(graph.Edge[int]{To: to, Weight: weight})
}

func TestBaseGraphAddNode(t *testing.T) {
	bg := make(graph.BaseGraph[int])
	bg.AddNode(1)

	if _, ok := bg[1]; !ok {
		expected := map[int]set.Set[graph.Edge[int]]{1: map[graph.Edge[int]]struct{}{}}
		t.Errorf("BaseGraph.AddNode() = %v, Expected %v", bg, expected)
	}
}

func TestBaseGraphAddNodeIfNotExists(t *testing.T) {
	bg := make(graph.BaseGraph[int])
	bg.AddNodeIfNotExists(1)
	if _, ok := bg[1]; !ok {
		expected := map[int]set.Set[graph.Edge[int]]{1: map[graph.Edge[int]]struct{}{}}
		t.Errorf("BaseGraph.AddNodeIfNotExists() = %v, Expected %v", bg, expected)
	}

	bg.AddNode(2)
	bg[1].Add(graph.Edge[int]{To: 2, Weight: 1})
	bg.AddNodeIfNotExists(1)
	if actual, ok := bg[1]; !ok { // If it adds node, it will replace the edges with an empty Set
		expected := map[int]set.Set[graph.Edge[int]]{1: map[graph.Edge[int]]struct{}{graph.Edge[int]{To: 2, Weight: 1}: {}}}
		t.Errorf("BaseGraph.AddNodeIfNotExists() = %v, Expected %v", actual, expected[1])
	}
}

func TestBaseGraphHasEdge(t *testing.T) {
	bg := testGraph{make(graph.BaseGraph[int])}
	bg.AddNode(1)
	bg.AddNode(2)
	if actual := bg.HasEdge(1, 2); actual {
		t.Errorf("BaseGraph.HasEdge() = %v, Expected %v", actual, false)
	}

	bg.AddEdge(1, 2, 5)
	if actual := bg.HasEdge(1, 2); !actual {
		t.Errorf("BaseGraph.HasEdge() = %v, Expected %v", actual, true)
	}
}

func TestBaseGraphWeightSuccess(t *testing.T) {
	bg := testGraph{make(graph.BaseGraph[int])}
	bg.AddNode(1)
	bg.AddNode(2)
	bg.AddEdge(1, 2, 5)
	if actual := bg.Weight(1, 2); actual != 5 {
		t.Errorf("BaseGraph.Weight() = %v, Expected %v", actual, 5)
	}
}

func TestBaseGraphWeightNonExistentEdge(t *testing.T) {
	defer func() {
		t.Helper()
		expectedError := "2 not an edge of 1"
		if r := recover(); r == nil {
			t.Error("expected panic on BaseGraph.Weight(), but no panic occurred")
		} else if r != expectedError {
			t.Errorf(`expected panic with error "%s", but got "%s"`, expectedError, r)
		}
	}()

	bg := testGraph{make(graph.BaseGraph[int])}
	bg.AddNode(1)
	bg.AddNode(2)
	bg.Weight(1, 2)
}

func TestBaseGraphNeighbours(t *testing.T) {
	bg := testGraph{make(graph.BaseGraph[int])}
	bg.AddNode(1)
	bg.AddNode(2)
	bg.AddNode(3)
	bg.AddEdge(1, 2, 2)
	bg.AddEdge(1, 3, 5)

	expectedNeighbours := []int{2, 3}
	if actual := bg.Neighbours(1); !slices.Equal(actual, expectedNeighbours) {
		t.Errorf("BaseGraph.Neighbours() = %v, Expected %v", actual, expectedNeighbours)
	}

	if actual := bg.Neighbours(2); len(actual) != 0 {
		t.Errorf("BaseGraph.Neighbours() = %v, Expected %v", actual, []int{})
	}
}

func TestBaseGraphNodes(t *testing.T) {
	bg := testGraph{make(graph.BaseGraph[int])}
	bg.AddNode(1)
	bg.AddNode(2)
	bg.AddNode(3)

	expected := []int{1, 2, 3}
	if actual := bg.Nodes(); !slices.Equal(actual, expected) {
		t.Errorf("BaseGraph.Nodes() = %v, Expected %v", actual, expected)
	}
}
