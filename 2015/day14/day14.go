package day14

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/stringy"
	"math"
	"regexp"
)

func part1(input string) int {
	reindeers := parseInput(input)
	totalTime := 2503

	maxDistance := math.MinInt
	for _, reindeer := range reindeers {
		cycleTime := reindeer.flyTime + reindeer.restTime
		nCycles, remainderTime := totalTime/cycleTime, totalTime%cycleTime
		distance := (nCycles*reindeer.flyTime + min(remainderTime, reindeer.flyTime)) * reindeer.speed
		maxDistance = max(maxDistance, distance)
	}

	return maxDistance
}

// Not my best code by far
// I don't like it one bit, but I hated part2 so it is what it is
func part2(input string) int {
	reindeers := parseInput(input)
	seconds := 2503

	for s := 0; s < seconds; s++ {
		maximumDistance := 0
		for _, reindeer := range reindeers {
			reindeer.tick()
			if reindeer.distance > maximumDistance {
				maximumDistance = reindeer.distance
			}
		}

		for _, reindeer := range reindeers {
			if reindeer.distance == maximumDistance {
				reindeer.points++
			}
		}
	}

	maximumPoints := -1
	for _, reindeer := range reindeers {
		if reindeer.points > maximumPoints {
			maximumPoints = reindeer.points
		}
	}

	return maximumPoints
}

func (r *reindeerInfo) tick() {
	if r.rest > 0 {
		r.rest--
	} else {
		r.flyDuration++
		r.distance += r.speed
	}

	if r.flyTime == r.flyDuration {
		r.flyDuration = 0
		r.rest = r.restTime
	}
}

type reindeerInfo struct {
	speed, flyTime, restTime            int
	rest, points, distance, flyDuration int // part2 specific fields
}

func parseInput(input string) []*reindeerInfo {
	inRegex := regexp.MustCompile(`^\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds\.$`)

	reindeers := make([]*reindeerInfo, 0)
	for line := range stringy.SplitLines(input) {
		matches := inRegex.FindAllStringSubmatch(line, -1)

		velocity := cast.ToInt(matches[0][1])
		flySeconds := cast.ToInt(matches[0][2])
		restSeconds := cast.ToInt(matches[0][3])
		reindeers = append(reindeers, &reindeerInfo{
			speed:    velocity,
			flyTime:  flySeconds,
			restTime: restSeconds,
		})
	}

	return reindeers
}
