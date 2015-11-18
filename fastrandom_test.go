package fastrandom

import (
	"math/rand"
	"testing"
)

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

func BenchmarkStandard(b *testing.B) {
	benchmarkStandardNoRange(b)
}

func BenchmarkStandard100(b *testing.B) {
	benchmarkStandard(b, 100)
}

func BenchmarkFast100(b *testing.B) {
	benchmarkFast(b, 100)
}

func BenchmarkStandard1000(b *testing.B) {
	benchmarkStandard(b, 100)
}

func BenchmarkFast1000(b *testing.B) {
	benchmarkFast(b, 100)
}

func BenchmarkStandard1073741824(b *testing.B) {
	benchmarkStandard(b, 100)
}

func BenchmarkFast1073741824(b *testing.B) {
	benchmarkFast(b, 100)
}
