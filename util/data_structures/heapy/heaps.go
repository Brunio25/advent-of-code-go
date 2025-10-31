package heapy

import (
	"cmp"
)

type MinHeap[T cmp.Ordered] struct {
	BaseHeap[T]
}

func (mh *MinHeap[T]) Less(i, j int) bool {
	return mh.Nodes[i] < mh.Nodes[j]
}

type MaxHeap[T cmp.Ordered] struct {
	BaseHeap[T]
}

func (mh *MaxHeap[T]) Less(i, j int) bool {
	return mh.Nodes[i] > mh.Nodes[j]
}
