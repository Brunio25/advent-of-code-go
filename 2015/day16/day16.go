package day16

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/stringy"
	"regexp"
)

type tapeInfo struct {
	value    int
	checkVal func(int, int) bool
}

func equals(tapeInfoValue int, other int) bool {
	return tapeInfoValue == other
}

func greaterThan(tapeInfoValue int, other int) bool {
	return tapeInfoValue < other
}

func lesserThan(tapeInfoValue int, other int) bool {
	return tapeInfoValue > other
}

func part1(input string) int {
	sues := parseInput(input)

	tickerTape := map[string]tapeInfo{
		"children":    {value: 3, checkVal: equals},
		"cats":        {value: 7, checkVal: equals},
		"samoyeds":    {value: 2, checkVal: equals},
		"pomeranians": {value: 3, checkVal: equals},
		"akitas":      {value: 0, checkVal: equals},
		"vizslas":     {value: 0, checkVal: equals},
		"goldfish":    {value: 5, checkVal: equals},
		"trees":       {value: 3, checkVal: equals},
		"cars":        {value: 2, checkVal: equals},
		"perfumes":    {value: 1, checkVal: equals},
	}

	return findCorrectSue(sues, tickerTape)
}

func part2(input string) int {
	sues := parseInput(input)

	tickerTape := map[string]tapeInfo{
		"children":    {value: 3, checkVal: equals},
		"cats":        {value: 7, checkVal: greaterThan},
		"samoyeds":    {value: 2, checkVal: equals},
		"pomeranians": {value: 3, checkVal: lesserThan},
		"akitas":      {value: 0, checkVal: equals},
		"vizslas":     {value: 0, checkVal: equals},
		"goldfish":    {value: 5, checkVal: lesserThan},
		"trees":       {value: 3, checkVal: greaterThan},
		"cars":        {value: 2, checkVal: equals},
		"perfumes":    {value: 1, checkVal: equals},
	}

	return findCorrectSue(sues, tickerTape)
}

func findCorrectSue(sues []sue, tickerTape map[string]tapeInfo) int {
OUTER:
	for _, currSue := range sues {
		for _, entr := range currSue.entries {
			if val := tickerTape[entr.key]; !val.checkVal(val.value, entr.value) {
				continue OUTER
			}
		}
		return currSue.id

	}

	panic("Sue not found")
}

type sue struct {
	id      int
	entries []entry
}

type entry struct {
	key   string
	value int
}

func parseInput(input string) []sue {
	inRegex := regexp.MustCompile(`^Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)$`)

	sues := make([]sue, 0)
	for line := range stringy.SplitLines(input) {
		matches := inRegex.FindStringSubmatch(line)
		sues = append(sues, sue{
			id: cast.ToInt(matches[1]),
			entries: []entry{
				{key: matches[2], value: cast.ToInt(matches[3])},
				{key: matches[4], value: cast.ToInt(matches[5])},
				{key: matches[6], value: cast.ToInt(matches[7])},
			},
		})
	}

	return sues
}
