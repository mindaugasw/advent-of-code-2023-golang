package task04

import (
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"math"
	"strings"
)

// https://adventofcode.com/2023/day/4

func init() {
	tasks.Register(4, "A", SolveA)
	tasks.Register(4, "B", SolveB)
}

type card struct {
	winningNumCount int
	copies          int
}

func SolveA(input string) (result int, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	sum := 0
	for _, line := range lines {
		c, err := parseCard(line)
		if err != nil {
			return 0, err
		}

		cardWorth := int(math.Pow(2.0, float64(c.winningNumCount-1)))
		sum += cardWorth
	}

	return sum, nil
}

func SolveB(input string) (result int, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	var cards = make([]card, len(lines))

	for i, line := range lines {
		c, err := parseCard(line)
		if err != nil {
			return 0, err
		}

		cards[i].winningNumCount = c.winningNumCount
		cards[i].copies += c.copies

		for range cards[i].copies {
			for j := 1; j < c.winningNumCount+1; j++ {
				cards[i+j].copies++
			}
		}
	}

	sum := 0
	for _, c := range cards {
		sum += c.copies
	}

	return sum, nil
}

func parseCard(line string) (c card, err error) {
	c.copies = 1
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
			c.winningNumCount++
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

		number := internal.ParseInt(numberStr)
		list[number] = true
	}

	return
}
