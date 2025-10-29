package grid_test

import (
	"advent-of-code-go/util/data_structures/grid"
	"testing"
)

func TestSumIntGrid(t *testing.T) {
	g := grid.NewGrid[int](2, 2)
	g[0][0], g[0][1], g[1][0], g[1][1] = 101, -1, 50, 34
	expected := 184
	if actual := grid.SumIntGrid(g); actual != expected {
		t.Errorf("grid.SumIntGrid() = %v, Expected %v", actual, expected)
	}

}
