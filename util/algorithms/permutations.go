package algorithms

import "slices"

func Permutations[T any](items []T) [][]T {
	return permutationsRec[T](items, len(items))
}

func permutationsRec[T any](items []T, size int) [][]T {
	if size == 1 {
		return [][]T{slices.Concat(items)}
	}

	perms := make([][]T, 0)
	for i := 0; i < size; i++ {
		perms = append(perms, permutationsRec[T](items, size-1)...)
		if size%2 == 0 {
			items[i], items[size-1] = items[size-1], items[i]
		} else {
			items[0], items[size-1] = items[size-1], items[0]
		}
	}
	return perms
}
