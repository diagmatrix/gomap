package gomap

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Heightmap -> RGBA image
func HeightMapToImage(m *HeightMap, sl uint8) *image.RGBA {
	var w, h int = len(m.Pix), len(m.Pix[0])
	mapper := mapperSeaLevel(sl)
	var res *image.RGBA = image.NewRGBA(image.Rect(0, 0, w, h))
	for i := 0; i < res.Bounds().Max.X; i++ {
		for j, alt := range m.Pix[i] {
			res.Set(i, j, mapper(alt))
		}
	}
	return res
}

// Saves a colored version of the heightmap
func SaveHeightMapColored(s string, h *HeightMap, sl uint8) error {
	f, err := os.Create(s)
	if err != nil {
		// Failed opening/creating file
		return err
	}
	defer f.Close()
	err = png.Encode(f, HeightMapToImage(h, sl))
	if err != nil {
		// Failed encoding
		return err
	}
	return nil
}

// ----------------------------------------------------------------------------
// Auxiliaries
var (
	DEEPWATER = color.RGBA{0, 83, 209, 255}
	WATER     = color.RGBA{45, 118, 228, 255}
	SHORE     = color.RGBA{100, 193, 240, 255}
	LAND      = color.RGBA{81, 168, 95, 255}
	MOUNTAIN  = color.RGBA{115, 66, 12, 255}
)

func mapperSeaLevel(sl uint8) func(i uint8) color.RGBA {
	var d uint8 = sl / 3
	var w uint8 = 2 * sl / 3
	var s uint8 = sl
	var l uint8 = (3 * sl / 2) % 255
	var m uint8 = (2 * sl) % 255
	return func(i uint8) color.RGBA {
		if i < d {
			return DEEPWATER
		} else if i < w {
			return WATER
		} else if i < s {
			return SHORE
		} else if i < l {
			return LAND
		} else if i < m {
			return MOUNTAIN
		} else {
			return color.RGBA{255, 255, 255, 255}
		}
	}
}
