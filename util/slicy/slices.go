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

func Fold[E any, R any](slice []E, acc R, operation func(R, E) R) R {
	for _, e := range slice {
		acc = operation(acc, e)
	}
	return acc
}

func Reduce[E any](slice []E, operation func(E, E) E) E {
	if len(slice) == 0 {
		panic("empty collection can't be reduced")
	}

	return Fold(slice[1:], slice[0], operation)
}
