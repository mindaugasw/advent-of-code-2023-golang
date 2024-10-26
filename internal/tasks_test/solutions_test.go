package tasks_test

import (
	"fmt"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"os"
	"slices"
	"strings"
	"testing"

	// Tasks self-register using the registry and init()
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task01"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task02"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task03"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task04"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task05"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task06"
)

type dataSet struct {
	number      int
	part        string
	input       string
	expected    int
	longRunning bool
}

func newSet(number int, part string, input string, expected int) dataSet {
	if !strings.Contains(input, ".") {
		input += ".txt"
	}

	return dataSet{number, part, input, expected, false}
}

func newSetLongRunning(number int, part string, input string, expected int) dataSet {
	ds := newSet(number, part, input, expected)
	ds.longRunning = true
	return ds
}

func TestSolutionProvider(t *testing.T) {
	runLongTests := slices.Contains(os.Args, "long")

	dataSets := []dataSet{
		newSet(1, "A", "exampleA", 142),
		newSet(1, "A", "input", 54927),
		newSet(1, "B", "exampleB", 281),
		newSet(1, "B", "input", 54581),
		newSet(2, "A", "example", 8),
		newSet(2, "A", "input", 2317),
		newSet(2, "B", "example", 2286),
		newSet(2, "B", "input", 74804),
		newSet(3, "A", "example", 4361),
		newSet(3, "A", "input", 530849),
		newSet(3, "B", "example", 467835),
		newSet(3, "B", "input", 84900879),
		newSet(4, "A", "example", 13),
		newSet(4, "A", "input", 21568),
		newSet(4, "B", "example", 30),
		newSet(4, "B", "input", 11827296),
		newSet(5, "A", "example", 35),
		newSetLongRunning(5, "A", "input", 322500873),
		newSet(5, "B", "example", 46),
		newSetLongRunning(5, "B", "input", 108956227),
		newSet(6, "A", "example", 288),
		newSet(6, "A", "input", 588588),
		newSet(6, "B", "example", 71503),
		newSet(6, "B", "input", 34655848),
	}

	for _, set := range dataSets {
		if set.longRunning && !runLongTests {
			continue
		}

		testName := fmt.Sprintf("%d%s-%s", set.number, set.part, set.input)
		t.Run(
			testName, func(t *testing.T) {
				taskFunc, ok := tasks.Get(set.number, set.part)

				if !ok {
					t.Fatalf("Could not get taskFunc")
				}

				result, err := taskFunc(set.input)

				if err != nil {
					t.Fatalf("Error returned: %v", err)
				}

				if result != set.expected {
					t.Fatalf("Incorrect result. Expected '%v', got '%v'", set.expected, result)
				}
			},
		)
	}
}
