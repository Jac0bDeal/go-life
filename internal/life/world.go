package life

import (
	"math/rand"
	"time"
)

// World represents a Game of Life world.
type World struct {
	cells [][]int

	width  int
	height int
}

// newWorld generates a new, empty World and returns a pointer to it.
func newWorld(width, height int) *World {
	cells := make([][]int, height)
	for row := range cells {
		cells[row] = make([]int, width)
	}

	return &World{
		cells: cells,

		width:  width,
		height: height,
	}
}

// NewRandomWorld generates a new, randomized World of the specified width and height and returns a pointer to it.
// The percent of cells that should initially be alive should be passed as a value on the range [0.00, 1.00].
func NewRandomWorld(width int, height int, percentAlive float64) *World {
	rand.Seed(time.Now().UnixNano())

	world := newWorld(width, height)

	numAlive := int(float64(width*height) * percentAlive)
	for n := 0; n < numAlive; n++ {
		randIdx := rand.Intn(width * height)
		randX := randIdx % width
		randY := randIdx / width

		world.Set(randX, randY, 1)
	}

	return world
}

// CountAliveNeighbors returns the number of neighbors that are alive for a given cell.
func (w *World) CountAliveNeighbors(x, y int) int {
	var alive int

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if w.Get(x+i, y+j) == 1 {
				alive++
			}
		}
	}

	return alive
}

// Get gets the value of a cell at position (x, y), applying periodic boundary conditions.
func (w *World) Get(x, y int) int {
	x = (x + w.width) % w.width
	y = (y + w.height) % w.height
	return w.cells[x][y]
}

// Set sets the value of a cell at position (x, y) to the passed value, applying periodic boundary conditions.
func (w *World) Set(x, y, value int) {
	x = (x + w.width) % w.width
	y = (y + w.height) % w.height
	w.cells[x][y] = value
}

// State returns the current state of the world cells as a matrix of values.
func (w *World) State() [][]int {
	return w.cells
}

// Update performs an update to the world based on the set Kernel
func (w *World) Update() {
	newCells := make([][]int, w.height)
	for row := range newCells {
		newCells[row] = make([]int, w.width)
	}

	for i, row := range w.cells {
		for j, value := range row {
			aliveNeighbors := w.CountAliveNeighbors(i, j)
			if aliveNeighbors == 3 || (aliveNeighbors == 2 && value == 1) {
				newCells[i][j] = 1
			}
		}
	}

	w.cells = newCells
}
