// This code is an adaptation based on:
// http://git.gnome.org/browse/gegl/tree/operations/common/perlin
package pnoise

import (
	"math"
	"math/rand"
)

//-----------------------------------------------------------------------------
// Constants and variables for perlin noise
const (
	B  = 0x100  //256
	N  = 0x1000 //4096
	BM = 0xFF   //255
	NP = 12
	NM = 0xFFF // 4095
)

type Perlin2 struct {
	alpha float64
	beta  float64
	n     int32
	p     [B + B + 2]int32
	g     [B + B + 2][2]float64
}

func NewPerlin2Rand(alpha, beta float64, n int32, src rand.Source) *Perlin2 {
	var p Perlin2
	p.alpha = alpha
	p.beta = beta
	p.n = n
	r := rand.New(src)
	var i int32

	for i = 0; i < B; i++ {
		p.p[i] = i
		for j := 0; j < 2; j++ {
			p.g[i][j] = float64(r.Int31n(B+B)-B) / B
		}
		normalize(&p.g[i])
	}
	for ; i > 0; i-- {
		k := r.Int31n(B)
		p.p[i], p.p[k] = p.p[k], p.p[i]
	}
	for i = 0; i < B+2; i++ {
		p.p[B+i] = p.p[i]
		for j := 0; j < 2; j++ {
			p.g[B+i][j] = p.g[i][j]
		}
	}
	return &p
}

func (p *Perlin2) Noise(x, y float64) float64 {
	var scale float64 = 1
	var sum float64 = 0
	var freq float64 = 1
	var max float64 = 0
	for i := 0; i < int(p.n); i++ {
		val := noise(x*freq, y*freq, p)
		sum += val / scale
		scale *= p.alpha
		freq *= p.beta
		max += scale
	}
	return sum / scale
}

func NoiseMap(w, h int, a, b float64, o int32, src rand.Source) ([][]float64, float64, float64) {
	p := NewPerlin2Rand(a, b, o, src)
	m := make([][]float64, w)
	var min, max float64 = 0, 0
	for y := range m {
		m[y] = make([]float64, h)
		for x := range m[y] {
			var dx float64 = float64(x) / float64(h)
			var dy float64 = float64(y) / float64(w)
			val := p.Noise(dx, dy)
			m[y][x] = val
			if x == 0 {
				min, max = val, val
			}
			if val < min {
				min = val
			}
			if val > max {
				max = val
			}
		}
	}
	return m, min, max
}

// ----------------------------------------------------------------------------
// Auxiliary functions
func fade(t float64) float64 {
	return t * t * t * (10 + t*(6*t-15)) // 6t^5-15t^4+10t^3
}
func lerp(t, a, b float64) float64 {
	return a + t*(b-a)
}
func normalize(v *[2]float64) {
	mod := math.Sqrt(v[0]*v[0] + v[1]*v[1])
	v[0], v[1] = v[0]/mod, v[1]/mod
}
func at(x, y float64, v [2]float64) float64 {
	return x*v[0] + y*v[1]
}
func noise(v0, v1 float64, p *Perlin2) float64 {
	var t [2]float64 = [2]float64{v0 + N, v1 + N}

	var bx0 int32 = int32(t[0]) & BM
	var bx1 int32 = (bx0 + 1) & BM
	var by0 int32 = int32(t[1]) & BM
	var by1 int32 = (by0 + 1) & BM

	var rx0 float64 = t[0] - math.Floor(t[0])
	var rx1 float64 = rx0 - 1
	var ry0 float64 = t[1] - math.Floor(t[1])
	var ry1 float64 = rx0 - 1

	var i int32 = p.p[bx0]
	var j int32 = p.p[bx1]

	var b00 int32 = p.p[i+by0]
	var b10 int32 = p.p[j+by0]
	var b01 int32 = p.p[i+by1]
	var b11 int32 = p.p[j+by1]

	var sx float64 = fade(rx0)
	var sy float64 = fade(ry0)

	var ua float64 = at(rx0, ry0, p.g[b00])
	var va float64 = at(rx1, ry0, p.g[b10])
	var a float64 = lerp(sx, ua, va)

	var ub float64 = at(rx0, ry1, p.g[b01])
	var vb float64 = at(rx1, ry1, p.g[b11])
	var b float64 = lerp(sx, ub, vb)

	return lerp(sy, a, b)
}
