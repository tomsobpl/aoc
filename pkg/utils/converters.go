package utils

import (
	"strconv"
	"strings"
)

func ConvertStringToInt(input string) int {
	result, _ := strconv.Atoi(input)
	return result
}

func ConvertStringToLines(input string) []string {
	return strings.Split(input, "\n")
}
