package mathy

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Max[T Number](nums ...T) T {
	maxNum := nums[0]
	for _, n := range nums {
		if n > maxNum {
			maxNum = n
		}
	}

	return maxNum
}

func Min[T Number](nums ...T) T {
	minNum := nums[0]
	for _, n := range nums {
		if n < minNum {
			minNum = n
		}
	}

	return minNum
}
