package grid_test

import (
	"advent-of-code-go/util/data_structures/grid"
	"testing"
)

func TestFromInputToRuneGrid(t *testing.T) {
	input := "ab\ncd"
	expected := [][]rune{{'a', 'b'}, {'c', 'd'}}

	actual := grid.FromInputToRuneGrid(input)
	if !areGridsEqual(actual, expected) {
		t.Errorf("grid.FromInputToRuneGrid() = %s, Expected = %v", actual, expected)
	}
}
