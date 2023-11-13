package game

import (
	"errors"
	"testing"
)

func TestNext(t *testing.T) {
	type testResult struct {
		err error
		p   position
	}
	subtests := []struct {
		name   string
		p      position
		result testResult
	}{
		{
			name: "zero position",
			p:    position{},
			result: testResult{
				err: nil,
				p: position{
					x:    1,
					y:    0,
					size: 9,
				},
			},
		},
		{
			name: "last position",
			p: position{
				x:    8,
				y:    8,
				size: 9,
			},
			result: testResult{
				err: positionsExhausted{},
				p: position{
					x: 8,
					y: 8,
				},
			},
		},
		{
			name: "first row full",
			p: position{
				x:    1,
				y:    0,
				size: 2,
			},
			result: testResult{
				err: nil,
				p: position{
					x:    0,
					y:    1,
					size: 2,
				},
			},
		},
	}
	for _, test := range subtests {
		t.Run(test.name, func(t *testing.T) {
			err := test.p.next()
			if !errors.Is(err, test.result.err) {
				t.Errorf("Unexpected error. Wanted %v, got %v", test.result.err, err)
			}
			if test.p.x != test.result.p.x || test.p.y != test.result.p.y {
				t.Error("Next position is incorrect", test.p)
			}
		})
	}
}
