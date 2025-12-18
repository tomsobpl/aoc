package main

import (
	"fmt"
	"strconv"

	"github.com/tomsobpl/aoc/pkg/core"
	"github.com/tomsobpl/aoc/pkg/utils"
)

const (
	emptyCell     rune = '.'
	splitter      rune = '^'
	startLocation rune = 'S'
	tachyonBeam   rune = '|'
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

func (s Solution) solvePartOne(payload string) core.AocResult {
	grid := s.preparePayload(payload)
	grid.FlipVertically()

	splits := 0

	for y := 1; y < grid.Height(); y++ {
		for x, v := range grid.Row(y) {
			switch v {
			case emptyCell:
				if grid.Cells[y-1][x] == startLocation || grid.Cells[y-1][x] == tachyonBeam {
					grid.Cells[y][x] = tachyonBeam
				}
			case splitter:
				if grid.Cells[y-1][x] == tachyonBeam {
					grid.Cells[y][x-1] = tachyonBeam
					grid.Cells[y][x+1] = tachyonBeam
					splits++
				}
			}
		}
	}

	return core.NewAocResult(strconv.Itoa(splits))
}

func (s Solution) solvePartTwo(payload string) core.AocResult {
	grid := s.preparePayload(payload)
	grid.FlipVertically()

	splitters := make(map[string]int)
	timelines := 0

	for x, v := range grid.Row(0) {
		if v == startLocation {
			timelines = s.followTimeline(grid, 1, x, timelines+1, splitters)
			break
		}
	}

	return core.NewAocResult(strconv.Itoa(timelines))
}

func (s Solution) followTimeline(grid *utils.GridOfRunes, y int, x int, timeline int, splitters map[string]int) int {
	k := fmt.Sprintf("%d,%d", y, x)

	if _, ok := (splitters)[k]; ok {
		return splitters[k]
	}

	if y >= grid.Height() {
		return timeline
	}

	switch grid.Cells[y][x] {
	case emptyCell:
		return s.followTimeline(grid, y+1, x, timeline, splitters)
	case splitter:
		splitters[k] = s.followTimeline(grid, y+1, x-1, timeline, splitters)
		splitters[k] += s.followTimeline(grid, y+1, x+1, timeline, splitters)
		return splitters[k]
	}

	return timeline
}

func (s Solution) preparePayload(rawPayload string) *utils.GridOfRunes {
	grid, err := utils.GridOfRunesFromString(rawPayload)

	if err != nil {
		fmt.Printf("Failed to parse grid of runes: %s\n", err)
		return nil
	}

	return grid
}

func NewSolution() core.AocSolution {
	return Solution{}
}
