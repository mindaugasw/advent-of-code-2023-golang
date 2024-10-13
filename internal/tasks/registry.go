package tasks

import (
	"strconv"
	"strings"
)

type TaskFunction func() string

var taskRegistry = make(map[string]TaskFunction)

func Register(number int, part string, taskFunc TaskFunction) {
	taskRegistry[fullName(number, part)] = taskFunc
}

func Get(number int, part string) (taskFunc TaskFunction, ok bool) {
	taskFunc, ok = taskRegistry[fullName(number, part)]

	return taskFunc, ok
}

func fullName(number int, part string) string {
	return strconv.Itoa(number) + strings.ToUpper(part)
}
