package game

import (
	"fmt"
	"os"
	"reflect"
	"runtime/pprof"
	"testing"
)

func BenchmarkBruteForceSynchronousSolve(b *testing.B) {
	f, err := os.Create("sync_cpu.prof")
	if err != nil {
		b.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		b.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < b.N; i++ {
		m := newModel(9)
		_, err := bruteForceSynchronousSolve(&m)
		if err != nil {
			b.Error(err)
		}
	}
}

func TestBruteForceSynchronousSolve(t *testing.T) {
	m := newModel(9)
	solved, err := bruteForceSynchronousSolve(&m)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(solved)
}

func TestMissingNumbers(t *testing.T) {
	tests := []struct {
		name   string
		size   uint8
		input  []uint8
		output []uint8
	}{
		{
			name:   "no missing numbers",
			size:   3,
			input:  []uint8{3, 2, 1},
			output: nil,
		},
		{
			name:   "one missing number",
			size:   3,
			input:  []uint8{3, 1},
			output: []uint8{2},
		},
		{
			name:   "all missing numbers",
			size:   3,
			output: []uint8{1, 2, 3},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := missingNumbers(test.input, test.size)
			if !reflect.DeepEqual(got, test.output) {
				t.Errorf("Wanted %v, got %v", test.output, got)
			}
		})
	}
}

//func TestCommonNumbers(t *testing.T) {
//	tests := []struct {
//		name   string
//		size   uint8
//		input  []uint8
//		output []uint8
//	}{
//		{name: "no common numbers"},
//		{name: "one common number"},
//		{name: "all common numbers"},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			got := commonNumbers(test.input)
//			if !reflect.DeepEqual(got, test.output) {
//				t.Error("")
//			}
//		})
//	}
//}
