package heapy_test

import (
	"advent-of-code-go/util/data_structures/heapy"
	"container/heap"
	"testing"
)

type tuple struct {
	fst int
	snd int
}

type tupleHeap struct {
	heapy.BaseHeap[tuple]
}

func (th *tupleHeap) Less(i, j int) bool {
	return th.Nodes[i].snd < th.Nodes[j].snd
}

var tuplesInput = []tuple{
	{10, 1},
	{1000, -101},
	{0, 0},
	{101, 19},
	{-3, 30},
}

var tuplesExpected = []tuple{
	{1000, -101},
	{0, 0},
	{10, 1},
	{101, 19},
	{-3, 30},
}

var tuplesExpectedAfterRemove = []tuple{
	{1000, -101},
	{0, 0},
	{101, 19},
	{-3, 30},
}

func TestBaseHeap(t *testing.T) {
	th := initTupleHeap(t)

	// Validate Heap is sorted
	validateTupleHeap(th, tuplesExpected, t)

	th = initTupleHeap(t)
	heap.Remove(th, 1)
	validateTupleHeap(th, tuplesExpectedAfterRemove, t)
}

func initTupleHeap(t *testing.T) *tupleHeap {
	t.Helper()
	th := &tupleHeap{}
	heap.Init(th)

	for _, tupl := range tuplesInput {
		heap.Push(th, tupl)
	}

	return th
}

func validateTupleHeap(th *tupleHeap, expected []tuple, t *testing.T) {
	t.Helper()
	i := 0
	var actualHeap []tuple
	for th.Len() != 0 {
		popTupl := heap.Pop(th).(tuple)
		actualHeap = append(actualHeap, popTupl)
		if popTupl != expected[i] {
			t.Errorf("heap.BaseHeap() = %v, Expected %v", actualHeap, expected[:i])
		}
		i++
	}
}
