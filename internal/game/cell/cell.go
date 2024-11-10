package cell

var (
	Alive = &Cell{}
	Dead  *Cell
)

// Cell represents a cell in the Conway's Game of Life.
type Cell struct{}

// NextGeneration calculates the next generation of the cell based on the number of alive neighbors.
func (c *Cell) NextGeneration(aliveNeighbors int) *Cell {
	if c.canBeReproduced(aliveNeighbors) {
		return Alive
	}

	if c.isUnderpopulated(aliveNeighbors) {
		return Dead
	}

	if c.isOverpopulated(aliveNeighbors) {
		return Dead
	}

	return c
}

// canBeReproduced returns true if the dead cell has exactly three alive neighbors.
func (c *Cell) canBeReproduced(aliveNeighbors int) bool {
	return c == Dead && aliveNeighbors == 3
}

// isUnderpopulated returns true if an alive cell has fewer than two alive neighbors.
func (c *Cell) isUnderpopulated(aliveNeighbors int) bool {
	return aliveNeighbors < 2
}

// isOverpopulated returns true if an alive cell has more than three alive neighbors.
func (c *Cell) isOverpopulated(aliveNeighbors int) bool {
	return aliveNeighbors > 3
}
