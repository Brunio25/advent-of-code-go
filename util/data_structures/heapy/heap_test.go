package heapy_test

import (
	"advent-of-code-go/util/data_structures/heapy"
	"container/heap"
	"math/rand"
	"slices"
	"testing"
)

func TestMinHeap(t *testing.T) {
	mh := &heapy.MinHeap[int]{}
	heap.Init(mh)
	expected := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		expected = append(expected, x)
		heap.Push(mh, x)
	}

	slices.Sort(expected)

	actualOrder := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		actualOrder = append(actualOrder, heap.Pop(mh).(int))
		if actualOrder[i] != expected[i] {
			t.Errorf("heapy.MinHeap() = %v, Expected %v", actualOrder[:i], expected[:i])
		}
	}
}

func TestMaxHeap(t *testing.T) {
	mh := &heapy.MaxHeap[int]{}
	heap.Init(mh)
	expected := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		expected = append(expected, x)
		heap.Push(mh, x)
	}

	slices.Sort(expected)
	slices.Reverse(expected)

	actualOrder := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		actualOrder = append(actualOrder, heap.Pop(mh).(int))
		if actualOrder[i] != expected[i] {
			t.Errorf("heapy.MaxHeap() = %v, Expected %v", actualOrder[:i], expected[:i])
		}
	}
}
