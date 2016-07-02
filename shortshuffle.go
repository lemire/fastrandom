package fastrandom

import "fmt"

///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
// This file contains a fast function to compute permutations of up to 12 elements
// without *any* division from a single 32-bit key
///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////


// precomputed factorials... 2!, 3!..., 12!
var precomputedFactorial = [11]uint32{2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800, 39916800, 479001600}

// magic number of fast division
type magicNumber struct {
	multiplier uint32
	needAdd    bool
	shift      uint
}

// magic numbers to divide by... 2!, 3!..., 12!
var precomputedMagicNumber = [11]magicNumber{{0x80000000, false, 0}, {0xaaaaaaab, false, 2}, {0xaaaaaaab, false, 4}, {0x88888889, false, 6}, {0x6c16c16d, true, 9}, {0xa01a01a1, true, 12}, {0xa01a01a1, true, 15}, {0xb8ef1d2b, false, 18}, {0x24fc9f6f, false, 19}, {0x035cc8ad, false, 19}, {0x011eed8f, false, 21}}

// (x*y) >> 32
func high32(x uint32, y uint32) uint32 {
	return uint32((uint64(x) * uint64(y)) >> 32)
}

// fast division function
func fastDiv(val uint32, mn *magicNumber) uint32 {
	q := high32(mn.multiplier, val)
	if mn.needAdd {
		q = ((val - q) >> 1) + q
	}
	return q >> mn.shift
}

// Given a key in [0,N!) generates a shuffle *without* using any division.
// You can generate a value in [0,N!) using randuint32pcg_dgryski
// you probably don't want to use such a function... this is just a demo
func FastShortShuffle(key uint32, N uint) []uint {
	if (N < 2) || (N > 12) {
		fmt.Errorf("out of bound")
	}
	if key >= precomputedFactorial[N-2] {
		fmt.Errorf("bad key")
	}
	answer := make([]uint, N)
	for i := uint(0); i < N; i++ {
		answer[i] = i // fill answer with values 0,1,2...,N-1
	}
	for i := N; i >= 3; i-- {
		mn := precomputedMagicNumber[i-3]
		fa := precomputedFactorial[i-3]
		divresult := fastDiv(key, &mn) // value in [0,i)
		key = key - divresult*fa       // value in [0,(i-1)!)
		// we swap
		tmp := answer[divresult]
		answer[divresult] = answer[i-1]
		answer[i-1] = tmp
	}
	// here i = 2, key is zero or one
	tmp := answer[key]
	answer[key] = answer[1]
	answer[1] = tmp
	return answer
}
