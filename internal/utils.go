package internal

import (
	"fmt"
	"strconv"
)

// If simulates ternary operator functionality: returns true or false value based on condition.
//
// From https://stackoverflow.com/a/59375088/4110469
func If[T any](condition bool, valTrue, valFalse T) T {
	if condition {
		return valTrue
	}

	return valFalse
}

func ParseInt(s string) int {
	return int(ParseInt64(s))
}

func ParseInt64(s string) int64 {
	result, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(fmt.Sprintf("Could not convert string \"%v\" to int: %v", s, err))
	}

	return result
}
