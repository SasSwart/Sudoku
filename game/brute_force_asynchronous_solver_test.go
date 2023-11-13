package game

import (
	"fmt"
	"os"
	"runtime/pprof"
	"testing"
)

func BenchmarkBruteForceAsynchronousSolve(b *testing.B) {
	f, err := os.Create("async_cpu.prof")
	if err != nil {
		b.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		b.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < b.N; i++ {
		m := newModel(9)
		_, err := bruteForceAsynchronousSolve(&m)
		if err != nil {
			b.Error(err)
		}
	}
}

func TestBruteForceAsynchronousSolve(t *testing.T) {
	m := newModel(9)
	solved, err := bruteForceAsynchronousSolve(&m)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(solved)
}
