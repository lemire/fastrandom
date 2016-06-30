package fastrandom

import "github.com/davidminor/gorand/pcg"

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
