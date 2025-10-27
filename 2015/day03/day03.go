package day03

func part1(input string) int {
	isVisited := make(map[[2]int]bool)
	x, y := 0, 0

	visits := 1
	putIfAbsent(0, 0, isVisited)

	for _, m := range input {
		visits += moveSanta(&x, &y, m, isVisited)
	}

	return visits
}

func part2(input string) int {
	isVisited := make(map[[2]int]bool)
	x1, y1, x2, y2 := 0, 0, 0, 0

	visits := 1
	putIfAbsent(0, 0, isVisited)
	santaTurn := true
	for _, m := range input {
		if santaTurn {
			visits += moveSanta(&x1, &y1, m, isVisited)
		} else {
			visits += moveSanta(&x2, &y2, m, isVisited)
		}

		santaTurn = !santaTurn
	}

	return visits
}

func moveSanta(x, y *int, move rune, isVisited map[[2]int]bool) int {
	switch move {
	case '<':
		*x--
	case '>':
		*x++
	case '^':
		*y++
	case 'v':
		*y--
	}

	notVisited := putIfAbsent(*x, *y, isVisited)
	if notVisited {
		return 1
	}

	return 0
}

// Returns true if the key was absent and false if the key was present
func putIfAbsent(x, y int, isVisited map[[2]int]bool) bool {
	if _, isPresent := isVisited[[2]int{x, y}]; !isPresent {
		isVisited[[2]int{x, y}] = true
		return true
	}

	return false
}
