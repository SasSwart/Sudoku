package game

type position struct {
	x, y uint8
	size uint8
}

type positionsExhausted struct {
}

func (p positionsExhausted) Error() string {
	return "positions exhausted"
}

func (p *position) next() error {
	if p.x < p.size-1 {
		p.x++
		return nil
	}
	if p.y >= p.size-1 {
		return positionsExhausted{}
	}
	p.x = 0
	p.y++
	return nil
}
