package gomap

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/diagmatrix/gomap/internal/pnoise"
)

// Heightmap
type HeightMap struct {
	Pix  [][]uint8       // Map
	Rect image.Rectangle // Bounds
}

var _time float64 = 0

// ----------------------------------------------------------------------------
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
func NewHeightMapRN(w, h int) *HeightMap {
	seed()
	matrix := make([][]uint8, w)
	for i := range matrix {
		matrix[i] = make([]uint8, h)
		for j := range matrix[i] {
			matrix[i][j] = uint8(rand.Int31n(65535))
		}
	}
	return &HeightMap{
		Pix:  matrix,
		Rect: image.Rect(0, 0, w, h),
	}
}

func NewHeightMapPN(w, h int, a, b float64, o int32, src rand.Source) *HeightMap {
	nm, min, max := pnoise.NoiseMap(w, h, a, b, o, src)
	matrix := make([][]uint8, w)
	for i := 0; i < w; i++ {
		matrix[i] = make([]uint8, h)
		for j := 0; j < h; j++ {
			matrix[i][j] = toGrayscale(nm[i][j], min, max)
		}
	}

	return &HeightMap{
		Pix:  matrix,
		Rect: image.Rect(0, 0, w, h),
	}
}

// ----------------------------------------------------------------------------
// Other functions
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
func seed() {
	rand.Seed(time.Now().Unix())
}
func toGrayscale(t, min, max float64) uint8 {
	return uint8(((t - min) / (max - min)) * 255)
}
