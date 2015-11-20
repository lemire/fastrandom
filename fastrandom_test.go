package fastrandom

import (
	"math/rand"
	"testing"
)

var rgen = Init(4111)

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

//go test -bench=.
func benchmarkMTNoRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rgen.Next()
	}
}
func benchmarkFastMT(b *testing.B, r uint32) {
	for i := 0; i < b.N; i++ {
		randuint32mt(r, rgen)
	}
}

func BenchmarkStandard(b *testing.B) {
	benchmarkStandardNoRange(b)
}

func BenchmarkStandardMT(b *testing.B) {
	benchmarkMTNoRange(b)
}
func BenchmarkStandard100(b *testing.B) {
	benchmarkStandard(b, 100)
}

func BenchmarkFast100(b *testing.B) {
	benchmarkFast(b, 100)
}

func BenchmarkFastMT100(b *testing.B) {
	benchmarkFastMT(b, 100)
}

func BenchmarkStandard1000(b *testing.B) {
	benchmarkStandard(b, 1000)
}

func BenchmarkFast1000(b *testing.B) {
	benchmarkFast(b, 1000)
}

func BenchmarkFastMT1000(b *testing.B) {
	benchmarkFastMT(b, 1000)
}

func BenchmarkStandard1000000(b *testing.B) {
	benchmarkStandard(b, 1000000)
}

func BenchmarkFast1000000(b *testing.B) {
	benchmarkFast(b, 1000000)
}

func BenchmarkFastMT1000000(b *testing.B) {
	benchmarkFastMT(b, 1000000)
}
