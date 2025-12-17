package utils

import (
	"fmt"
	"slices"
)

type GridOfRunes struct {
	Cells [][]rune
}

func (g *GridOfRunes) Dump() {
	for _, row := range g.Cells {
		fmt.Println(string(row))
	}
}

func (g *GridOfRunes) FlipHorizontally() {
	for _, row := range g.Cells {
		slices.Reverse(row)
	}
}

func (g *GridOfRunes) FlipVertically() {
	slices.Reverse(g.Cells)
}

func (g *GridOfRunes) GetAdjacentValuesTo(x int, y int) []rune {
	values := make([]rune, 0)

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			if x+dx >= 0 && x+dx < len(g.Cells[y]) && y+dy >= 0 && y+dy < len(g.Cells) {
				values = append(values, g.Cells[y+dy][x+dx])
			}
		}
	}

	return values
}

func (g *GridOfRunes) Height() int {
	return len(g.Cells)
}

func (g *GridOfRunes) Row(rowIndex int) []rune {
	return g.Cells[rowIndex]
}

func GridOfRunesFromString(input string) (*GridOfRunes, error) {
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

	return &GridOfRunes{Cells: grid}, nil
}
