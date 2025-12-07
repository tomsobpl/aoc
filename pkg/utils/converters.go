package utils

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func ConvertStringToGridOfRunes(input string) ([][]rune, error) {
	rows := ConvertStringToNotEmptyLines(input)
	size := len(rows[0])
	grid := make([][]rune, len(rows))

	slices.Reverse(rows)

	for i, row := range rows {
		if len(row) != size {
			return nil, fmt.Errorf("invalid row length, expected %d, got %d", size, len(row))
		}

		grid[i] = []rune(row)
	}

	return grid, nil
}

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
