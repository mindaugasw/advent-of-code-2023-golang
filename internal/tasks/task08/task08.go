package task08

import (
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"strings"
)

// https://adventofcode.com/2023/day/8

func init() {
	tasks.Register(8, "A", SolveA)
}

type node [2]string

func SolveA(input string) (result int, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	instructions := lines[0]
	nodes := parseNodes(lines)
	result = traverseNodes(instructions, nodes)

	return
}

func parseNodes(lines []string) map[string]node {
	nodes := make(map[string]node)

	for _, line := range lines[1:] {
		split := strings.Split(line, " = (")
		homeNode := split[0]

		split = strings.Split(split[1], ", ")
		leftNode := split[0]

		split = strings.Split(split[1], ")")
		rightNode := split[0]

		nodes[homeNode] = node{leftNode, rightNode}
	}

	return nodes
}

func traverseNodes(instructions string, nodes map[string]node) int {
	steps := 0
	currentNode := nodes["AAA"]

	for {
		direction := instructions[steps%len(instructions)]
		directionIndex := internal.If(direction == 'L', 0, 1)
		nextKey := currentNode[directionIndex]
		steps++

		if nextKey == "ZZZ" {
			return steps
		}

		currentNode = nodes[nextKey]
	}
}
