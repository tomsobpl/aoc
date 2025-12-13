package main

import (
	"regexp"
	"strconv"

	"github.com/tomsobpl/aoc/pkg/core"
	"github.com/tomsobpl/aoc/pkg/utils"
)

var idRangePattern = regexp.MustCompile(`^\d+-\d+$`)

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

func (s Solution) solvePartOne(payload string) core.AocResult {
	intRanges, ids := s.preparePayload(payload)
	freshIngredients := 0

	for _, id := range ids {
		for _, ir := range intRanges {
			if ir.Contains(id) {
				freshIngredients++
				break
			}
		}
	}

	return core.NewAocResult(strconv.Itoa(freshIngredients))
}

func (s Solution) solvePartTwo(payload string) core.AocResult {
	freshIngredients := 0

	intRanges, _ := s.preparePayload(payload)
	intRanges = utils.ReduceIntegerRanges(intRanges)

	for _, ir := range intRanges {
		freshIngredients += ir.IntegersCount()
	}

	return core.NewAocResult(strconv.Itoa(freshIngredients))
}

func (s Solution) preparePayload(rawPayload string) ([]utils.IntegerRange, []int) {
	intRanges := make([]utils.IntegerRange, 0)
	ids := make([]int, 0)

	for _, line := range utils.ConvertStringToLines(rawPayload) {
		if idRangePattern.MatchString(line) {
			parts := utils.ConvertStringToInts(line, "-")
			intRanges = append(intRanges, utils.IntegerRange{From: parts[0], To: parts[1]})
		} else if id, err := strconv.Atoi(line); err == nil {
			ids = append(ids, id)
		}
	}

	return intRanges, ids
}

func NewSolution() core.AocSolution {
	return Solution{}
}
