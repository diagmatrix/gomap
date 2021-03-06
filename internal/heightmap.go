package gomap

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/diagmatrix/gomap/internal/generator"
)

// ----------------------------------------------------------------------------
// Heightmap
type HeightMap struct {
	Pix  [][]uint8       // Map
	Rect image.Rectangle // Bounds
}

// Image interface functions
func (h *HeightMap) At(x, y int) color.Color {
	out := x < h.Rect.Min.X || x >= h.Rect.Dx() || y < h.Rect.Min.Y || y >= h.Rect.Dy()
	if out {
		return color.Gray{}
	} else {
		return color.Gray{uint8(h.Pix[x][y])}
	}
}
func (h *HeightMap) Bounds() image.Rectangle {
	return h.Rect
}
func (h *HeightMap) ColorModel() color.Model {
	return color.GrayModel
}

// ----------------------------------------------------------------------------
// Heightmap generator functions
// Random noise
func NewHeightMapRN(w, h, o int) *HeightMap {
	seed()
	matrix := make([][]uint8, w)
	for i := range matrix {
		matrix[i] = make([]uint8, h)
		for j := range matrix[i] {
			matrix[i][j] = uint8(rand.Int31n(256))
		}
	}
	return &HeightMap{
		Pix:  matrix,
		Rect: image.Rect(0, 0, w, h),
	}
}

// Diamond Square (works better with h,w = 2^k)
func NewHeightMapDS(s int, o, r uint8) *HeightMap {
	seed()
	// Empty matrix
	matrix := make([][]uint8, s)
	for i := range matrix {
		matrix[i] = make([]uint8, s)
		for j := range matrix[i] {
			matrix[i][j] = uint8(rand.Int31n(256))
		}
	}
	var n uint8
	for n = 1; n < o+1; n++ {
		matrix = generator.DiamondSquare(&matrix, s, r/n)
	}
	return &HeightMap{
		Pix:  matrix,
		Rect: image.Rect(0, 0, len(matrix), len(matrix)),
	}
}

// ----------------------------------------------------------------------------
// Other functions
// Saves the heightmap
func SaveHeightMap(s string, h *HeightMap) error {
	f, err := os.Create(s)
	if err != nil {
		// Failed opening/creating file
		return err
	}
	defer f.Close()
	err = png.Encode(f, h)
	if err != nil {
		// Failed encoding
		return err
	}
	return nil
}

// Initializes seed
func seed() {
	rand.Seed(time.Now().Unix())
}
