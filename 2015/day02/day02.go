package day02

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/mathy"
	"regexp"
	"strings"
)

func part1(input string) int {
	boxes := parseInput(input)
	requiredPaper := 0
	for _, b := range boxes {
		lw := b.l * b.w
		wh := b.w * b.h
		hl := b.h * b.l

		requiredPaper += 2*lw + 2*wh + 2*hl + mathy.Min(lw, wh, hl)
	}

	return requiredPaper
}

func part2(input string) int {
	boxes := parseInput(input)
	requiredPaper := 0
	for _, b := range boxes {
		lw := 2*b.l + 2*b.w
		wh := 2*b.w + 2*b.h
		hl := 2*b.h + 2*b.l

		requiredPaper += mathy.Min(lw, wh, hl) + (b.l * b.w * b.h)
	}

	return requiredPaper
}

type boxDimensions struct {
	l int
	w int
	h int
}

func parseInput(input string) []boxDimensions {
	reg := regexp.MustCompile(`^(\d+)x(\d+)x(\d+)`)

	res := make([]boxDimensions, 0, 15)
	for line := range strings.Lines(input) {
		matches := reg.FindStringSubmatch(line)
		res = append(res, boxDimensions{
			l: cast.ToInt(matches[1]),
			w: cast.ToInt(matches[2]),
			h: cast.ToInt(matches[3]),
		})
	}

	return res
}
