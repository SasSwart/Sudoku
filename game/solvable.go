package game

type solvable interface {
	cloneable
	getFirstOpenPosition() (position, error)
	getNextOpenPosition(p position) (position, error)
	set(p position, value uint8) error
	unset(p position)
	get(p position) uint8
	getRow(i uint8) []uint8
	getWidth() uint8
	getColumn(i uint8) []uint8
	full() bool
	getGrid(x, y uint8) [][]uint8
}

func solved(s solvable) error {
	if !s.full() {
		return boardNotFull{}
	}
	if err := checkInvalidRows(s); err != nil {
		return err
	}
	if err := checkInvalidColumns(s); err != nil {
		return err
	}
	if err := checkInvalidGrids(s); err != nil {
		return err
	}
	return nil
}

func sum(col []uint8) uint8 {
	var s uint8 = 0
	for _, u := range col {
		s += u
	}
	return s
}

func checkInvalidRows(s solvable) error {
	var i uint8
	for i = 0; i < s.getWidth(); i++ {
		sumOfRow := sum(s.getRow(i))
		if sumOfRow != (s.getWidth()*(s.getWidth()+1))/2 {
			return invalidRow{}
		}
	}
	return nil
}
func checkInvalidColumns(s solvable) error {
	var i uint8
	for i = 0; i < s.getWidth(); i++ {
		sumOfColumn := sum(s.getColumn(i))
		if sumOfColumn != (s.getWidth()*(s.getWidth()+1))/2 {
			return invalidRow{}
		}
	}
	return nil
}
func checkInvalidGrids(s solvable) error {
	var i, j uint8
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			grid := s.getGrid(i, j)
			missingNums := missingNumbers(flatten(grid), 9)
			if len(missingNums) > 0 {
				return invalidGrid{}
			}
		}
	}
	return nil
}

type boardNotFull struct{}

func (b boardNotFull) Error() string {
	return "Board Not Full"
}

type invalidRow struct{}

func (b invalidRow) Error() string {
	return "Invalid Row"
}

type invalidColumn struct{}

func (b invalidColumn) Error() string {
	return "Invalid Column"
}

type invalidGrid struct{}

func (b invalidGrid) Error() string {
	return "Invalid Grid"
}

type cloneable interface {
	clone() cloneable
}
