package game

import (
	"errors"
	"testing"
)

var fullBoard = [][]uint8{
	{7, 1, 4, 2, 3, 6, 5, 8, 9},
	{5, 2, 3, 1, 8, 9, 4, 7, 6},
	{6, 8, 9, 4, 5, 7, 1, 2, 3},
	{1, 3, 6, 7, 2, 8, 9, 5, 4},
	{2, 9, 7, 3, 4, 5, 8, 6, 1},
	{4, 5, 8, 9, 6, 1, 7, 3, 2},
	{3, 4, 5, 8, 9, 2, 6, 1, 7},
	{8, 7, 2, 6, 1, 4, 3, 9, 5},
	{9, 6, 1, 5, 7, 3, 2, 4, 8},
}

func TestCheckInvalidGrids(t *testing.T) {
	m := newModel(9)
	m.board = fullBoard
	expectedErr := boardNotFull{}
	t.Run("", func(t *testing.T) {
		err := checkInvalidGrids(&m)
		if !errors.Is(err, expectedErr) {
			t.Errorf("Unexpected error. Wanted %v, got %v", expectedErr, err)
		}

	})
}
