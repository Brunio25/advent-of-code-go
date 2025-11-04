package day12

import (
	"advent-of-code-go/util/cast"
	"encoding/json"
	"regexp"
)

var numRegex = regexp.MustCompile(`-?\d+`)

func part1(input string) int {
	strNums := numRegex.FindAllString(input, -1)
	sum := 0
	for _, strNum := range strNums {
		sum += cast.ToInt(strNum)
	}
	return sum
}

func part2(input string) int {
	var data any
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		panic("invalid json")
	}

	return jsonSum(data)
}

func jsonSum(parent any) int {
	switch v := parent.(type) {
	case float64:
		return int(v)
	case []any:
		sum := 0
		for _, elem := range v {
			sum += jsonSum(elem)
		}
		return sum
	case map[string]any:
		for _, val := range v {
			if str, ok := val.(string); ok && str == "red" {
				return 0
			}
		}
		sum := 0
		for _, elem := range v {
			sum += jsonSum(elem)
		}
		return sum
	default:
		return 0
	}
}
