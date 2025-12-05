package slicy

func CountFunc[S ~[]E, E comparable](slice S, predicate func(E) bool) int {
	count := 0
	for _, elem := range slice {
		if predicate(elem) {
			count++
		}
	}

	return count
}
