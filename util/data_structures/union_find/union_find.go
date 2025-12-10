package union_find

type UnionFind struct {
	parent []int
	sizes  []int
}

// If at a certain point it's required to use this class with
// a set not being an int (like Coordinates2D), it may be useful
// to implement a wrapper class which does that exactly
//
// Check out the solution for day 08 2025 for how to do it:
// func createCoordsToUnionFind and coordsToUf[pair.p1]

func NewUnionFind(n int) UnionFind {
	parent := make([]int, n)
	sizes := make([]int, n)
	for i := range parent {
		parent[i] = i
		sizes[i] = 1
	}
	return UnionFind{parent: parent, sizes: sizes}
}

func (uf UnionFind) Parent(a int) int {
	return uf.parent[a]
}

func (uf UnionFind) Size(a int) int {
	return uf.sizes[uf.Find(a)]
}

func (uf UnionFind) Len() int {
	return len(uf.parent)
}

func (uf UnionFind) Find(a int) int {
	if uf.parent[a] == a {
		return a
	}

	return uf.Find(uf.parent[a])
}

func (uf UnionFind) Union(a, b int) {
	aRep := uf.Find(a)
	bRep := uf.Find(b)
	if aRep == bRep {
		return
	}

	if uf.sizes[aRep] < uf.sizes[bRep] {
		aRep, bRep = bRep, aRep
	}

	uf.parent[bRep] = aRep
	uf.sizes[aRep] += uf.sizes[bRep]
}

func (uf UnionFind) SameSet(a, b int) bool {
	return uf.Find(a) == uf.Find(b)
}

func (uf UnionFind) GetSizes() []int {
	res := make([]int, 0)
	for i, rep := range uf.parent {
		if i == rep {
			res = append(res, uf.sizes[rep])
		}
	}
	return res
}
