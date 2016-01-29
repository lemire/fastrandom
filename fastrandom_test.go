package fastrandom

import (
	"math/rand"
	"testing"
)
import "github.com/davidminor/gorand/pcg"

var p = pcg.NewPcg32(111)

var seed = uint32(123456789)

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

func benchmarkFastPCG4(b *testing.B, r uint32) {
	b.StopTimer()
	array := make([]int, r, r)
	for i := 0; i < int(r); i++ {
		array[i] = i
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		for i := r; i > 0; i-=4 {
			idx1 := randuint32pcg(uint32(i), &p)
			idx2 := randuint32pcg(uint32(i-1), &p)
			idx3 := randuint32pcg(uint32(i-2), &p)
			idx4 := randuint32pcg(uint32(i-3), &p)
			tmp1 := array[idx1]
			array[idx1] = array[i-1]
			array[i-1] = tmp1
			tmp2 := array[idx2]
			array[idx2] = array[i-2]
			array[i-2] = tmp2
			tmp3 := array[idx3]
			array[idx3] = array[i-3]
			array[i-3] = tmp3
			tmp4 := array[idx4]
			array[idx4] = array[i-4]
			array[i-4] = tmp4
		}
	}
}


func benchmarkFastPCG2(b *testing.B, r uint32) {
	b.StopTimer()
	array := make([]int, r, r)
	for i := 0; i < int(r); i++ {
		array[i] = i
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {

		for i := r; i > 0; i-=2 {
			idx1 := randuint32pcg(uint32(i), &p)
			idx2 := randuint32pcg(uint32(i-1), &p)
			tmp1 := array[idx1]
			array[idx1] = array[i-1]
			array[i-1] = tmp1
			tmp2 := array[idx2]
			array[idx2] = array[i-2]
			array[i-2] = tmp2
		}
	}
}


func benchmarkFastPCGalone(b *testing.B, r uint32) {
	b.StopTimer()
	z := uint32(0)
	b.StartTimer()
	for j := 0; j < b.N; j++ {

		for i := r; i > 0; i-- {
			idx := randuint32pcg(uint32(i), &p)
			z = z + idx
		}
	}
}

func benchmarkFastPCGaloneu(b *testing.B, r uint32) {
	b.StopTimer()
	z := uint32(0)
	b.StartTimer()
	for j := 0; j < b.N; j++ {

		for i := r; i > 0; i-- {
			idx := p.Next() % i
			z = z + idx
		}
	}
}
func benchmarkFastPCGaloneut(b *testing.B, r uint32) {
	b.StopTimer()
	z := uint32(0)
	b.StartTimer()
	for j := 0; j < b.N; j++ {

		for i := r; i > 0; i-- {
			idx := uint32(((uint64 (p.Next()))*uint64(i))>>32)
			z = z + idx
		}
	}
}


func benchmarkFastPCGaloneus(b *testing.B, r uint32) {
	b.StopTimer()
	z := uint32(0)
	b.StartTimer()
	for j := 0; j < b.N; j++ {

		for i := r; i > 0; i-- {
			idx := p.Next() 
			z = z + idx
		}
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
func benchmarkPCG4(b *testing.B, r uint32) {
	b.StopTimer()
	array := make([]int, r, r)
	for i := 0; i < int(r); i++ {
		array[i] = i
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		for i := r; i > 0; i-= 4 {
			idx1 := p.NextN(uint32(i))
			idx2 := p.NextN(uint32(i-1))
			idx3 := p.NextN(uint32(i-2))
			idx4 := p.NextN(uint32(i-3))
		 	tmp1 := array[idx1]
			array[idx1] = array[i-1]
			array[i-1] = tmp1
		 	tmp2 := array[idx2]
			array[idx2] = array[i-2]
			array[i-2] = tmp2
		 	tmp3 := array[idx3]
			array[idx3] = array[i-3]
			array[i-3] = tmp3
		 	tmp4 := array[idx4]
			array[idx4] = array[i-4]
			array[i-4] = tmp4
		}
	}
}


func benchmarkFlip(b *testing.B, r uint32) {
        b.StopTimer()
        array := make([]int, r, r)
        for i := 0; i < int(r); i++ {
                array[i] = i
        }
        b.StartTimer()
        for j := 0; j < b.N; j++ {
                for i := r; i > 0; i-- {
                        idx := r-i
                        tmp := array[idx]
                        array[idx] = array[i-1]
                        array[i-1] = tmp
                }
        }
}


func BenchmarkPCG1000(b *testing.B) {
	benchmarkPCG(b, 1000)
}
func BenchmarkPCG4_1000(b *testing.B) {
	benchmarkPCG4(b, 1000)
}
func BenchmarkFastPCGalone1000(b *testing.B) {
	benchmarkFastPCGalone(b, 1000)
}
func BenchmarkFastPCGaloneu1000(b *testing.B) {
	benchmarkFastPCGaloneu(b, 1000)
}
func BenchmarkFastPCGaloneus1000(b *testing.B) {
	benchmarkFastPCGaloneus(b, 1000)
}
func BenchmarkFastPCGaloneut1000(b *testing.B) {
	benchmarkFastPCGaloneut(b, 1000)
}
func BenchmarkFastPCG1000(b *testing.B) {
	benchmarkFastPCG(b, 1000)
}
func BenchmarkFastPCG2_1000(b *testing.B) {
	benchmarkFastPCG2(b, 1000)
}
func BenchmarkFastPCG4_1000(b *testing.B) {
	benchmarkFastPCG4(b, 1000)
}
func BenchmarkFlip1000(b *testing.B) {
	benchmarkFlip(b, 1000)
}


