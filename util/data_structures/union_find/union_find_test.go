package union_find_test

import (
	"advent-of-code-go/util/data_structures/union_find"
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	expectedLength := 10
	uf := union_find.NewUnionFind(expectedLength)

	if actualLen := uf.Len(); actualLen != expectedLength {
		t.Errorf("union_find.NewUnionFind() Length = %v, Expected %v", actualLen, expectedLength)
	}

	for i := 0; i < uf.Len(); i++ {
		if actualParent := uf.Parent(i); actualParent != i {
			t.Errorf("union_find.NewUnionFind() Parent = %v, Expected %v", actualParent, i)
		}

		if actualSize := uf.Size(i); actualSize != 1 {
			t.Errorf("union_find.NewUnionFind() Size = %v, Expected %v", actualSize, i)
		}
	}
}

func TestUnionFind_Find(t *testing.T) {
	expected := 3
	uf := union_find.NewUnionFind(5)
	if actual := uf.Find(expected); actual != expected {
		t.Errorf("union_find.Find() = %v, Expected %v", actual, expected)
	}
}

// Wrote this massive test function because didn't feel like writing proper unit tests
func TestUnionFind(t *testing.T) {
	expectedParent := 1
	expectedSize := 4

	uf := union_find.NewUnionFind(5)
	uf.Union(1, 2)
	uf.Union(3, 2)
	uf.Union(3, 4)

	if actual := uf.Find(0); actual != 0 {
		t.Errorf("union_find.Find() = %v, Expected %v", actual, 0)
	}

	if actual := uf.Size(0); actual != 1 {
		t.Errorf("union_find.Find() Size = %v, Expected %v", actual, 1)
	}

	for i := 1; i < 5; i++ {
		if actual := uf.Find(i); actual != expectedParent {
			t.Errorf("union_find.Find() = %v, Expected %v", actual, expectedParent)
		}

		if actual := uf.Size(i); actual != expectedSize {
			t.Errorf("union_find.Find() Size = %v, Expected %v", actual, expectedSize)
		}
	}

	if actual := uf.SameSet(1, 4); !actual {
		t.Errorf("union_find.SameSet() = %v, Expected %v", actual, true)
	}

	actualSizes := uf.GetSizes()
	if actualSizes[0] != 1 || actualSizes[1] != expectedSize {
		t.Errorf("union_find.GetSizes() = %v, Expected %v", actualSizes, expectedSize)
	}
}
