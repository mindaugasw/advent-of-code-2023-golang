package task03

import (
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"regexp"
	"strconv"
)

// https://adventofcode.com/2023/day/3

func init() {
	tasks.Register(3, "A", SolveA)
}

func SolveA(input string) (result string, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	pattern := regexp.MustCompile("\\d+")
	sum := 0

	for i := 0; i < len(lines); i++ {
		findResult := pattern.FindAllStringIndex(lines[i], -1)

		for _, matchedNumber := range findResult {
			matchedNumberIdxStart, matchedNumberIdxEnd := matchedNumber[0], matchedNumber[1]

			if isPartNumber(lines, i, matchedNumberIdxStart, matchedNumberIdxEnd-matchedNumberIdxStart) {
				partNumber, err := strconv.Atoi(lines[i][matchedNumberIdxStart:matchedNumberIdxEnd])
				if err != nil {
					return "", err
				}

				sum += partNumber
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func isPartNumber(schematic []string, line int, index int, length int) bool {
	lineLength := len(schematic[0])
	// surroundings is a block around the number.
	// idxStart/End is the index (inclusive) of edge symbols that will go into the surroundings block.
	// E.g.:
	//          .....
	//          .455.
	//          ...*.
	// idxStart ^   ^ idxEnd
	var surroundings []string
	idxStart := max(0, index-1)
	idxEnd := min(lineLength-1, index+length)

	// add line above
	if line != 0 {
		surroundings = append(surroundings, schematic[line-1][idxStart:idxEnd+1])
	}

	// add left symbol
	if index != 0 {
		surroundings = append(surroundings, schematic[line][idxStart:idxStart+1])
	}

	// add right symbol
	if idxEnd != (lineLength - 1) {
		surroundings = append(surroundings, schematic[line][idxEnd:idxEnd+1])
	}

	// add line below
	if line != (len(schematic) - 1) {
		surroundings = append(surroundings, schematic[line+1][idxStart:idxEnd+1])
	}

	for _, symbols := range surroundings {
		for _, symbol := range symbols {
			if symbol != '.' {
				return true
			}
		}
	}

	return false
}
