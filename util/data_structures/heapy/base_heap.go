package heapy

type BaseHeap[T any] struct {
	Nodes []T
}

func (h *BaseHeap[T]) Push(x any) {
	h.Nodes = append(h.Nodes, x.(T))
}

func (h *BaseHeap[T]) Pop() any {
	old := h.Nodes
	n := len(old)
	x := old[n-1]
	h.Nodes = old[0 : n-1]
	return x
}

func (h *BaseHeap[K]) Len() int {
	return len(h.Nodes)
}

func (h *BaseHeap[K]) Swap(i, j int) {
	h.Nodes[i], h.Nodes[j] = h.Nodes[j], h.Nodes[i]
}
