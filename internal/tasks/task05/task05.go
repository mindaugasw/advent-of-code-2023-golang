package task05

import (
	"errors"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"strings"
)

// https://adventofcode.com/2023/day/5

func init() {
	tasks.Register(5, "A", SolveA)
	tasks.Register(5, "B", SolveB)
}

type mapRange struct {
	destination int64
	source      int64
	length      int64
}
type mapSection []mapRange
type almanac []mapSection

func SolveA(input string) (result int, err error) {
	return solve(input, false)
}

func SolveB(input string) (result int, err error) {
	return solve(input, true)
}

func solve(input string, treatAsRanges bool) (result int, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	seedSection, err := parseSeedRanges(lines[0], treatAsRanges)
	if err != nil {
		return
	}

	a := parseAlmanac(seedSection, lines[1:])
	result64, err := solveAlmanac(a)

	if err != nil {
		return
	}

	return int(result64), nil
}

func parseSeedRanges(line string, treatAsRanges bool) (section mapSection, err error) {
	line = strings.Split(line, ": ")[1]
	numberStrList := strings.Split(line, " ")
	increment := internal.If(treatAsRanges, 2, 1)

	for i := 0; i < len(numberStrList); i += increment {
		firstNum := internal.ParseInt64(numberStrList[i])

		if treatAsRanges {
			rangeLen := internal.ParseInt64(numberStrList[i+1])
			section = append(section, mapRange{firstNum, firstNum, rangeLen})
		} else {
			section = append(section, mapRange{firstNum, firstNum, 1})
		}
	}

	return
}

func parseAlmanac(seedSection mapSection, lines []string) (a almanac) {
	a = append(a, seedSection)
	i := 1

	for i < len(lines) {
		section := mapSection{}

		for ; i < len(lines); i++ {
			if strings.Contains(lines[i], "map:") {
				i++
				break
			}

			rangeParts := strings.Split(lines[i], " ")
			section = append(
				section, mapRange{
					destination: internal.ParseInt64(rangeParts[0]),
					source:      internal.ParseInt64(rangeParts[1]),
					length:      internal.ParseInt64(rangeParts[2]),
				},
			)
		}

		a = append(a, section)
	}

	return
}

func solveAlmanac(a almanac) (int64, error) {
	for resultGuess := int64(0); ; resultGuess++ {
		num := resultGuess
		found := false

		for i := len(a) - 1; i >= 0; i-- {
			num, found = a[i].translateDestinationToSource(num)

			if i == 0 && found {
				return resultGuess, nil
			}
		}

		if resultGuess == 325_000_000 {
			// Since this solution utilises an infinite loop, we stop it if it
			// missed the correct answer to prevent it from running forever
			return 0, errors.New("missed the correct answer")
		}
	}
}

func (ms mapSection) translateDestinationToSource(destination int64) (source int64, found bool) {
	for _, mr := range ms {
		if destination >= mr.destination && destination <= (mr.destination+mr.length) {
			destOffset := destination - mr.destination
			source = mr.source + destOffset
			return source, true
		}
	}

	return destination, false
}
