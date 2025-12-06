package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/tomsobpl/aoc/pkg/core"
	"github.com/tomsobpl/aoc/pkg/utils"
)

var idRangePattern = regexp.MustCompile(`^\d+-\d+$`)

type idRange struct {
	from int
	to   int
}

type Solution struct {
}

func (r idRange) invalidDoubleRepetitionIds() []int {
	var result []int
	for i := r.from; i <= r.to; i++ {
		s := strconv.Itoa(i)

		if stringIsComposedFromDoubleRepetitionOfSequence(s) {
			result = append(result, i)
		}
	}

	return result
}

func (r idRange) invalidMultipleRepetitionIds() []int {
	var result []int
	for i := r.from; i <= r.to; i++ {
		s := strconv.Itoa(i)

		if stringIsComposedFromMultipleRepetitionsOfSequence(s) {
			result = append(result, i)
		}
	}

	return result
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
	var invalidIds []int

	for _, r := range s.preparePayload(payload) {
		invalidIds = append(invalidIds, r.invalidDoubleRepetitionIds()...)
	}

	sum := 0
	for _, id := range invalidIds {
		sum += id
	}

	return core.NewAocResult(strconv.Itoa(sum))
}

func (s Solution) solvePartTwo(payload string) core.AocResult {
	var invalidIds []int

	for _, r := range s.preparePayload(payload) {
		invalidIds = append(invalidIds, r.invalidMultipleRepetitionIds()...)
	}

	sum := 0
	for _, id := range invalidIds {
		sum += id
	}

	return core.NewAocResult(strconv.Itoa(sum))
}

func (s Solution) preparePayload(rawPayload string) []idRange {
	lines := utils.ConvertStringToLines(rawPayload)
	pairs := strings.Split(lines[0], ",")
	ranges := make([]idRange, len(pairs))

	for i, pair := range pairs {
		if range_, err := newIdRangeFromString(pair); err == nil {
			ranges[i] = range_
		}
	}

	return ranges
}

func newIdRangeFromString(raw string) (idRange, error) {
	if !idRangePattern.MatchString(raw) {
		return idRange{}, fmt.Errorf("invalid id range format: %s", raw)
	}

	parts := strings.Split(raw, "-")

	return idRange{from: utils.ConvertStringToInt(parts[0]), to: utils.ConvertStringToInt(parts[1])}, nil
}

func stringIsComposedFromDoubleRepetitionOfSequence(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	half := len(s) / 2
	return s[:half] == s[half:]
}
func stringIsComposedFromMultipleRepetitionsOfSequence(s string) bool {
	if len(s) < 2 {
		return false
	}

	for sequenceLength := 1; sequenceLength <= len(s)/2; sequenceLength++ {
		if len(s)%sequenceLength != 0 {
			continue
		}

		sequence := s[:sequenceLength]
		isRepeating := true

		for i := sequenceLength; i < len(s); i += sequenceLength {
			if s[i:i+sequenceLength] != sequence {
				isRepeating = false
				break
			}
		}

		if isRepeating {
			return true
		}
	}

	return false
}

func NewSolution() core.AocSolution {
	return Solution{}
}
