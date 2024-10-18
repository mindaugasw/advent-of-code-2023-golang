package task02

import (
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"strconv"
	"strings"
)

func init() {
	tasks.Register(2, "A", SolveA)
}

type cubeSet map[string]int

type game struct {
	id   int
	sets []cubeSet
}

func newCubeSet() cubeSet {
	return map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
}

func SolveA(input string) (result string, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	condition := cubeSet{"red": 12, "green": 13, "blue": 14}
	sum := 0

	for _, line := range lines {
		g, err := parseLine(line)

		if err != nil {
			return "", err
		}

		if wasGamePossible(g, condition) {
			sum += g.id
		}
	}

	return strconv.Itoa(sum), nil
}

func parseLine(line string) (g game, err error) {
	gameStr, setsStr := func() (string, string) {
		split := strings.Split(line, ":")
		return split[0], split[1]
	}()

	gameId := strings.Split(gameStr, " ")[1]
	g.id, err = strconv.Atoi(gameId)

	if err != nil {
		return
	}

	setsSplit := strings.Split(setsStr, ";")

	for _, setStr := range setsSplit {
		colorsSplit := strings.Split(setStr, ",")
		cubes := newCubeSet()
		g.sets = append(g.sets, cubes)

		for _, colorSplit := range colorsSplit {
			colorParts := strings.Split(strings.TrimSpace(colorSplit), " ")
			colorCountStr, colorName := colorParts[0], colorParts[1]
			colorCount, err := strconv.Atoi(colorCountStr)

			if err != nil {
				return game{}, err
			}

			cubes[colorName] = colorCount
		}
	}

	return g, nil
}

func wasGamePossible(g game, condition cubeSet) bool {
	for _, set := range g.sets {
		for conditionColor, conditionCount := range condition {
			if set[conditionColor] > conditionCount {
				return false
			}
		}
	}

	return true
}
