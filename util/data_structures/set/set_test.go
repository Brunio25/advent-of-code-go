package set_test

import (
	"advent-of-code-go/util/data_structures/set"
	"slices"
	"testing"
)

func TestSetAdd(t *testing.T) {
	s := set.Set[string]{}
	s.Add("elem")
	if s.Size() != 1 || !s.Contains("elem") {
		t.Errorf("set.Add() = %v, Expected %v", s, []string{"elem"})
	}
}

func TestSetRemove(t *testing.T) {
	s := set.Set[int]{}
	s.Add(1)
	s.Add(2)
	s.Remove(1)
	if s.Size() != 1 || s.Contains(1) || !s.Contains(2) {
		t.Errorf("set.Remove() = %v, Expected %v", s, []int{2})
	}
}

func TestSetContains(t *testing.T) {
	s := set.Set[rune]{}
	s.Add('a')
	if actual := s.Contains('a'); !actual {
		t.Errorf("set.Contains() = %v, Expected %v", actual, true)
	}
	if actual := s.Contains('r'); actual {
		t.Errorf("set.Contains() = %v, Expected %v", actual, false)
	}
}

func TestSetSize(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := set.Set[int]{}
	for _, elem := range elems {
		s.Add(elem)
	}

	if actual := s.Size(); actual != 8 {
		t.Errorf("set.Size() = %v, Expected %v", actual, 8)
	}
}

func TestSetAll(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := set.Set[int]{}
	for _, elem := range elems {
		s.Add(elem)
	}

	actualElems := make([]int, 0, len(elems))
	for elem := range s.All() {
		actualElems = append(actualElems, elem)
	}
	slices.Sort(actualElems)

	if !slices.Equal(elems, actualElems) {
		t.Errorf("set.All() = %v, Expected %v", actualElems, elems)
	}
}
