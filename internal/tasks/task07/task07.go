package task07

import (
	"cmp"
	"fmt"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"maps"
	"slices"
	"strings"
)

// https://adventofcode.com/2023/day/7

func init() {
	tasks.Register(7, "A", SolveA)
}

type hand struct {
	cards       string
	cNumbers    []int
	occurrences []int
	bid         int
	strength    int64
}

func SolveA(input string) (result int, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	hands, err := parseInput(lines)
	if err != nil {
		return
	}

	slices.SortFunc(
		hands, func(a hand, b hand) int {
			return cmp.Compare(a.strength, b.strength)
		},
	)

	for i := 0; i < len(hands); i++ {
		cardWinnings := hands[i].bid * (i + 1)
		result += cardWinnings
	}

	return
}

func parseInput(lines []string) (hands []hand, err error) {
	for _, line := range lines {
		h, err := parseHand(line)
		if err != nil {
			return nil, err
		}

		hands = append(hands, h)
	}

	return
}

func parseHand(line string) (h hand, err error) {
	split := strings.Split(line, " ")

	h.cards = split[0]
	h.bid = internal.ParseInt(split[1])
	h.cNumbers = cardsToNumbers(h.cards)
	h.occurrences = countOccurrences(h.cNumbers)
	h.strength = calculateHandStrength(h)

	return
}

func cardsToNumbers(cards string) []int {
	var result = make([]int, 5)

	for i, card := range cards {
		var cardNumber int
		switch card {
		case 'A':
			cardNumber = 14
		case 'K':
			cardNumber = 13
		case 'Q':
			cardNumber = 12
		case 'J':
			cardNumber = 11
		case 'T':
			cardNumber = 10
		default:
			cardNumber = internal.ParseInt(string(card))
		}

		result[i] = cardNumber
	}

	return result
}

func countOccurrences(cNumbers []int) []int {
	occurrences := make(map[int]int)

	for _, num := range cNumbers {
		occurrences[num]++
	}

	occurrenceValues := slices.Collect(maps.Values(occurrences))
	slices.SortFunc(occurrenceValues, func(a int, b int) int { return cmp.Compare(b, a) })

	return occurrenceValues
}

// calculateHandStrength returns a single number identifying hand strength.
// First digit is hand type (e.g. five of a kind, full house, etc.).
// Subsequent digits are all card values in a row, to allow comparing same type hands.
func calculateHandStrength(h hand) int64 {
	var handType int

	switch {
	case h.occurrences[0] == 5: // Five of a kind
		handType = 7
	case h.occurrences[0] == 4: // Four of a kind
		handType = 6
	case h.occurrences[0] == 3 && h.occurrences[1] == 2: // Full house
		handType = 5
	case h.occurrences[0] == 3: // Three of a kind
		handType = 4
	case h.occurrences[0] == 2 && h.occurrences[1] == 2: // Two pair
		handType = 3
	case h.occurrences[0] == 2: // One pair
		handType = 2
	default: // High card
		handType = 1
	}

	strengthStr := fmt.Sprintf(
		"%d%02d%02d%02d%02d%02d",
		handType,
		h.cNumbers[0],
		h.cNumbers[1],
		h.cNumbers[2],
		h.cNumbers[3],
		h.cNumbers[4],
	)

	return internal.ParseInt64(strengthStr)
}
