package tasks

import (
	"strconv"
	"strings"
)

type TaskFunction func(string) (string, error)

var taskRegistry = make(map[string]TaskFunction)

func Register(number int, part string, taskFunc TaskFunction) {
	taskRegistry[GetFullName(number, part)] = taskFunc
}

func Get(number int, part string) (taskFunc TaskFunction, ok bool) {
	taskFunc, ok = taskRegistry[GetFullName(number, part)]

	return taskFunc, ok
}

func GetFullName(number int, part string) string {
	return strconv.Itoa(number) + strings.ToUpper(part)
}