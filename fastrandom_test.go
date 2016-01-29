package fastrandom

import (
	"testing"
)
import "github.com/davidminor/gorand/pcg"

var p = pcg.NewPcg32(111)

var seed = uint32(123456789)

//go test -bench=.

// precomputes the random numbers
func benchmarkPrePCG(b *testing.B, r uint32) {
	b.StopTimer()
	array := make([]int, r, r)
	randarray := make([]uint32, r+1, r+1)
	for i := 0; i < int(r); i++ {
		array[i] = i
	}
	for i := r; i > 0; i-- {
		randarray[i] = randuint32pcg(uint32(i), &p)
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {

		for i := r; i > 0; i-- {
			idx := randarray[i]
			tmp := array[idx]
			array[idx] = array[i-1]
			array[i-1] = tmp
		}
	}
}


// our fast approach that avoids division
func benchmarkFastPCG(b *testing.B, r uint32) {
	b.StopTimer()
	array := make([]int, r, r)
	for i := 0; i < int(r); i++ {
		array[i] = i
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {

		for i := r; i > 0; i-- {
			idx := randuint32pcg(uint32(i), &p)
			tmp := array[idx]
			array[idx] = array[i-1]
			array[i-1] = tmp
		}
	}
}

// standard PCG
func benchmarkPCG(b *testing.B, r uint32) {
	b.StopTimer()
	array := make([]int, r, r)
	for i := 0; i < int(r); i++ {
		array[i] = i
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		for i := r; i > 0; i-- {
			idx := p.NextN(uint32(i))
			tmp := array[idx]
			array[idx] = array[i-1]
			array[i-1] = tmp
		}
	}
}

func BenchmarkStandardShuffleWithPCGWithDivision1000(b *testing.B) {
	benchmarkPCG(b, 1000)
}
func BenchmarkShuffleWithPrecomputedRandomNumbers1000(b *testing.B) {
	benchmarkPrePCG(b, 1000)
}

func BenchmarkStandardShuffleWithPCGButNoDivisionFastPCG1000(b *testing.B) {
	benchmarkFastPCG(b, 1000)
}

