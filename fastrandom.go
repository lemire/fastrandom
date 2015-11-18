package fastrandom

import "math/rand"

/// "github.com/CasualSuperman/Mersenne-Twister"
type Generator struct {
	state [624]uint32
	index uint32
}

func Init(a uint32) *Generator {
	var t Generator
	t.init(a)
	return &t
}

func (g *Generator) init(a uint32) {
	g.state[0] = a
	var i uint32 = 0
	for i = 1; i < 624; i++ {
		g.state[i] = uint32(0x6c078965*(g.state[i-1]^
			g.state[i-1]>>30) + i)
	}
}

func (g *Generator) Next() uint32 {
	if g.index == 0 {
		g.generateNumbers()
	}
	y := g.state[g.index]
	g.index++
	g.index %= 624
	y ^= y >> 11
	y ^= y << 7 & 0x9d2c5680
	y ^= y << 15 & 0xefc60000
	y ^= y >> 18
	return y
}
func (g *Generator) generateNumbers() {
	var i uint32
	for i = 0; i < 624; i++ {
		y := g.state[i]>>31 + g.state[(i+1)%624]&0x7fffffff
		g.state[i] = g.state[(i+397)%624] ^ y>>1
		if y%2 == 1 {
			g.state[i] ^= 0x9908b0df
		}
	}
}

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
	if leftover > -r-1 { //2^32 -r +lsbset <= leftover
		threshold := uint32(0xFFFFFFFF)/r*r - 1 //uint32((uint64(1)<<32)/uint64(r))*r - 1
		random32bit = rand.Uint32()
		multiresult = uint64(random32bit) * uint64(r)
		leftover = uint32(multiresult)

		for leftover > threshold {
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
	if leftover > -r-1 { //2^32 -r +lsbset <= leftover
		threshold := uint32(0xFFFFFFFF)/r*r - 1 //uint32((uint64(1)<<32)/uint64(r))*r - 1
		random32bit = mt.Next()
		multiresult = uint64(random32bit) * uint64(r)
		leftover = uint32(multiresult)

		for leftover > threshold {
			random32bit = rand.Uint32()
			multiresult = uint64(random32bit) * uint64(r)
			leftover = uint32(multiresult)
		}
	}
	return uint32(multiresult >> 32) // [0, r)
}
