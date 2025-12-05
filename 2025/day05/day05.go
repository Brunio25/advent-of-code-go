package day05

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/slicy"
	"advent-of-code-go/util/stringy"
	"iter"
	"slices"
	"strings"
)

func part1(input string) int {
	freshIngredients, ingredientIds := parseInput(input)

	return slicy.CountFunc(ingredientIds, func(ingrId int) bool {
		return slices.ContainsFunc(freshIngredients, func(rng idRange) bool {
			return rng.start <= ingrId && ingrId <= rng.end
		})
	})
}

func part2(input string) int {
	freshIngredients, _ := parseInput(input)
	slices.SortFunc(freshIngredients, idRangeComparator)

	count := 0
	last := 0
	for _, rng := range freshIngredients {
		start := max(rng.start, last+1)
		count += max(0, rng.end-start+1)
		last = max(last, rng.end)
	}

	return count
}

func idRangeComparator(a, b idRange) int {
	switch {
	case a.start == b.start:
		return 0
	case a.start < b.start:
		return -1
	default:
		return 1
	}
}

func parseInput(input string) ([]idRange, []int) {
	idRanges := make([]idRange, 0)

	next, _ := iter.Pull(stringy.SplitLines(input))
	for {
		line, ok := next()
		if !ok || line == "" {
			break
		}

		splitRange := strings.Split(line, "-")
		idRanges = append(idRanges, idRange{
			start: cast.ToInt(splitRange[0]),
			end:   cast.ToInt(splitRange[1]),
		})
	}

	ingredients := make([]int, 0)
	for {
		line, ok := next()
		if !ok {
			break
		}

		ingredients = append(ingredients, cast.ToInt(line))
	}

	return idRanges, ingredients
}

type idRange struct {
	start int
	end   int
}
