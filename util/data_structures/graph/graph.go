package graph

type UnidirectionalGraph[K NodeId] struct {
	BaseGraph[K]
}

func (ug UnidirectionalGraph[K]) AddEdge(from, to K, weight int) {
	ug.AddNodeIfNotExists(from)

	edges := ug.BaseGraph[from]
	edges.Add(Edge[K]{To: to, Weight: weight})
}

func NewUnidirectionalGraph[K NodeId]() UnidirectionalGraph[K] {
	return UnidirectionalGraph[K]{
		BaseGraph: make(BaseGraph[K]),
	}
}

type BidirectionalGraph[K NodeId] struct {
	BaseGraph[K]
}

func (bg BidirectionalGraph[K]) AddEdge(from, to K, weight int) {
	bg.AddNodeIfNotExists(from)
	bg.AddNodeIfNotExists(to)

	bg.BaseGraph[from].Add(Edge[K]{To: to, Weight: weight})
	bg.BaseGraph[to].Add(Edge[K]{To: from, Weight: weight})
}

func NewBidirectionalGraph[K NodeId]() BidirectionalGraph[K] {
	return BidirectionalGraph[K]{
		BaseGraph: make(BaseGraph[K]),
	}
}
