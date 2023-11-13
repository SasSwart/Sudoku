package game

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	var wanted uint8 = 1
	p := position{
		x: 0,
		y: 0,
	}
	m := newModel(9)
	err := m.set(p, wanted)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	got := m.get(p)
	if got != wanted {
		t.Errorf("result: %d, got %d", wanted, got)
		t.Fail()
	}
}

func newPartialModel() model {
	m := newModel(9)
	_ = m.set(position{0, 0, 9}, 1)
	return m
}
func newFullModel(size uint8) model {
	m := newModel(size)
	for i, _ := range m.board {
		for j, _ := range m.board[i] {
			_ = m.set(position{uint8(i), uint8(j), size}, 1)
		}
	}
	return m
}
func newMostlyFullModel(size uint8) model {
	m := newFullModel(size)
	m.unset(position{size - 1, size - 1, size})
	return m
}
func newFirstRowFullModel(size uint8) model {
	m := newModel(size)
	for i, _ := range m.board {
		_ = m.set(position{uint8(0), uint8(i), size}, 1)
	}
	return m
}

func TestGetFirstOpenPosition(t *testing.T) {
	tests := []struct {
		name string
		m    model
		p    position
		err  error
	}{
		{
			name: "empty board",
			m:    newModel(9),
			p:    position{},
		},
		{
			name: "first cell set",
			m:    newPartialModel(),
			p:    position{1, 0, 9},
		},
		{
			name: "first row set",
			m:    newFirstRowFullModel(2),
			p:    position{0, 1, 2},
		},
		{
			name: "only last cell open",
			m:    newMostlyFullModel(9),
			p:    position{8, 8, 9},
		},
		{
			name: "full board",
			m:    newFullModel(9),
			p:    position{0, 0, 9},
			err:  positionsExhausted{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.m.getFirstOpenPosition()
			if !errors.Is(err, test.err) {
				t.Errorf("Unexpected error. Wanted %v, got %v", test.err, err)
			}
			if !(got.x == test.p.x && got.y == test.p.y) {
				t.Errorf("Wanted %v, got %v", test.p, got)
			}
		})
	}
}

func TestGetGrid(t *testing.T) {
	tests := []struct {
		name   string
		m      model
		result [][]uint8
		grid   position
	}{
		{
			name: "first quadrant of empty board",
			m:    newModel(9),
			result: [][]uint8{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "first quadrant of board with first cell set",
			m:    newPartialModel(),
			result: [][]uint8{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "first quadrant of board with first row set",
			m:    newFirstRowFullModel(9),
			result: [][]uint8{
				{1, 0, 0},
				{1, 0, 0},
				{1, 0, 0},
			},
		},
		{
			name: "first quadrant of board with only last cell open",
			m:    newMostlyFullModel(9),
			result: [][]uint8{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
		},
		{
			name: "first quadrant of full board",
			m: model{
				board: fullBoard,
				width: 9,
			},
			result: [][]uint8{
				{7, 1, 4},
				{5, 2, 3},
				{6, 8, 9},
			},
		},
		{
			name: "second quadrant of full board",
			m: model{
				board: fullBoard,
				width: 9,
			},
			result: [][]uint8{
				{1, 3, 6},
				{2, 9, 7},
				{4, 5, 8},
			},
			grid: position{0, 1, 9},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(test.m)
			got := test.m.getGrid(test.grid.x, test.grid.y)

			if !reflect.DeepEqual(got, test.result) {
				t.Errorf("Wanted %v, got %v", test.result, got)
			}
		})
	}
}
