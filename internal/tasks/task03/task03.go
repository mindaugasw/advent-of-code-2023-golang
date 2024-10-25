package task03

import (
	"fmt"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"regexp"
)

// https://adventofcode.com/2023/day/3

func init() {
	tasks.Register(3, "A", SolveA)
	tasks.Register(3, "B", SolveB)
}

var gearMap map[string][]int

func SolveA(input string) (result int, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	gearMap = make(map[string][]int)
	pattern := regexp.MustCompile("\\d+")
	sum := 0

	for i := 0; i < len(lines); i++ {
		findResult := pattern.FindAllStringIndex(lines[i], -1)

		for _, matchedNumber := range findResult {
			matchedNumberIdxStart, matchedNumberIdxEnd := matchedNumber[0], matchedNumber[1]

			isPart, err := isPartNumber(lines, i, matchedNumberIdxStart, matchedNumberIdxEnd-matchedNumberIdxStart)
			if err != nil {
				return 0, err
			}

			if isPart {
				partNumber := internal.ParseInt(lines[i][matchedNumberIdxStart:matchedNumberIdxEnd])
				sum += partNumber
			}
		}
	}

	return sum, nil
}

func SolveB(input string) (result int, err error) {
	// SolveA is needed to fill gears map
	_, err = SolveA(input)
	if err != nil {
		return
	}

	sum := 0
	for _, values := range gearMap {
		if len(values) != 2 {
			continue
		}

		sum += values[0] * values[1]
	}

	return sum, nil
}

func isPartNumber(schematic []string, line int, index int, length int) (isPart bool, err error) {
	numberStr := schematic[line][index : index+length]
	number := internal.ParseInt(numberStr)
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

	if line != 0 {
		lineAbove := schematic[line-1][idxStart : idxEnd+1]
		surroundings = append(surroundings, lineAbove)
		logGears(number, lineAbove, line-1, idxStart)
	}

	if index != 0 {
		leftSymbol := schematic[line][idxStart : idxStart+1]
		surroundings = append(surroundings, leftSymbol)
		logGears(number, leftSymbol, line, idxStart)
	}

	if idxEnd != (lineLength - 1) {
		rightSymbol := schematic[line][idxEnd : idxEnd+1]
		surroundings = append(surroundings, rightSymbol)
		logGears(number, rightSymbol, line, idxEnd)
	}

	if line != (len(schematic) - 1) {
		lineBelow := schematic[line+1][idxStart : idxEnd+1]
		surroundings = append(surroundings, lineBelow)
		logGears(number, lineBelow, line+1, idxStart)
	}

	for _, symbols := range surroundings {
		for _, symbol := range symbols {
			if symbol != '.' {
				return true, nil
			}
		}
	}

	return false, nil
}

func logGears(number int, surroundText string, line int, blockStartIdx int) {
	for i, symbol := range surroundText {
		if symbol != '*' {
			continue
		}

		positionText := fmt.Sprintf("%v;%v", line, blockStartIdx+i)
		gearMap[positionText] = append(gearMap[positionText], number)
	}
}
