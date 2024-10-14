package task01

import (
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"regexp"
	"strconv"
)

func init() {
	tasks.Register(1, "A", SolveA)
}

func SolveA(input string) (result string, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	pattern := regexp.MustCompile("\\d")
	sum := 0

	for _, line := range lines {
		matches := pattern.FindAllString(line, -1)

		if len(matches) == 0 {
			continue
		}

		lineDigits := matches[0] + matches[len(matches)-1]
		lineNumber, err := strconv.Atoi(lineDigits)

		if err != nil {
			return "", err
		}

		sum += lineNumber
	}

	return strconv.Itoa(sum), nil
}
