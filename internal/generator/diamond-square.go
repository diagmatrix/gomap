package generator

import (
	"math/rand"
)

/*
Diamond-Square algorithm.
	base: Heightmap
	s: Size (s x s)
	r: octave
Returns a heightmap size 2*s-1
*/
func DiamondSquare(base *[][]uint8, s int, r uint8) [][]uint8 {
	var n int = 2*s - 1
	// The new heightmap
	var hm [][]uint8 = make([][]uint8, n)
	for i := range hm {
		hm[i] = make([]uint8, n)
	}
	// Resize
	for i := 0; i < n; i += 2 {
		for j := boolC(!(i%2 == 0)); j < n; j += 2 {
			hm[i][j] = (*base)[i/2][j/2]
		}
	}
	// Diamond
	/*
		0 0 0
		0 X 0
		0 0 0
	*/
	for i := 1; i < n; i += 2 {
		for j := 1; j < n; j += 2 {
			var a, b, c, d uint8 = hm[i-1][j-1], hm[i-1][j+1], hm[i+1][j-1], hm[i+1][j+1] // Corners
			var avg int = int(a+b+c+d) / 4
			avg += randInt(r)
			if avg > 255 {
				avg = 255
			} else if avg < 0 {
				avg = 0
			}
			hm[i][j] = uint8(avg)
		}
	}
	// Square
	/*
		0 + 0
		+ 0 +
		0 + 0
	*/
	for i := 0; i < n; i++ {
		for j := boolC(i%2 == 0); j < n; j += 2 {
			// Surrounding values
			var a, b, c, d uint8 = 0, 0, 0, 0
			var count int = 0
			if i != 0 {
				a = hm[i-1][j]
				count++
			}
			if j != 0 {
				b = hm[i][j-1]
				count++
			}
			if j+1 != n {
				c = hm[i][j+1]
				count++
			}
			if i+1 != n {
				d = hm[i+1][j]
				count++
			}
			var avg int = int(a+b+c+d) / count
			avg += randInt(r)
			if avg > 255 {
				avg = 255
			} else if avg < 0 {
				avg = 0
			}
			hm[i][j] = uint8(avg)
		}
	}
	return hm
}

// Makes a random int in [-r,r]
func randInt(r uint8) int {
	return -int(r) + 2*rand.Intn(int(r)+1)
}
func boolC(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
