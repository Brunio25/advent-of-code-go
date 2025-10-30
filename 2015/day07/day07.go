package day07

import (
	"advent-of-code-go/util/cast"
	"advent-of-code-go/util/stringy"
	"strconv"
	"strings"
)

func part1(input string) int {
	wireToGate := parseInput(input)
	wireResults := make(map[string]uint16) // not inline so I can consult the map for debugging purposes outside dfs func
	aNum := dfs("a", wireToGate, wireResults)
	return int(aNum)
}

func part2(input string) int {
	wireToGate := parseInput(input)
	wireResults := make(map[string]uint16) // not inline so I can consult the map for debugging purposes outside dfs func
	aNum := dfs("a", wireToGate, wireResults)

	wireToGate["b"] = cast.ToString(aNum)
	return int(dfs("a", wireToGate, map[string]uint16{}))
}

func dfs(wire string, wireToGate map[string]string, mem map[string]uint16) uint16 {
	if memVal, inMap := mem[wire]; inMap {
		return memVal
	}

	// If it's a number, the "wire" is actually a signal
	if n, ok := tryToNumber(wire); ok {
		return n
	}

	gateStr := wireToGate[wire]
	if n, ok := tryToNumber(gateStr); ok {
		mem[wire] = n
		return n
	}

	gateParts := strings.Split(gateStr, " ")

	var n uint16
	switch {
	case len(gateParts) == 1:
		n = dfs(gateParts[0], wireToGate, mem)
	case gateParts[0] == "NOT":
		n = ^dfs(gateParts[1], wireToGate, mem)
	case gateParts[1] == "AND":
		n = dfs(gateParts[0], wireToGate, mem) & dfs(gateParts[2], wireToGate, mem)
	case gateParts[1] == "OR":
		n = dfs(gateParts[0], wireToGate, mem) | dfs(gateParts[2], wireToGate, mem)
	case gateParts[1] == "LSHIFT":
		n = dfs(gateParts[0], wireToGate, mem) << dfs(gateParts[2], wireToGate, mem)
	case gateParts[1] == "RSHIFT":
		n = dfs(gateParts[0], wireToGate, mem) >> dfs(gateParts[2], wireToGate, mem)
	}

	mem[wire] = n
	return n
}

func tryToNumber(s string) (uint16, bool) {
	n, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, false
	}

	return uint16(n), true
}

func parseInput(input string) map[string]string {
	wireInstructions := make(map[string]string)

	for instruction := range stringy.SplitLines(input) {
		instParts := strings.Split(instruction, " -> ")
		wireInstructions[instParts[1]] = instParts[0]
	}

	return wireInstructions
}
