package grid

import (
	"github.com/ivanlemeshev/gameoflife/internal/game/cell"
)

// Grid represents a grid of cells.
type Grid struct {
	generation int
	width      int
	height     int
	grid       [][]*cell.Cell
}

// New creates a new cell grid with the given width and height.
func New(width, height int) *Grid {
	return &Grid{
		width:  width,
		height: height,
		grid:   newEmptyGrid(width, height),
	}
}

// Generation returns the current generation of the cell grid.
func (g *Grid) Generation() int {
	return g.generation
}

// State returns the current state of the cell grid.
func (g *Grid) State() [][]*cell.Cell {
	return g.grid
}

// ToggleCell makes the cell alive or dead depending on the current state in the x-th column and y-th row.
func (g *Grid) ToggleCell(x, y int) {
	if g.grid[y][x] == cell.Dead {
		g.grid[y][x] = cell.Alive
		return
	}

	g.grid[y][x] = cell.Dead
}

// NextGeneration moves the cell grid to the next generation.
func (g *Grid) NextGeneration() {
	// We need to keep the state the same while we calculate the next generation.
	// That's why we need a new grid.
	nextGenerationGrid := newEmptyGrid(g.width, g.height)

	for y := range g.grid {
		for x := range g.grid[y] {
			aliveNeighbors := g.countAliveNeighbors(x, y)
			cell := g.grid[y][x]
			nextGenerationCell := cell.NextGeneration(aliveNeighbors)
			nextGenerationGrid[y][x] = nextGenerationCell
		}
	}

	g.grid = nextGenerationGrid
	g.generation++
}

// countAliveNeighbors counts the number of alive neighbors of a cell.
func (g *Grid) countAliveNeighbors(x, y int) int {
	aliveNeighbors := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			// Skip the cell itself.
			if dx == 0 && dy == 0 {
				continue
			}

			// Skip the cell if it's outside the grid.
			// We consider them as dead cells.
			ny := y + dy
			nx := x + dx

			if ny < 0 || ny >= len(g.grid) {
				continue
			}

			if nx < 0 || nx >= len(g.grid[ny]) {
				continue
			}

			// Count the cell as an alive neighbor.
			if g.grid[ny][nx] == cell.Alive {
				aliveNeighbors++
			}
		}
	}

	return aliveNeighbors
}

// newEmptyGrid creates a new empty grid of cells with the given width and height.
// All cells are dead in the beginning.
func newEmptyGrid(width, height int) [][]*cell.Cell {
	grid := make([][]*cell.Cell, height)
	for y := range height {
		grid[y] = make([]*cell.Cell, width)
		for x := range width {
			grid[y][x] = cell.Dead
		}
	}

	return grid
}
