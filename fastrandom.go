package fastrandom

import "math/rand"
import "github.com/davidminor/gorand/pcg"

// uses the standard Go algorithm based on divisions to generate a number in [0,r), uses the default source for random bits (golang library)
func divbased_randuint32(r uint32) uint32 {
	random32bit := rand.Uint32()
	if r&(r-1) == 0 {
		return random32bit & (r - 1)
	}
	max := uint32(0xFFFFFFFF) % r
	for random32bit <= max {
		random32bit = rand.Uint32()
	}
	return random32bit % r
}



// return a pseudo-random number in [0,r), avoiding divisions, uses the default source for random bits (golang library)
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


// return a pseudo-random number in [0,r), avoiding divisions, uses the PCG for random bits
func fancyranduint32pcg(r uint32, pcg *pcg.Pcg32) uint32 {
        random32bit := pcg.Next()
        if r&(r-1) == 0 {
                return random32bit & (r - 1)
        }
        if r > 0x80000000 { // if r > 1<<31
                for random32bit >= r {
                        random32bit = pcg.Next()
                }
                return random32bit // [0, r)
        }
        multiresult := uint64(random32bit) * uint64(r)
        leftover := uint32(multiresult)
        if leftover < r {
                threshold := uint32(0xFFFFFFFF) % r
                for leftover <= threshold {
                        random32bit = pcg.Next()
                        multiresult = uint64(random32bit) * uint64(r)
                        leftover = uint32(multiresult)
                }
        }
        return uint32(multiresult >> 32) // [0, r)
}


// return a pseudo-random number in [0,r), avoiding divisions, uses the PCG for random bits
func randuint32pcg(r uint32, pcg *pcg.Pcg32) uint32 {
	random32bit := pcg.Next()
	multiresult := uint64(random32bit) * uint64(r)
	leftover := uint32(multiresult) 
	if leftover < r {
		threshold := uint32(-r) % r
		for leftover < threshold {
			random32bit = pcg.Next()
			multiresult = uint64(random32bit) * uint64(r)
			leftover = uint32(multiresult)
		}
	}
	return uint32(multiresult >> 32) // [0, r)
}
