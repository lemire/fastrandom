package fastrandom

import "math/rand"

// return a pseudo-random number, uses the default source for random bits (golang library)
func randuint32(r uint32) uint32 {
	random32bit := rand.Uint32()
	if r > 0x80000000 { // if r > 1<<31
		for random32bit >= r {
			random32bit = rand.Uint32()
		}
		return random32bit // [0, r)
	}
	multiresult := uint64(random32bit) * uint64(r)
	candidate := multiresult >> 32
	leftover := uint32(multiresult)
	if leftover > -r-1 { //2^32 -r +lsbset <= leftover
		threshold := uint32((uint64(1)<<32)/uint64(r))*r - 1
		random32bit = rand.Uint32()
		multiresult = uint64(random32bit) * uint64(r)
		candidate = multiresult >> 32
		leftover = uint32(multiresult)

		for leftover > threshold {
			random32bit = rand.Uint32()
			multiresult = uint64(random32bit) * uint64(r)
			candidate = multiresult >> 32
			leftover = uint32(multiresult)
		}
	}
	return uint32(candidate) // [0, r)
}
