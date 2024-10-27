package task08

import (
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"strings"
)

// https://adventofcode.com/2023/day/8

func init() {
	tasks.Register(8, "A", SolveA)
	tasks.Register(8, "B", SolveB)
}

type node struct {
	key      string
	children [2]string
	isStart  bool
	isEnd    bool
}

func SolveA(input string) (result int, err error) {
	return solve(input, false)
}

func SolveB(input string) (result int, err error) {
	return solve(input, true)
}

func solve(input string, simultaneousTraversal bool) (result int, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	instructions := lines[0]
	nodes := parseNodes(lines)

	if !simultaneousTraversal {
		result = traverseNodesSimple(instructions, nodes)
	} else {
		result = traverseNodesSimultaneous(instructions, nodes)
	}

	return
}

func parseNodes(lines []string) map[string]node {
	nodes := make(map[string]node)

	for _, line := range lines[1:] {
		split := strings.Split(line, " = (")
		nodeKey := split[0]

		split = strings.Split(split[1], ", ")
		leftNode := split[0]

		split = strings.Split(split[1], ")")
		rightNode := split[0]

		nodes[nodeKey] = node{
			key:      nodeKey,
			children: [2]string{leftNode, rightNode},
			isStart:  nodeKey[2] == 'A',
			isEnd:    nodeKey[2] == 'Z',
		}
	}

	return nodes
}

func traverseNodesSimple(instructions string, nodes map[string]node) int {
	steps := 0
	currentNode := nodes["AAA"]

	for {
		direction := instructions[steps%len(instructions)]
		directionIndex := internal.If(direction == 'L', 0, 1)
		nextKey := currentNode.children[directionIndex]
		steps++

		if nextKey == "ZZZ" {
			return steps
		}

		currentNode = nodes[nextKey]
	}
}

func traverseNodesSimultaneous(instructions string, nodes map[string]node) int {
	steps := 0
	var currentNodes []node

	for _, n := range nodes {
		if n.isStart {
			currentNodes = append(currentNodes, n)
		}
	}

	stepsToEndNode := make([]int, len(currentNodes))

	for {
		direction := instructions[steps%len(instructions)]
		directionIndex := internal.If(direction == 'L', 0, 1)
		steps++

		for i := 0; i < len(currentNodes); i++ {
			nextKey := currentNodes[i].children[directionIndex]
			currentNodes[i] = nodes[nextKey]

			if currentNodes[i].isEnd && stepsToEndNode[i] == 0 {
				stepsToEndNode[i] = steps
			}
		}

		allPathsReachedEnd := true

		for _, s := range stepsToEndNode {
			if s == 0 {
				allPathsReachedEnd = false
				break
			}
		}

		if allPathsReachedEnd {
			break
		}
	}

	return internal.LeastCommonMultiple(stepsToEndNode[0], stepsToEndNode[1], stepsToEndNode[2:]...)
}
