package day01

func part1(input string) int {
	floor := 0
	for _, r := range input {
		if r == '(' {
			floor++
		} else {
			floor--
		}
	}

	return floor
}

func part2(input string) int {
	floor := 0
	for i, r := range input {
		if r == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}

	return -1
}
