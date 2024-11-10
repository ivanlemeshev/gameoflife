package grid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivanlemeshev/gameoflife/internal/game/cell"
	"github.com/ivanlemeshev/gameoflife/internal/game/grid"
)

func TestCellGrid_New(t *testing.T) {
	width := 3
	height := 3
	sg := grid.New(width, height)
	state := sg.State()
	for y := range height {
		for x := range width {
			assert.Equal(t, cell.Dead, state[y][x])
		}
	}
}

func TestCellGrid_ToggleCell(t *testing.T) {
	width := 3
	height := 3

	sg := grid.New(width, height)

	t.Run("toggle dead cells", func(t *testing.T) {
		expected := [][]*cell.Cell{
			{cell.Dead, cell.Dead, cell.Dead},
			{cell.Dead, cell.Dead, cell.Dead},
			{cell.Dead, cell.Dead, cell.Dead},
		}

		assert.Equal(t, expected, sg.State())

		for y := range height {
			for x := range width {
				expected[y][x] = cell.Alive
				sg.ToggleCell(x, y)
				assert.Equal(t, expected, sg.State())
			}
		}
	})

	t.Run("toggle alive cells", func(t *testing.T) {
		expected := [][]*cell.Cell{
			{cell.Alive, cell.Alive, cell.Alive},
			{cell.Alive, cell.Alive, cell.Alive},
			{cell.Alive, cell.Alive, cell.Alive},
		}

		assert.Equal(t, expected, sg.State())

		for y := range height {
			for x := range width {
				expected[y][x] = cell.Dead
				sg.ToggleCell(x, y)
				assert.Equal(t, expected, sg.State())
			}
		}
	})
}

func TestCellGrid_NextGeneration(t *testing.T) {
	tt := []struct {
		name           string
		addAliveCells  func(sg *grid.Grid)
		expectedBefore [][]*cell.Cell
		expectedAfter  [][]*cell.Cell
	}{
		{
			name:          "all dead cells should be dead",
			addAliveCells: func(sg *grid.Grid) {},
			expectedBefore: [][]*cell.Cell{
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
			expectedAfter: [][]*cell.Cell{
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
		},
		{
			name: "the dead cell should be dead if there are fewer than three alive neighbors",
			addAliveCells: func(sg *grid.Grid) {
				sg.ToggleCell(0, 0)
				sg.ToggleCell(0, 1)
			},
			expectedBefore: [][]*cell.Cell{
				{cell.Alive, cell.Dead, cell.Dead},
				{cell.Alive, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
			expectedAfter: [][]*cell.Cell{
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
		},
		{
			name: "the dead cell should be alive if there are three alive neighbors",
			addAliveCells: func(sg *grid.Grid) {
				sg.ToggleCell(0, 0)
				sg.ToggleCell(1, 0)
				sg.ToggleCell(2, 0)
			},
			expectedBefore: [][]*cell.Cell{
				{cell.Alive, cell.Alive, cell.Alive},
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
			expectedAfter: [][]*cell.Cell{
				{cell.Dead, cell.Alive, cell.Dead},
				{cell.Dead, cell.Alive, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
		},
		{
			name: "the dead cell should be dead if there are more than three alive neighbors",
			addAliveCells: func(sg *grid.Grid) {
				sg.ToggleCell(0, 0)
				sg.ToggleCell(1, 0)
				sg.ToggleCell(2, 0)
				sg.ToggleCell(0, 1)
			},
			expectedBefore: [][]*cell.Cell{
				{cell.Alive, cell.Alive, cell.Alive},
				{cell.Alive, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
			expectedAfter: [][]*cell.Cell{
				{cell.Alive, cell.Alive, cell.Dead},
				{cell.Alive, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
		},
		{
			name: "the alive cell should be dead if there are no alive neighbors",
			addAliveCells: func(sg *grid.Grid) {
				sg.ToggleCell(1, 1)
			},
			expectedBefore: [][]*cell.Cell{
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Alive, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
			expectedAfter: [][]*cell.Cell{
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
		},
		{
			name: "the alive cell should be dead if there are fewer than two alive neighbors",
			addAliveCells: func(sg *grid.Grid) {
				sg.ToggleCell(1, 0)
				sg.ToggleCell(1, 1)
			},
			expectedBefore: [][]*cell.Cell{
				{cell.Dead, cell.Alive, cell.Dead},
				{cell.Dead, cell.Alive, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
			expectedAfter: [][]*cell.Cell{
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
		},
		{
			name: "the alive cell should be alive if there are two alive neighbors",
			addAliveCells: func(sg *grid.Grid) {
				sg.ToggleCell(0, 0)
				sg.ToggleCell(1, 0)
				sg.ToggleCell(1, 1)
			},
			expectedBefore: [][]*cell.Cell{
				{cell.Alive, cell.Alive, cell.Dead},
				{cell.Dead, cell.Alive, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
			expectedAfter: [][]*cell.Cell{
				{cell.Alive, cell.Alive, cell.Dead},
				{cell.Alive, cell.Alive, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
		},
		{
			name: "the alive cell should be alive if there are thee alive neighbors",
			addAliveCells: func(sg *grid.Grid) {
				sg.ToggleCell(0, 0)
				sg.ToggleCell(1, 0)
				sg.ToggleCell(2, 0)
				sg.ToggleCell(1, 1)
			},
			expectedBefore: [][]*cell.Cell{
				{cell.Alive, cell.Alive, cell.Alive},
				{cell.Dead, cell.Alive, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
			expectedAfter: [][]*cell.Cell{
				{cell.Alive, cell.Alive, cell.Alive},
				{cell.Alive, cell.Alive, cell.Alive},
				{cell.Dead, cell.Dead, cell.Dead},
			},
		},
		{
			name: "the alive cell should be dead if there are more than thee alive neighbors",
			addAliveCells: func(sg *grid.Grid) {
				sg.ToggleCell(0, 0)
				sg.ToggleCell(1, 0)
				sg.ToggleCell(2, 0)
				sg.ToggleCell(0, 1)
				sg.ToggleCell(1, 1)
			},
			expectedBefore: [][]*cell.Cell{
				{cell.Alive, cell.Alive, cell.Alive},
				{cell.Alive, cell.Alive, cell.Dead},
				{cell.Dead, cell.Dead, cell.Dead},
			},
			expectedAfter: [][]*cell.Cell{
				{cell.Alive, cell.Dead, cell.Alive},
				{cell.Alive, cell.Dead, cell.Alive},
				{cell.Dead, cell.Dead, cell.Dead},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sg := grid.New(3, 3)

			tc.addAliveCells(sg)
			assert.Equal(t, tc.expectedBefore, sg.State())

			sg.NextGeneration()
			assert.Equal(t, tc.expectedAfter, sg.State())
		})
	}
}

func TestCellGrid_Generation(t *testing.T) {
	sg := grid.New(3, 3)
	assert.Equal(t, 0, sg.Generation())

	for i := 1; i < 10; i++ {
		sg.NextGeneration()
		assert.Equal(t, i, sg.Generation())
	}
}
