package tasks_test

import (
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"testing"
)

func TestSolutionProvider(t *testing.T) {
	dataSets := []struct {
		number   int
		part     string
		input    string
		expected string
	}{
		{1, "A", "input.txt", "54927"},
		{1, "A", "example.txt", "142"},
	}

	for _, set := range dataSets {
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
	}
}
