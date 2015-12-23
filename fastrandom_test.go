package fastrandom

import (
	"math/rand"
	"testing"
)
import "github.com/davidminor/gorand/pcg"

var rgen = Init(4111)
var p = pcg.NewPcg32(111)

//go test -bench=.
func benchmarkStandardNoRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int31()
	}
}
func benchmarkStandard(b *testing.B, r int32) {
	for i := 0; i < b.N; i++ {
		rand.Int31n(r)
	}
}
func benchmarkFast(b *testing.B, r uint32) {
	for i := 0; i < b.N; i++ {
		randuint32(r)
	}
}

func benchmarkDivBased(b *testing.B, r uint32) {
	for i := 0; i < b.N; i++ {
		divbased_randuint32(r)
	}
}

//go test -bench=.
func benchmarkMTNoRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rgen.Next()
	}
}

func benchmarkDivBasedMT(b *testing.B, r uint32) {
	for i := 0; i < b.N; i++ {
		divbased_randuint32mt(r, rgen)
	}
}

func benchmarkFastMT(b *testing.B, r uint32) {
	for i := 0; i < b.N; i++ {
		randuint32mt(r, rgen)
	}
}

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

func BenchmarkPCG1000(b *testing.B) {
	benchmarkPCG(b, 1000)
}
func BenchmarkFastPCG1000(b *testing.B) {
	benchmarkFastPCG(b, 1000)
}

func BenchmarkPCG1000000(b *testing.B) {
	benchmarkPCG(b, 1000000)
}
func BenchmarkFastPCG1000000(b *testing.B) {
	benchmarkFastPCG(b, 1000000)
}


func BenchmarkPCG10000000(b *testing.B) {
	benchmarkPCG(b, 10000000)
}
func BenchmarkFastPCG10000000(b *testing.B) {
	benchmarkFastPCG(b, 10000000)
}