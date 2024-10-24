package task04

import (
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"strconv"
	"strings"
)

// https://adventofcode.com/2023/day/4

func init() {
	tasks.Register(4, "A", SolveA)
}

func SolveA(input string) (result int, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	sum := 0
	for _, line := range lines {
		cardWorth, err := calculateCardWorth(line)
		if err != nil {
			return 0, err
		}

		sum += cardWorth
	}

	return sum, nil
}

func calculateCardWorth(line string) (score int, err error) {
	line = strings.Split(line, ":")[1]
	winningStr, playingStr := func() (string, string) {
		split := strings.Split(line, "|")
		return split[0], split[1]
	}()

	winningList, err := parseNumberList(winningStr)
	if err != nil {
		return
	}

	playingList, err := parseNumberList(playingStr)
	if err != nil {
		return
	}

	for winningNum := range winningList {
		_, found := playingList[winningNum]

		if found {
			score *= 2
			score = max(score, 1)
		}
	}

	return
}

func parseNumberList(numbers string) (list map[int]bool, err error) {
	listStr := strings.Split(strings.TrimSpace(numbers), " ")
	list = make(map[int]bool)

	for _, numberStr := range listStr {
		if strings.TrimSpace(numberStr) == "" {
			continue
		}

		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return nil, err
		}

		list[number] = true
	}

	return
}
