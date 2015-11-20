package fastrandom

import "math/rand"

// return a pseudo-random number, uses the default source for random bits (golang library)
func randuint32(r uint32) uint32 {
	random32bit := rand.Uint32()
	if r&(r-1) == 0 {
		return random32bit & (r - 1)
	}
	if r > 0x80000000 { // if r > 1<<31
		for random32bit >= r {
			random32bit = rand.Uint32()
		}
		return random32bit // [0, r)
	}
	multiresult := uint64(random32bit) * uint64(r)
	leftover := uint32(multiresult)
	if leftover < r {
		threshold := uint32(0xFFFFFFFF) % r
		for leftover <= threshold {
			random32bit = rand.Uint32()
			multiresult = uint64(random32bit) * uint64(r)
			leftover = uint32(multiresult)
		}
	}
	return uint32(multiresult >> 32) // [0, r)
}

// return a pseudo-random number, uses the Mersenne Twister for random bits
func randuint32mt(r uint32, mt *Generator) uint32 {
	random32bit := mt.Next()
	if r&(r-1) == 0 {
		return random32bit & (r - 1)
	}
	if r > 0x80000000 { // if r > 1<<31
		for random32bit >= r {
			random32bit = rand.Uint32()
		}
		return random32bit // [0, r)
	}
	multiresult := uint64(random32bit) * uint64(r)
	leftover := uint32(multiresult)
	if leftover < r {
		threshold := uint32(0xFFFFFFFF) % r
		for leftover <= threshold {
			random32bit = rand.Uint32()
			multiresult = uint64(random32bit) * uint64(r)
			leftover = uint32(multiresult)
		}
	}
	return uint32(multiresult >> 32) // [0, r)
}
