package main

import (
	"errors"
	"fmt"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"os"
	// Tasks self-register using the registry and init()
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task01"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task02"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task03"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task04"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task05"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task06"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task07"
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task08"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	name, taskFunc, input, err := findTask()

	if err != nil {
		return
	}

	fmt.Printf("Solving %v ... ", name)
	result, err := taskFunc(input)

	if err != nil {
		return
	}

	fmt.Printf("Completed:\n%v\n", result)

	return nil
}

func findTask() (name string, taskFunc tasks.TaskFunction, input string, err error) {
	if len(os.Args) < 3 {
		err = errors.New("usage: go run main.go 1 A [input.txt]")
		return
	}

	taskNumber := internal.ParseInt(os.Args[1])
	taskPart := os.Args[2]
	taskFunc, ok := tasks.Get(taskNumber, taskPart)

	if !ok {
		err = errors.New(fmt.Sprintf("could not find task %v %v", taskNumber, taskPart))
		return
	}

	if len(os.Args) > 3 {
		input = os.Args[3]
	} else {
		input = "input.txt"
	}

	return tasks.GetFullName(taskNumber, taskPart), taskFunc, input, nil
}
