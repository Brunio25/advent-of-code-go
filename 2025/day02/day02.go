package day02

import (
	"advent-of-code-go/util/cast"
	"strings"

	"github.com/dlclark/regexp2"
)

func part1(input string) int {
	return sumInvalidIds(input, regexp2.MustCompile(`^(\w+)\1$`, regexp2.None))
}

func part2(input string) int {
	return sumInvalidIds(input, regexp2.MustCompile(`^(\w+)\1+$`, regexp2.None))
}

func sumInvalidIds(input string, invalidIdRegex *regexp2.Regexp) int {
	res := 0
	for rang := range strings.SplitSeq(input, ",") {
		splitRang := strings.Split(rang, "-")
		for i := cast.ToInt(splitRang[0]); i <= cast.ToInt(splitRang[1]); i++ {
			if isRepeat, _ := invalidIdRegex.MatchString(cast.ToString(i)); isRepeat {
				res += i
			}
		}
	}

	return res
}
