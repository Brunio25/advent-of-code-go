package grid_test

import (
	"advent-of-code-go/util/data_structures/grid"
	"advent-of-code-go/util/geom"
	"fmt"
	"slices"
	"testing"
)

func TestNewGridSuccess(t *testing.T) {
	var testCases = []struct {
		xSize int
		ySize int
	}{
		{10, 20},
		{5, 5},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`"X=%v,Y=%v`, tc.xSize, tc.ySize), func(t *testing.T) {
			if g := grid.NewGrid[int](tc.xSize, tc.ySize); len(g) != tc.ySize || len(g[0]) != tc.xSize {
				t.Errorf("grid.NewGrid() xLength = %v yLength = %v, Expected xLength = %v yLength = %v",
					len(g[0]), len(g), tc.xSize, tc.ySize)
			}
		})
	}
}

func TestNewGridInvalidInput(t *testing.T) {
	defer func() {
		t.Helper()
		expectedError := `grid cannot have negative or zero dimensions`
		if r := recover(); r == nil {
			t.Error("expected panic on grid.NewGrid(), but no panic occurred")
		} else if r != expectedError {
			t.Errorf(`expected panic on grid.NewGrid() with error "%s" but got "%s"`, expectedError, r)
		}
	}()
	grid.NewGrid[int](0, 0)
}

func TestNewGridDefaultSuccess(t *testing.T) {
	var testCases = []struct {
		xSize      int
		ySize      int
		defaultVal int
	}{
		{10, 20, 10},
		{5, 5, -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`"X=%v,Y=%v`, tc.xSize, tc.ySize), func(t *testing.T) {
			g := grid.NewGridDefault[int](tc.xSize, tc.ySize, tc.defaultVal)
			if len(g) != tc.ySize || len(g[0]) != tc.xSize && isGridFilledWithValue(g, tc.defaultVal) {
				t.Errorf("grid.NewGrid() xLength = %v yLength = %v, Expected xLength = %v yLength = %v",
					len(g[0]), len(g), tc.xSize, tc.ySize)
			}
		})
	}
}

func TestNewGridDefaultInvalidInput(t *testing.T) {
	defer func() {
		t.Helper()
		expectedError := `grid cannot have negative or zero dimensions`
		if r := recover(); r == nil {
			t.Error("expected panic on grid.NewGrid(), but no panic occurred")
		} else if r != expectedError {
			t.Errorf(`expected panic on grid.NewGrid() with error "%s" but got "%s"`, expectedError, r)
		}
	}()
	grid.NewGridDefault(0, 0, 3)
}

func TestValueAtCoordinates(t *testing.T) {
	value := 101
	g := grid.NewGrid[int](4, 4)
	g[2][3] = value
	if actual := g.ValueAtCoordinates(geom.Coordinates{X: 3, Y: 2}); actual != value {
		t.Errorf("grid.ValueAtCoordinates() = %v, Expected %v", actual, value)
	}
}

func TestValueAt(t *testing.T) {
	value := 101
	g := grid.NewGrid[int](4, 4)
	g[2][3] = value
	if actual := g.ValueAt(3, 2); actual != value {
		t.Errorf("grid.ValueAt() = %v, Expected %v", actual, value)
	}
}

func TestSetValueFromTo(t *testing.T) {
	g := grid.NewGrid[bool](7, 7)
	from, to := geom.Coordinates{X: 2, Y: 1}, geom.Coordinates{X: 5, Y: 4}
	g.SetValueFromTo(from, to, grid.RightDown, true)

	expectedTrueCoords := []geom.Coordinates{{X: 2, Y: 1}, {X: 3, Y: 2}, {X: 4, Y: 3}, {X: 5, Y: 4}}
	for y, row := range g {
		for x, val := range row {
			if val && !slices.Contains(expectedTrueCoords, geom.Coordinates{X: x, Y: y}) {
				t.Errorf("grid.SetValueFromTo() %v, Expected true %v", g, expectedTrueCoords)
			}
		}
	}
}

func TestSetValueFromToFunc(t *testing.T) {
	g := grid.NewGrid[bool](7, 7)
	from, to := geom.Coordinates{X: 2, Y: 1}, geom.Coordinates{X: 5, Y: 4}
	g.SetValueFromToFunc(from, to, grid.RightDown, func(b bool) bool { return true })

	expectedTrueCoords := []geom.Coordinates{{X: 2, Y: 1}, {X: 3, Y: 2}, {X: 4, Y: 3}, {X: 5, Y: 4}}
	for y, row := range g {
		for x, val := range row {
			if val && !slices.Contains(expectedTrueCoords, geom.Coordinates{X: x, Y: y}) {
				t.Errorf("grid.SetValueFromTo()\n%v\n\n-> Expected true %v", g, expectedTrueCoords)
			}
		}
	}
}

func TestIn(t *testing.T) {
	g := grid.NewGrid[int](2, 2)
	if g.In(geom.Coordinates{X: -1}) {
		t.Errorf("grid.In() = %v, Expected %v", true, false)
	}

	if !g.In(geom.Coordinates{X: 1, Y: 1}) {
		t.Errorf("grid.In() = %v, Expected %v", false, true)
	}
}

func TestForEach(t *testing.T) {
	g := grid.NewGrid[int](2, 2)
	g[0][0], g[0][1], g[1][0], g[1][1] = 0, 1, 2, 3
	expectedCoordsWithValues := []struct {
		c   geom.Coordinates
		val int
	}{
		{geom.Coordinates{X: 0, Y: 0}, 0},
		{geom.Coordinates{X: 1, Y: 0}, 1},
		{geom.Coordinates{X: 0, Y: 1}, 2},
		{geom.Coordinates{X: 1, Y: 1}, 3},
	}

	i := 0
	g.ForEach(func(x, y int, v int) {
		t.Helper()
		expected := expectedCoordsWithValues[i]
		if expected.c.X != x || expected.c.Y != y || expected.val != v {
			t.Errorf("grid.ForEach() = (Coordinates{%v, %v}, val = %v), Expected (Coordinates{%v, %v}, val = %v)",
				x, y, v, expected.c.X, expected.c.Y, expected.val)
			return
		}
		i++
	})
}

func TestIterator(t *testing.T) {
	g := grid.NewGrid[int](2, 2)
	g[0][0], g[0][1], g[1][0], g[1][1] = 0, 1, 2, 3
	expectedCoordsWithValues := []struct {
		c   geom.Coordinates
		val int
	}{
		{geom.Coordinates{X: 0, Y: 0}, 0},
		{geom.Coordinates{X: 1, Y: 0}, 1},
		{geom.Coordinates{X: 0, Y: 1}, 2},
		{geom.Coordinates{X: 1, Y: 1}, 3},
	}

	i := 0
	for c, v := range g.Iterator() {
		expected := expectedCoordsWithValues[i]
		if expected.c != c || expected.val != v {
			t.Errorf("grid.Iterator() = (Coordinates%v, val = %v), Expected (Coordinates%v, val = %v)",
				c, v, expected.c, expected.val)
			return
		}
		i++
	}
}

func TestCountFunc(t *testing.T) {
	g := grid.NewGrid[rune](2, 40)
	g[8][1] = 'a'
	g[38][0] = 'a'
	if actual := g.CountFunc(func(_ rune) bool { return true }); actual != 80 {
		t.Errorf("grid.CountFunc() = %v, Expected %v", actual, 80)
	}

	if actual := g.CountFunc(func(r rune) bool { return r == 'a' }); actual != 2 {
		t.Errorf("grid.CountFunc() = %v, Expected %v", actual, 2)
	}
}

func TestAdjacentCoords(t *testing.T) {
	g := grid.NewGrid[bool](5, 5)

	expectedFst := []geom.Coordinates{{1, 0}, {1, 1}, {0, 1}}
	if actual := g.AdjacentCoords(0, 0); !slices.Equal(actual, expectedFst) {
		t.Errorf("grid.AdjacentCoords() = %v, Expected %v", actual, expectedFst)
	}

	expectedSnd := []geom.Coordinates{
		{3, 2}, {4, 2}, {4, 3}, {4, 4},
		{3, 4}, {2, 4}, {2, 3}, {2, 2},
	}
	if actual := g.AdjacentCoords(3, 3); !slices.Equal(actual, expectedSnd) {
		t.Errorf("grid.AdjacentCoords() = %v, Expected %v", actual, expectedSnd)
	}

	expectedTrd := []geom.Coordinates{{4, 2}, {4, 4}, {3, 4}, {3, 3}, {3, 2}}
	if actual := g.AdjacentCoords(4, 3); !slices.Equal(actual, expectedTrd) {
		t.Errorf("grid.AdjacentCoords() = %v, Expected %v", actual, expectedTrd)
	}
}

func TestCopy(t *testing.T) {
	g := grid.NewGridDefault[int](3, 3, 10)
	g[2][2] = 71
	g[1][0] = 9

	actual := g.Copy()
	if !areGridsEqual(g, actual) {
		t.Errorf("grid.Copy() = %v, Expected %v", g, actual)
	}

	g[1][0] = -1
	if actual[1][0] == -1 {
		t.Errorf("grid.Copy() = %v, Expected %v", g, actual)
	}
}

// Helper Functions

func isGridFilledWithValue(g grid.Grid[int], v int) bool {
	for _, row := range g {
		for _, cell := range row {
			if cell != v {
				return false
			}
		}
	}
	return true
}

func areGridsEqual(g1 grid.Grid[int], g2 grid.Grid[int]) bool {
	for y, row := range g1 {
		if !slices.Equal(row, g2[y]) {
			return false
		}
	}
	return true
}
