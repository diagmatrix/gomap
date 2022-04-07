package smoothing

import (
	"math"
)

// Smoothing KNN
func SmoothKNN(hm *[][]uint8, k int) [][]uint8 {
	// Kernel heightmap
	var w, h int = len(*hm), len((*hm)[0])
	var res [][]float64 = make([][]float64, w)
	for i := range res {
		res[i] = make([]float64, h)
	}
	var rad int = k / 2
	var dmax float64 = float64(rad) * math.Sqrt2
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			// Kernel creation (k x k)
			var xmin, xmax int = max(0, x-rad), min(w-1, x+rad)
			var ymin, ymax int = max(0, y-rad), min(h-1, y+rad)
			var sum float64 = 0
			var count int = 0
			for i := xmin; i <= xmax; i++ {
				for j := ymin; j <= ymax; j++ {
					// Weight depends on distance
					sum += (dist(x, y, i, j) / dmax) * float64((*hm)[i][j])
					count++
				}
			}
			var kernel_avg float64 = sum / float64(count)
			res[x][y] = kernel_avg
		}
	}
	// Result heightmap
	var nhm [][]uint8 = make([][]uint8, w)
	for i := range nhm {
		nhm[i] = make([]uint8, h)
		for j := range nhm[i] {
			var val uint8 = uint8(res[i][j] * (float64(k+rad+1) / float64(k))) // It's a hack
			nhm[i][j] = uint8(val)
		}
	}

	return nhm
}

/*
TODO:
	- Implement other smoothing functions
	- Fix the hack
*/

// ----------------------------------------------------------------------------
// Auxiliary functions

// Min
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Max
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Normalized distance
func dist(x0, y0, x1, y1 int) float64 {
	return math.Sqrt(float64(x0-x1)*float64(x0-x1) + float64(y0-y1)*float64(y0-y1))
}
