package main

import (
	"strconv"

	"github.com/tomsobpl/aoc/pkg/core"
	"github.com/tomsobpl/aoc/pkg/utils"
	"github.com/tomsobpl/aoc/pkg/utils/math"
)

type area struct {
	a utils.Point
	b utils.Point
}

func (a area) c() utils.Point {
	return utils.Point{X: a.b.X, Y: a.a.Y}
}

func (a area) d() utils.Point {
	return utils.Point{X: a.a.X, Y: a.b.Y}
}

func (a area) bl() utils.Point {
	return utils.Point{X: min(a.a.X, a.b.X), Y: min(a.a.Y, a.b.Y)}
}

func (a area) br() utils.Point {
	return utils.Point{X: max(a.a.X, a.b.X), Y: min(a.a.Y, a.b.Y)}
}

func (a area) tl() utils.Point {
	return utils.Point{X: min(a.a.X, a.b.X), Y: max(a.a.Y, a.b.Y)}
}

func (a area) tr() utils.Point {
	return utils.Point{X: max(a.a.X, a.b.X), Y: max(a.a.Y, a.b.Y)}
}

func (a area) tilesCovered() int {
	x := max(a.a.X, a.b.X) - min(a.a.X, a.b.X) + 1
	y := max(a.a.Y, a.b.Y) - min(a.a.Y, a.b.Y) + 1

	return x * y
}

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
	areas := s.preparePayload(payload)
	value := 0

	for _, area := range areas {
		value = max(value, area.tilesCovered())
	}

	return core.NewAocResult(strconv.Itoa(value))
}

func (s Solution) solvePartTwo(payload string) core.AocResult {
	areas := s.preparePayload(payload)
	value := 0

	for _, area := range areas {
		value = max(value, area.tilesCovered())
	}

	return core.NewAocResult(strconv.Itoa(value))
}

func (s Solution) preparePayload(rawPayload string) []area {
	tiles := make([]utils.Point, 0)
	areas := make([]area, 0)

	for _, line := range utils.ConvertStringToNotEmptyLines(rawPayload) {
		tiles = append(tiles, utils.PointFromSlice(utils.ConvertStringToInts(line, ",")))
	}

	for _, tile := range math.CartesianProduct(tiles) {
		areas = append(areas, area{tile[0], tile[1]})
	}

	return areas
}

func NewSolution() core.AocSolution {
	return Solution{}
}
