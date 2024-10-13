package main

import (
	"fmt"
	"github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks"
	"os"
	"strconv"

	// Tasks self-register using the registry and init()
	_ "github.com/mindaugasw/advent-of-code-2023-golang/internal/tasks/task01"
)

func main() {
	taskFunc, ok := findTask()

	if !ok {
		return
	}

	fmt.Print("Solving... ")
	result := taskFunc()
	fmt.Printf("Completed:\n\n%v\n", result)
}

func findTask() (tasks.TaskFunction, bool) {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go 1 A")
		return nil, false
	}

	taskNumber, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Error: could not parse task number: " + os.Args[1])
		return nil, false
	}

	taskPart := os.Args[2]

	taskFunc, ok := tasks.Get(taskNumber, taskPart)

	if !ok {
		fmt.Printf("Error: could not find task %v %v\n", taskNumber, taskPart)
	}

	return taskFunc, true
}
