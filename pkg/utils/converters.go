package utils

import (
	"strconv"
	"strings"
)

func ConvertStringToInt(input string) int {
	result, _ := strconv.Atoi(input)
	return result
}

func ConvertStringToInts(input string) []int {
	result := make([]int, len(input))

	for i, s := range strings.Split(input, "") {
		result[i] = ConvertStringToInt(s)
	}

	return result
}

func ConvertStringToLines(input string) []string {
	return strings.Split(input, "\n")
}

func ConvertStringToNotEmptyLines(input string) []string {
	var result []string

	for _, line := range ConvertStringToLines(input) {
		if line != "" {
			result = append(result, line)
		}
	}

	return result
}
