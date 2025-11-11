package day17

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/stringy"
	"maps"
	"math"
)

var target = 150

// I could clean up this file so that part1 uses the same approach as part2,
// but I don't feel like it right now. Also, this way I can later glance at what
// works best for these two different requirements

func part1(input string) int {
	containers := parseInput(input)
	return possibleCombs(containers)
}

func possibleCombs(containers []int) int {
	mem := make([]int, target+1)
	mem[0] = 1

	for _, container := range containers {
		for i := target; i >= container; i-- {
			mem[i] += mem[i-container]
		}
	}

	return mem[target]
}

func part2(input string) int {
	containers := parseInput(input)

	var results = make([][]int, 0)
	findValidContainerCombs(0, 0, []int{}, containers, &results)

	lenToCount := make(map[int]int)
	for _, row := range results {
		if val, ok := lenToCount[len(row)]; ok {
			lenToCount[len(row)] = val + 1
			continue
		}
		lenToCount[len(row)] = 1
	}

	smallestN := math.MaxInt
	for k := range maps.Keys(lenToCount) {
		if smallestN > k {
			smallestN = k
		}
	}

	return lenToCount[smallestN]
}

func findValidContainerCombs(i, currentSum int, path []int, containers []int, result *[][]int) {
	if currentSum == target {
		*result = append(*result, path)
		return
	}

	if currentSum > target || i == len(containers) {
		return
	}

	findValidContainerCombs(i+1, currentSum+containers[i], append(path, containers[i]), containers, result)
	findValidContainerCombs(i+1, currentSum, path, containers, result)
}

func parseInput(input string) []int {
	containers := make([]int, 0)
	for line := range stringy.SplitLines(input) {
		containers = append(containers, cast.ToInt(line))
	}
	return containers
}
