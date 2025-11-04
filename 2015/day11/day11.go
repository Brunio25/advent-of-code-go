package day11

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/data_structures/set"
	"strings"
)

var passwordValidators = []func(pass string) bool{
	func(pass string) bool {
		for i := 2; i < len(pass); i++ {
			if runeI := pass[i]; pass[i-1] == runeI-1 && pass[i-2] == runeI-2 {
				return true
			}
		}
		return false
	},
	func(pass string) bool {
		return !strings.ContainsAny(pass, "iol")
	},
	func(pass string) bool {
		pairs := set.Set[string]{}
		for i := 1; i < len(pass); i++ {
			if pass[i-1] == pass[i] {
				pairs.Add(pass[i-1 : i+1])
			}
		}

		return pairs.Size() >= 2
	},
}

func part1(input string) string {
	nextPassword := passwordGenerator(input)
	var pass string

	for !matchAllValidators(pass) {
		pass = nextPassword()
	}

	return pass
}

func part2(input string) string {
	nextPassword := passwordGenerator(input)
	var pass string
	validCount := 0

	for validCount != 2 {
		pass = nextPassword()
		if matchAllValidators(pass) {
			validCount++
		}
	}

	return pass
}

func passwordGenerator(input string) func() string {
	next := []byte(input)
	first := true
	return func() string {
		for i := len(next) - 1; !first && i >= 0; i-- {
			if next[i] == 'z' {
				next[i] = 'a'
			} else {
				next[i]++
				break
			}
		}

		if first {
			first = false
		}

		return cast.ToString(next)
	}
}

func matchAllValidators(pass string) bool {
	for _, validate := range passwordValidators {
		if !validate(pass) {
			return false
		}
	}
	return true
}
