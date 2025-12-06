package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/tomsobpl/aoc/pkg/core"
	"github.com/tomsobpl/aoc/pkg/utils"
)

const (
	rotationLeft  = 'L'
	rotationRight = 'R'
)

var rotationPattern = regexp.MustCompile(`^([LR])\d+$`)

type Solution struct {
}

type dial struct {
	size  int
	value int
}

type rotation struct {
	direction rune
	distance  int
}

func newDial(size int, value int) *dial {
	return &dial{size: size, value: value}
}

func (d *dial) rotate(rotation rotation) int {
	switch rotation.direction {
	case rotationLeft:
		d.value -= rotation.distance % d.size

		if d.value < 0 {
			d.value += d.size
		}
	case rotationRight:
		d.value = (d.value + rotation.distance) % d.size
	}

	return d.value
}

func (d *dial) rotateWithTicksCount(rotation rotation) int {
	ticks := rotation.distance / d.size
	short := rotation.copyWithoutCycles(d.size)

	distanceToZero := 0

	switch short.direction {
	case rotationLeft:
		distanceToZero = d.value
	case rotationRight:
		distanceToZero = d.size - d.value
	}

	if 0 != distanceToZero && short.distance > distanceToZero {
		ticks++
	}

	if 0 == d.rotate(short) {
		ticks++
	}

	return ticks
}

func (r *rotation) copyWithoutCycles(cycleSize int) rotation {
	return rotation{direction: r.direction, distance: r.distance % cycleSize}
}

func newRotationFromString(raw string) (rotation, error) {
	if !rotationPattern.MatchString(raw) {
		return rotation{}, fmt.Errorf("invalid rotation format: %s", raw)
	}

	return rotation{
		direction: rune(raw[0]),
		distance:  utils.ConvertStringToInt(raw[1:]),
	}, nil
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
	rotations := s.preparePayload(payload)
	dial := newDial(100, 50)
	zeros := 0

	for _, rotation := range rotations {
		if dial.rotate(rotation) == 0 {
			zeros++
		}
	}

	return core.NewAocResult(strconv.Itoa(zeros))
}

func (s Solution) solvePartTwo(payload string) core.AocResult {
	rotations := s.preparePayload(payload)
	dial := newDial(100, 50)
	ticks := 0

	for _, rotation := range rotations {
		ticks += dial.rotateWithTicksCount(rotation)
	}

	return core.NewAocResult(strconv.Itoa(ticks))
}

func (s Solution) preparePayload(rawPayload string) []rotation {
	lines := utils.ConvertStringToLines(rawPayload)
	rotations := make([]rotation, len(lines))

	for i, line := range lines {
		if rotation, err := newRotationFromString(line); err == nil {
			rotations[i] = rotation
		}
	}

	return rotations
}

func NewSolution() core.AocSolution {
	return Solution{}
}
