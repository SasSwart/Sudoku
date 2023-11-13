package game

import (
	"errors"
	"fmt"
)

func flatten(slices [][]uint8) []uint8 {
	var flat []uint8
	for _, slice := range slices {
		flat = append(flat, slice...)
	}
	return flat
}

func commonNumbers(lists ...[]uint8) []uint8 {
	counts := make(map[uint8]int)
	for _, list := range lists {
		seen := make(map[uint8]bool)
		for _, num := range list {
			if !seen[num] {
				counts[num]++
				seen[num] = true
			}
		}
	}

	var common []uint8
	for num, cnt := range counts {
		if cnt == len(lists) {
			common = append(common, num)
		}
	}

	return common
}

func missingNumbers(input []uint8, n uint8) []uint8 {
	present := make(map[uint8]bool)
	for _, num := range input {
		present[num] = true
	}

	var missing []uint8
	var i uint8
	for i = 1; i <= n; i++ {
		if !present[i] {
			missing = append(missing, i)
		}
	}

	return missing
}

func bruteForceSynchronousSolve(m solvable) (solvable, error) {
	err := solved(m)
	if err == nil {
		return m, nil
	}

	if !errors.Is(err, boardNotFull{}) {
		return m, err
	}
	firstOpenPosition, err := m.getFirstOpenPosition()
	if err != nil {
		return m, fmt.Errorf("couldn't get first open position on a board that we know to not be full: %w", err)
	}

	row := m.getRow(firstOpenPosition.y)
	column := m.getColumn(firstOpenPosition.x)
	grid := flatten(m.getGrid(firstOpenPosition.x/3, firstOpenPosition.y/3))
	options := commonNumbers(
		missingNumbers(row[:], m.getWidth()),
		missingNumbers(column[:], m.getWidth()),
		missingNumbers(grid[:], m.getWidth()),
	)
	for _, val := range options {
		err = m.set(firstOpenPosition, val)

		if err != nil {
			return nil, fmt.Errorf("couldn't set first open position on a board that we know to not be full: %w", err)
		}

		m, err = bruteForceSynchronousSolve(m)
		if err == nil {
			return m, nil
		}
		m.unset(firstOpenPosition)
	}
	return m, fmt.Errorf("fuck")
}
