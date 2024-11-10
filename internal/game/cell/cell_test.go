package cell_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivanlemeshev/gameoflife/internal/game/cell"
)

func TestCell_NextGeneration(t *testing.T) {
	tt := []struct {
		name           string
		aliveNeighbors int
		cell           *cell.Cell
		expected       *cell.Cell
	}{
		{
			name:           "the dead cell should be dead if there are fewer than three alive neighbors",
			aliveNeighbors: 2,
			cell:           cell.Dead,
			expected:       cell.Dead,
		},
		{
			name:           "the dead cell should be alive if there are three alive neighbors",
			aliveNeighbors: 3,
			cell:           cell.Dead,
			expected:       cell.Alive,
		},
		{
			name:           "the dead should be dead if there are more than three alive neighbors",
			aliveNeighbors: 4,
			cell:           cell.Dead,
			expected:       cell.Dead,
		},
		{
			name:           "the alive cell should be dead if there are no alive neighbors",
			aliveNeighbors: 0,
			cell:           cell.Alive,
			expected:       cell.Dead,
		},
		{
			name:           "the alive cell should be dead if there are fewer than two alive neighbors",
			aliveNeighbors: 1,
			cell:           cell.Alive,
			expected:       cell.Dead,
		},
		{
			name:           "the alive should be alive if there are two alive neighbors",
			aliveNeighbors: 2,
			cell:           cell.Alive,
			expected:       cell.Alive,
		},
		{
			name:           "the alive should be alive if there are thee alive neighbors",
			aliveNeighbors: 3,
			cell:           cell.Alive,
			expected:       cell.Alive,
		},
		{
			name:           "the alive should be dead if there are more than thee alive neighbors",
			aliveNeighbors: 4,
			cell:           cell.Alive,
			expected:       cell.Dead,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.cell.NextGeneration(tc.aliveNeighbors))
		})
	}
}
