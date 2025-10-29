package grid

func SumIntGrid(g Grid[int]) int {
	sum := 0
	g.ForEach(func(_, _ int, v int) {
		sum += v
	})
	return sum
}
