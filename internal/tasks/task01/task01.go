package task01

import (
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"regexp"
	"strconv"
	"strings"
)

var replaceMap = []string{
	"one", "o1e",
	"two", "t2o",
	"three", "t3e",
	"four", "f4r",
	"five", "f5e",
	"six", "s6x",
	"seven", "s7n",
	"eight", "e8t",
	"nine", "n9e",
}

func init() {
	tasks.Register(1, "A", SolveA)
	tasks.Register(1, "B", SolveB)
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

func SolveB(input string) (result string, err error) {
	lines, err := internal.ReadLines(input)
	if err != nil {
		return
	}

	replacer := strings.NewReplacer(replaceMap...)
	pattern := regexp.MustCompile("\\d")
	sum := 0

	for _, line := range lines {
		// Here we fix overlapping numbers (e.g. oneight). Because regex matches only the first one,
		// but for the task this is valid as "18".
		// So we replace it "oneight" (original) => "o1eight" (1st replace) => "o1e8t" (2nd replace)
		// to allow matching both 1 and 8.
		lineReplaced := replacer.Replace(line)
		lineReplaced = replacer.Replace(lineReplaced)
		matches := pattern.FindAllString(lineReplaced, -1)

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
