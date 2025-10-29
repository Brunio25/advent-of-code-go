package grid

import (
	"advent-of-code-go/util/geom"
	"fmt"
	"iter"
	"strings"
)

type Grid[T CellValue] [][]T

type CellValue interface {
	int | rune | bool
}

func NewGrid[T CellValue](xSize, ySize int) Grid[T] {
	if xSize <= 0 || ySize <= 0 {
		panic("grid cannot have negative or zero dimensions")
	}

	g := make(Grid[T], ySize)
	for y := range g {
		g[y] = make([]T, xSize)
	}
	return g
}

func NewGridDefault[T CellValue](xSize, ySize int, defaultValue T) Grid[T] {
	if xSize <= 0 || ySize <= 0 {
		panic("grid cannot have negative or zero dimensions")
	}

	g := make(Grid[T], ySize)
	for y := range g {
		g[y] = make([]T, xSize)
		for x := range g[y] {
			g[y][x] = defaultValue
		}
	}
	return g
}

func (g Grid[T]) ValueAtCoordinates(c geom.Coordinates) T {
	return g[c.Y][c.X]
}

func (g Grid[T]) ValueAt(x, y int) T {
	return g[y][x]
}

func (g Grid[T]) SetValueFromTo(from, to geom.Coordinates, dir Direction, value T) {
	g.SetValueFromToFunc(from, to, dir, func(_ T) T { return value })
}

func (g Grid[T]) SetValueFromToFunc(from, to geom.Coordinates, dir Direction, transform func(T) T) {
	for c := range coordsBetween(from, to, dir) {
		g[c.Y][c.X] = transform(g.ValueAtCoordinates(c))
	}
}

func coordsBetween(from, to geom.Coordinates, dir Direction) iter.Seq[geom.Coordinates] {
	return func(yield func(geom.Coordinates) bool) {
		lastCoord := from
		if !yield(lastCoord) {
			return
		}

		for lastCoord.X != to.X || lastCoord.Y != to.Y {
			lastCoord = dir(lastCoord)
			if !yield(lastCoord) {
				return
			}
		}
	}
}

func (g Grid[T]) ForEach(fn func(x, y int, v T)) {
	for y := range g {
		for x, v := range g[y] {
			fn(x, y, v)
		}
	}
}

func (g Grid[T]) CountFunc(predicate func(T) bool) int {
	count := 0
	g.ForEach(func(_, _ int, v T) {
		if predicate(v) {
			count++
		}
	})
	return count
}

// Rudimentary print function
// Ideally, it would auto adjust spaces between elements (based on cell value length) TODO
func (g Grid[T]) String() string {
	strGrid := make([]string, len(g))
	for y, row := range g {
		elems := make([]string, len(row))
		for x, v := range row {
			elems[x] = fmt.Sprint(v)
		}
		strGrid[y] = strings.Join(elems, " ")
	}

	return strings.Join(strGrid, "\n")
}
