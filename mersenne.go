package fastrandom

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
