package day06

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/data_structures/grid"
	"advent-of-code-go/util/geom"
	"iter"
	"regexp"
	"slices"
	"strings"
)

func part1(input string) int {
	g := grid.NewGrid[bool](1000, 1000)
	for inst := range parseInput(input) {
		setGridValuesSquareShape(inst.from, inst.to, g, actionToFuncPart1(inst.action))
	}

	return g.CountFunc(func(val bool) bool {
		return val == true
	})
}

func actionToFuncPart1(action Action) func(bool) bool {
	switch action {
	case TurnOn:
		return func(_ bool) bool { return true }
	case TurnOff:
		return func(_ bool) bool { return false }
	default:
		return func(b bool) bool { return !b }
	}
}

func part2(input string) int {
	g := grid.NewGrid[int](1000, 1000)
	for inst := range parseInput(input) {
		setGridValuesSquareShape(inst.from, inst.to, g, actionToFuncPart2(inst.action))
	}

	return grid.SumIntGrid(g)
}

func actionToFuncPart2(action Action) func(int) int {
	switch action {
	case TurnOn:
		return func(bright int) int {
			return bright + 1
		}
	case TurnOff:
		return func(bright int) int {
			if bright == 0 {
				return 0
			}
			return bright - 1
		}
	default:
		return func(bright int) int {
			return bright + 2
		}
	}
}

func setGridValuesSquareShape[T grid.CellValue](from, to geom.Coordinates, g grid.Grid[T], transform func(T) T) {
	ys, xs := []int{from.Y, to.Y}, []int{from.X, to.X}
	slices.Sort(ys)
	slices.Sort(xs)

	for y := ys[0]; y <= ys[1]; y++ {
		g.SetValueFromToFunc(geom.Coordinates{X: xs[0], Y: y},
			geom.Coordinates{X: xs[1], Y: y},
			grid.Right,
			transform)
	}
}

type Instruction struct {
	action Action
	from   geom.Coordinates
	to     geom.Coordinates
}

type Action int

const (
	TurnOn Action = iota
	TurnOff
	Toggle
)

// Input Parsing Functions

func parseInput(input string) iter.Seq[Instruction] {
	return func(yield func(Instruction) bool) {
		for line := range strings.Lines(input) {
			line = strings.TrimRight(line, "\n")

			startCoords, endCoords := getCoordinates(line)

			inst := Instruction{
				action: getAction(line),
				from:   startCoords,
				to:     endCoords,
			}

			if !yield(inst) {
				return
			}
		}
	}
}

var extractCoordsPattern = regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)`)

func getCoordinates(line string) (start, end geom.Coordinates) {
	matches := extractCoordsPattern.FindStringSubmatch(line)
	start = geom.Coordinates{X: cast.ToInt(matches[1]), Y: cast.ToInt(matches[2])}
	end = geom.Coordinates{X: cast.ToInt(matches[3]), Y: cast.ToInt(matches[4])}
	return
}

func getAction(line string) Action {
	switch {
	case strings.HasPrefix(line, "turn on"):
		return TurnOn
	case strings.HasPrefix(line, "turn off"):
		return TurnOff
	default:
		return Toggle
	}
}
