package game

import "fmt"

// A model represents a Sudoku puzzle that may be completely empty, partially or fully solved. It does not support CoW.
// Setting and unsetting positions updates the model in place.
type model struct {
	board [][]uint8
	width uint8
}

var _ solvable = &model{}

func newModel(size uint8) model {
	board := make([][]uint8, size)
	for i, _ := range board {
		row := make([]uint8, size)
		board[i] = row
	}
	return model{
		board: board,
		width: size,
	}
}

func (m *model) getWidth() uint8 {
	return m.width
}
func (m *model) getRow(i uint8) []uint8 {
	return m.board[i]
}

func (m *model) getColumn(i uint8) []uint8 {
	c := make([]uint8, m.width)
	for j, row := range m.board {
		c[j] = row[i]
	}
	return c
}

func (m *model) getGrid(x, y uint8) [][]uint8 {
	grid := make([][]uint8, 3)
	startY := 3 * y
	startX := 3 * x
	for j := startY; j < startY+3; j++ {
		grid[j-startY] = m.board[j][startX : startX+3]
	}
	return grid
}

func (m *model) getFirstOpenPosition() (position, error) {
	return m.getNextOpenPosition(position{0, 0, m.width})
}
func (m *model) getNextOpenPosition(p position) (position, error) {
	for {
		if m.board[p.y][p.x] == 0 {
			return p, nil
		}
		if err := p.next(); err != nil {
			return position{}, err
		}
	}
}

func (m *model) full() bool {
	if _, err := m.getFirstOpenPosition(); err != nil {
		return true
	}
	return false
}

// set updates the value of a position on the model in place and returns an error if it is unable to do so.
// An error is returned if the position is already set. An unset position is represented by a 0.
// TODO: Implement a CoW model. Consider the memento pattern.
func (m *model) set(p position, value uint8) error {
	if m.board[p.y][p.x] != 0 {
		return fmt.Errorf("position is already set")
	}
	m.board[p.y][p.x] = value
	return nil
}

func (m *model) unset(p position) {
	m.board[p.y][p.x] = 0
}

func (m *model) get(p position) uint8 {
	return m.board[p.y][p.x]
}

func (m *model) clone() cloneable {
	var clone = newModel(9)
	for i, row := range m.board {
		for j, value := range row {
			clone.board[i][j] = value
		}
	}
	return &clone
}
