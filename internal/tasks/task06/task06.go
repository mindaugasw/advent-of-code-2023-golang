package task06

import (
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"strings"
)

// https://adventofcode.com/2023/day/6

func init() {
	tasks.Register(6, "A", SolveA)
}

func SolveA(input string) (result int, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	times := parseLine(lines[0])
	distances := parseLine(lines[1])
	result = 1

	for i := 0; i < len(times); i++ {
		result *= calculateWaysToWin(times[i], distances[i])
	}

	return
}

func parseLine(line string) []int {
	line = strings.Split(line, ":")[1]
	numberStrList := strings.Split(line, " ")
	var numberList []int

	for _, numStr := range numberStrList {
		trimmed := strings.TrimSpace(numStr)

		if trimmed == "" {
			continue
		}

		numberList = append(numberList, internal.ParseInt(numStr))
	}

	return numberList
}

func calculateWaysToWin(time int, distance int) int {
	ways := 0

	for buttonHeldFor := 0; buttonHeldFor < time+1; buttonHeldFor++ {
		speed := buttonHeldFor
		travelTime := time - buttonHeldFor
		distanceTraveled := travelTime * speed

		if distanceTraveled > distance {
			ways++
		}
	}

	return ways
}
