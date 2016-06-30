package fastrandom

import (
	"math/rand"
	"testing"
)
import "github.com/dgryski/go-pcgr"

var p = pcgr.Rand{0x0ddc0ffeebadf00d, 0xcafebabe}

//go test -bench=.

// standard
func benchmarkGo(b *testing.B, r uint32) {
	b.StopTimer()
	array := make([]int, r, r)
	for i := 0; i < int(r); i++ {
		array[i] = i
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {

		for i := r; i > 0; i-- {
			idx := rand.Int31n(int32(i))
			tmp := array[idx]
			array[idx] = array[i-1]
			array[i-1] = tmp
		}
	}
}

// our fast approach that avoids division
func benchmarkFastPCG_dgryski(b *testing.B, r uint32) {
	b.StopTimer()
	array := make([]int, r, r)
	for i := 0; i < int(r); i++ {
		array[i] = i
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		for i := r; i > 0; i-- {
			idx := randuint32pcg_dgryski(uint32(i), &p)
			tmp := array[idx]
			array[idx] = array[i-1]
			array[i-1] = tmp
		}
	}
}

// standard PCG
func benchmarkPCG_dgryski(b *testing.B, r uint32) {
	b.StopTimer()
	array := make([]int, r, r)
	for i := 0; i < int(r); i++ {
		array[i] = i
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		for i := r; i > 0; i-- {
			idx := p.Bound(uint32(i))
			tmp := array[idx]
			array[idx] = array[i-1]
			array[i-1] = tmp
		}
	}
}

func BenchmarkStandardShuffleWithGo1000(b *testing.B) {
	benchmarkGo(b, 1000)
}

func BenchmarkStandardShuffleWithPCGWithDivision1000_dgryski(b *testing.B) {
	benchmarkPCG_dgryski(b, 1000)
}

func BenchmarkStandardShuffleWithPCGButNoDivision1000_dgryski(b *testing.B) {
	benchmarkFastPCG_dgryski(b, 1000)
}
