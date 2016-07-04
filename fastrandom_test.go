package fastrandom

import (
	"fmt"
	"math/rand"
	"testing"
)
import "github.com/dgryski/go-pcgr"

var p = pcgr.Rand{0x0ddc0ffeebadf00d, 0xcafebabe}

func testFastShortShuffle(t *testing.T, key uint32, N uint) {
	for i := N; i >= 3; i-- {
		mn := precomputedMagicNumber[i-3]
		fa := precomputedFactorial[i-3]
		divresult := fastDiv(key, &mn) // value in [0,i)
		expresult := key / fa
		if divresult != expresult {
			t.Error("Expected ", expresult, ", got ", divresult)
		}
		key = key - divresult*fa // value in [0,(i-1)!)
	}
}

func TestFastDiv(t *testing.T) {
	N := uint(11)
	m := precomputedFactorial[N-2] - 1
	for ; m > 0; m-- {
		testFastShortShuffle(t, m, N)
	}
}

func toId(vals []uint, N uint) uint {
	v := uint(0)
	for i := uint(0); i < N; i++ {
		v = v*N + vals[i]
	}
	return v
}

func TestFastShuffle(t *testing.T) {
	N := uint(5)

	x := make(map[uint]bool)
	m := precomputedFactorial[N-2]
	for mm := uint32(0); mm < m; mm++ {
		sh := FastShortShuffle(mm, N)
		fmt.Println(sh)
		x[toId(sh, N)] = true
	}
	if uint32(len(x)) != precomputedFactorial[N-2] {
		t.Error("Expected ", precomputedFactorial[N-2], ", got ", len(x))
	}
}

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
			idx := Randuint32pcg_dgryski(uint32(i), &p)
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
