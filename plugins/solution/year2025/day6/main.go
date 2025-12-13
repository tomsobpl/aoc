package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/tomsobpl/aoc/pkg/core"
	"github.com/tomsobpl/aoc/pkg/utils"
)

type Solution struct {
}

func (s Solution) Solve(data core.AocInput) core.AocResult {
	switch data.Part() {
	case core.AocTaskPartOne:
		return s.solvePartOne(data.Payload())
	case core.AocTaskPartTwo:
		return s.solvePartTwo(data.Payload())
	}

	return nil
}

func (s Solution) addProblems(data []string) int {
	result := 0

	for _, value := range utils.ConvertStringsToInts(data) {
		result += value
	}

	return result
}

func (s Solution) multiplyProblems(data []string) int {
	result := 1

	for _, value := range utils.ConvertStringsToInts(data) {
		result *= value
	}

	return result
}

func (s Solution) solvePartOne(payload string) core.AocResult {
	grandTotal := 0

	for _, problem := range utils.TransposeStringMatrix(s.preparePayload(payload)) {
		switch problem[len(problem)-1] {
		case "+":
			grandTotal += s.addProblems(problem[:len(problem)-1])
		case "*":
			grandTotal += s.multiplyProblems(problem[:len(problem)-1])
		}
	}

	return core.NewAocResult(strconv.Itoa(grandTotal))
}

func (s Solution) solvePartTwo(payload string) core.AocResult {
	grandTotal := 0
	subNumbers := make([]string, 0)

	problems := utils.TransposeRuneMatrix(s.preparePayloadOfRunes(payload))
	slices.Reverse(problems)

	for _, row := range problems {
		if "" == strings.TrimSpace(string(row)) {
			subNumbers = make([]string, 0)
			continue
		}

		subNumbers = append(subNumbers, string(row[:len(row)-1]))

		switch row[len(row)-1] {
		case '+':
			grandTotal += s.addProblems(subNumbers)
		case '*':
			grandTotal += s.multiplyProblems(subNumbers)
		}
	}

	return core.NewAocResult(strconv.Itoa(grandTotal))
}

func (s Solution) preparePayload(rawPayload string) [][]string {
	columns := make([][]string, 0)

	for _, line := range utils.ConvertStringToNotEmptyLines(rawPayload) {
		columns = append(columns, strings.Fields(line))
	}

	return columns
}

func (s Solution) preparePayloadOfRunes(rawPayload string) [][]rune {
	runes := make([][]rune, 0)

	for _, line := range utils.ConvertStringToNotEmptyLines(rawPayload) {
		runes = append(runes, []rune(line))
	}

	return runes
}

func NewSolution() core.AocSolution {
	return Solution{}
}
