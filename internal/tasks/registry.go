package tasks

import (
	"strconv"
	"strings"
)

type TaskFunction func(string) (int, error)

var taskRegistry = make(map[string]TaskFunction)

func Register(number int, part string, taskFunc TaskFunction) {
	if part != "A" && part != "B" {
		panic("unknown task part: " + part)
	}

	fullName := GetFullName(number, part)

	if _, found := taskRegistry[fullName]; found {
		panic("attempted to register task twice: " + fullName)
	}

	taskRegistry[fullName] = taskFunc
}

func Get(number int, part string) (taskFunc TaskFunction, ok bool) {
	taskFunc, ok = taskRegistry[GetFullName(number, part)]

	return taskFunc, ok
}

func GetFullName(number int, part string) string {
	return strconv.Itoa(number) + strings.ToUpper(part)
}
