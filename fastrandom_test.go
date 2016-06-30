package fastrandom

import (
	"testing"
        "math/rand"
)
import "github.com/davidminor/gorand/pcg"

var p = pcg.NewPcg32(111)

var seed = uint32(123456789)

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
func benchmarkFastPCG(b *testing.B, r uint32) {
	b.StopTimer()
	array := make([]int, r, r)
	for i := 0; i < int(r); i++ {
		array[i] = i
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		for i := r; i > 0; i-- {
			// next line is much faster which suggests inlining issue
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
func BenchmarkStandardShuffleWithGo1000(b *testing.B) {
	benchmarkGo(b, 1000)
}

func BenchmarkStandardShuffleWithPCGWithDivision1000(b *testing.B) {
	benchmarkPCG(b, 1000)
}

func BenchmarkStandardShuffleWithPCGButNoDivision1000(b *testing.B) {
	benchmarkFastPCG(b, 1000)
}

