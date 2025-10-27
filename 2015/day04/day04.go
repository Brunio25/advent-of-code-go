package day04

import (
	"crypto/md5"
	"fmt"
	"math"
	"strings"
)

func part1(input string) int {
	return smallestNToComputeValidHash(input, 5)
}

func part2(input string) int {
	return smallestNToComputeValidHash(input, 6)
}

func smallestNToComputeValidHash(input string, prefixZeroesN int) int {
	for i := 0; i < math.MaxInt64; i++ {
		toHash := fmt.Sprintf("%s%d", input, i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
		if strings.HasPrefix(hash, strings.Repeat("0", prefixZeroesN)) {
			return i
		}
	}

	panic("no valid hash found")
}
