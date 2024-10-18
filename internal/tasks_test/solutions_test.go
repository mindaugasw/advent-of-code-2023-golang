package tasks_test

import (
	"fmt"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"testing"

	// Tasks self-register using the registry and init()
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task01"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task02"
)

func TestSolutionProvider(t *testing.T) {
	dataSets := []struct {
		number   int
		part     string
		input    string
		expected string
	}{
		{1, "A", "exampleA.txt", "142"},
		{1, "A", "input.txt", "54927"},
		{1, "B", "exampleB.txt", "281"},
		{1, "B", "input.txt", "54581"},
		{2, "A", "example.txt", "8"},
		{2, "A", "input.txt", "2317"},
	}

	for _, set := range dataSets {
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
