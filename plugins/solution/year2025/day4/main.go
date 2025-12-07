package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tomsobpl/aoc/pkg/core"
	"github.com/tomsobpl/aoc/pkg/utils"
)

const paperRoll rune = '@'

type grid struct {
	cells [][]rune
}

type Solution struct {
}

func (g grid) getAdjacentValuesTo(x int, y int) []rune {
	values := make([]rune, 0)

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			if x+dx >= 0 && x+dx < len(g.cells[y]) && y+dy >= 0 && y+dy < len(g.cells) {
				values = append(values, g.cells[y+dy][x+dx])
			}
		}
	}

	return values
}

func (g grid) rollIsAccessible(x int, y int) bool {
	adjacentRolls := 0

	for _, value := range g.getAdjacentValuesTo(x, y) {
		if value == paperRoll {
			adjacentRolls++
		}
	}

	return adjacentRolls < 4
}

func (g grid) dump() {
	for _, row := range g.cells {
		fmt.Println(string(row))
	}
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

func (s Solution) findAccessiblePaperRolls(grid *grid, removeRolls bool) int {
	rollsWithAccess := 0
	rollsForRemoval := make([]string, 0)

	for y, row := range grid.cells {
		for x, col := range row {
			if col == paperRoll && grid.rollIsAccessible(x, y) {
				rollsWithAccess++
				rollsForRemoval = append(rollsForRemoval, fmt.Sprintf("%d,%d", x, y))
			}
		}
	}

	if removeRolls {
		for _, roll := range rollsForRemoval {
			loc := strings.Split(roll, ",")
			grid.cells[utils.ConvertStringToInt(loc[1])][utils.ConvertStringToInt(loc[0])] = 'x'
		}
	}

	return rollsWithAccess
}

func (s Solution) solvePartOne(payload string) core.AocResult {
	grid := s.preparePayload(payload)
	return core.NewAocResult(strconv.Itoa(s.findAccessiblePaperRolls(&grid, false)))
}

func (s Solution) solvePartTwo(payload string) core.AocResult {
	grid := s.preparePayload(payload)
	removedRollsSum := 0

	for {
		rollsWithAccess := s.findAccessiblePaperRolls(&grid, true)
		removedRollsSum += rollsWithAccess

		if rollsWithAccess == 0 {
			break
		}
	}

	return core.NewAocResult(strconv.Itoa(removedRollsSum))
}

func (s Solution) preparePayload(rawPayload string) grid {
	cells, _ := utils.ConvertStringToGridOfRunes(rawPayload)
	return grid{cells: cells}
}
func NewSolution() core.AocSolution {
	return Solution{}
}
