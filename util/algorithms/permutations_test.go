package algorithms_test

import (
	"advent-of-code-go/util/algorithms"
	"slices"
	"testing"
)

func TestPermutationsInt(t *testing.T) {
	actual := algorithms.Permutations[int]([]int{1, 2, 3})
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	if len(actual) != len(expected) {
		t.Errorf("algorithms.Permutations() = %v, Expected %v", actual, expected)
	}

	for _, perm := range actual {
		if !sliceContains(expected, perm) {
			t.Errorf("algorithms.Permutations():  permutation %v expected to be in %v", perm, expected)
		}
	}
}

func TestPermutationsString(t *testing.T) {
	actual := algorithms.Permutations[string]([]string{"Ines", "John", "Matt", "Oscar"})

	if len(actual) != len(expectedStringPerms) {
		t.Errorf("algorithms.Permutations() = %v, Expected %v", actual, expectedStringPerms)
	}

	for _, perm := range actual {
		if !sliceContains(expectedStringPerms, perm) {
			t.Errorf("algorithms.Permutations():  permutation %v expected to be in %v", perm, expectedStringPerms)
		}
	}
}

func sliceContains[T comparable](items [][]T, v []T) bool {
	return slices.ContainsFunc(items, func(item []T) bool {
		return slices.Equal(item, v)
	})
}

var expectedStringPerms = [][]string{
	{"Ines", "John", "Matt", "Oscar"},
	{"Ines", "John", "Oscar", "Matt"},
	{"Ines", "Matt", "John", "Oscar"},
	{"Ines", "Matt", "Oscar", "John"},
	{"Ines", "Oscar", "John", "Matt"},
	{"Ines", "Oscar", "Matt", "John"},
	{"John", "Ines", "Matt", "Oscar"},
	{"John", "Ines", "Oscar", "Matt"},
	{"John", "Matt", "Ines", "Oscar"},
	{"John", "Matt", "Oscar", "Ines"},
	{"John", "Oscar", "Ines", "Matt"},
	{"John", "Oscar", "Matt", "Ines"},
	{"Matt", "Ines", "John", "Oscar"},
	{"Matt", "Ines", "Oscar", "John"},
	{"Matt", "John", "Ines", "Oscar"},
	{"Matt", "John", "Oscar", "Ines"},
	{"Matt", "Oscar", "Ines", "John"},
	{"Matt", "Oscar", "John", "Ines"},
	{"Oscar", "Ines", "John", "Matt"},
	{"Oscar", "Ines", "Matt", "John"},
	{"Oscar", "John", "Ines", "Matt"},
	{"Oscar", "John", "Matt", "Ines"},
	{"Oscar", "Matt", "Ines", "John"},
	{"Oscar", "Matt", "John", "Ines"},
}
