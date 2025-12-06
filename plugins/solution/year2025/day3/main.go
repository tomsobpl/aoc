package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/tomsobpl/aoc/pkg/core"
	"github.com/tomsobpl/aoc/pkg/utils"
)

var batteryBankPattern = regexp.MustCompile(`^\d{2,}$`)

type batteryBank struct {
	cells []int
}

type Solution struct {
}

func (b batteryBank) maxJoltage(startingCell int, activeCellsLimit int) int {
	if 0 == activeCellsLimit {
		return -1
	}

	cellIdx := startingCell
	lastIdx := len(b.cells) - activeCellsLimit + 1

	for i := startingCell; i < lastIdx; i++ {
		if b.cells[i] > b.cells[cellIdx] {
			cellIdx = i
		}
	}

	joltage := fmt.Sprintf("%d", b.cells[cellIdx])

	if tmp := b.maxJoltage(cellIdx+1, activeCellsLimit-1); tmp != -1 {
		joltage += fmt.Sprintf("%d", tmp)
	}

	return utils.ConvertStringToInt(joltage)
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
	result := s.findMaxJoltage(payload, 0, 2)
	return core.NewAocResult(strconv.Itoa(result))
}

func (s Solution) solvePartTwo(payload string) core.AocResult {
	result := s.findMaxJoltage(payload, 0, 12)
	return core.NewAocResult(strconv.Itoa(result))
}

func (s Solution) findMaxJoltage(payload string, startingCell int, activeCellsLimit int) int {
	result := 0

	for _, bank := range s.preparePayload(payload) {
		result += bank.maxJoltage(startingCell, activeCellsLimit)
	}

	return result
}

func (s Solution) preparePayload(rawPayload string) []batteryBank {
	lines := utils.ConvertStringToNotEmptyLines(rawPayload)
	banks := make([]batteryBank, len(lines))

	for i, bank := range lines {
		if bank_, err := newBatteryBankFromString(bank); err == nil {
			banks[i] = bank_
		}
	}

	return banks
}

func newBatteryBankFromString(raw string) (batteryBank, error) {
	if !batteryBankPattern.MatchString(raw) {
		return batteryBank{}, fmt.Errorf("invalid battery bank format: %s", raw)
	}

	return batteryBank{cells: utils.ConvertStringToInts(raw)}, nil
}

func NewSolution() core.AocSolution {
	return Solution{}
}
