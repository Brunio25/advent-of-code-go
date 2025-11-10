package day15

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/stringy"
	"iter"
	"regexp"
	"slices"
)

func part1(input string) int {
	ingredients := parseInput(input)
	return bestRecipe(ingredients, func(calories int, score int) int { return score })
}

func part2(input string) int {
	ingredients := parseInput(input)
	return bestRecipe(ingredients, func(calories int, score int) int {
		if calories == 500 {
			return score
		}
		return 0
	})
}

func bestRecipe(ingredients []ingredient, caloryMapper func(int, int) int) int {
	maxRecipe := 0
	for amounts := range amountsGenerator(len(ingredients), 100) {
		capacity, durability, flavor, texture, calories := 0, 0, 0, 0, 0
		for i, quantity := range amounts {
			capacity += quantity * ingredients[i].capacity
			durability += quantity * ingredients[i].durability
			flavor += quantity * ingredients[i].flavor
			texture += quantity * ingredients[i].texture
			calories += quantity * ingredients[i].calories
		}
		capacity, durability, flavor, texture = max(0, capacity), max(0, durability), max(0, flavor), max(0, texture)
		maxRecipe = max(maxRecipe, caloryMapper(calories, capacity*durability*flavor*texture))
	}

	return maxRecipe
}

func amountsGenerator(ingredientsCount int, teaspoons int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		if ingredientsCount == 1 {
			yield([]int{teaspoons})
			return
		}

		for i := 0; i < teaspoons; i++ {
			for rest := range amountsGenerator(ingredientsCount-1, teaspoons-i) {
				if !yield(slices.Concat([]int{i}, rest)) {
					return
				}
			}
		}
	}
}

type ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseInput(input string) []ingredient {
	inRegex := regexp.MustCompile(`^\w+: capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)$`)

	ingredients := make([]ingredient, 0)
	for line := range stringy.SplitLines(input) {
		matches := inRegex.FindAllStringSubmatch(line, -1)
		ingredients = append(ingredients, ingredient{
			capacity:   cast.ToInt(matches[0][1]),
			durability: cast.ToInt(matches[0][2]),
			flavor:     cast.ToInt(matches[0][3]),
			texture:    cast.ToInt(matches[0][4]),
			calories:   cast.ToInt(matches[0][5]),
		})
	}

	return ingredients
}
