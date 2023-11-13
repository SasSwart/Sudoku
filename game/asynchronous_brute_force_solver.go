package game

import (
	"errors"
	"fmt"
)

func bruteForceAsynchronousSolve(m solvable) (solvable, error) {
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

	modelChan := make(chan solvable)
	errorChan := make(chan error)
	for i, _ := range options {
		i := i
		newM := m.clone().(solvable)
		go func() {
			err = newM.set(firstOpenPosition, options[i])
			if err != nil {
				errorChan <- fmt.Errorf("couldn't set first open position on a board that we know to not be full: %w", err)
			}
			newM, err = bruteForceSynchronousSolve(newM)
			if err == nil {
				modelChan <- newM
			}
		}()
	}
	for range options {
		select {
		case s := <-modelChan:
			return s, nil
		case err := <-errorChan:
			fmt.Println(err)
		}
	}
	return m, fmt.Errorf("fuck")
}
