package set

import (
	"fmt"
	"iter"
	"maps"
)

type Set[E comparable] map[E]struct{}

func (s Set[E]) Add(e E) {
	s[e] = struct{}{}
}

func (s Set[E]) Remove(e E) {
	delete(s, e)
}

func (s Set[E]) Contains(e E) bool {
	_, ok := s[e]
	return ok
}

func (s Set[E]) Size() int {
	return len(s)
}

func (s Set[E]) All() iter.Seq[E] {
	return maps.Keys(s)
}

func (s Set[E]) String() string {
	res := make([]E, 0, s.Size())
	for elem := range maps.Keys(s) {
		res = append(res, elem)
	}

	return fmt.Sprint(res)
}
