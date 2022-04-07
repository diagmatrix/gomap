package main

import (
	"github.com/diagmatrix/gomap/internal/heightmap"
)

func main() {
	var h heightmap.HeightMap = *heightmap.NewHeightMapRN(1024, 1024)
	heightmap.SaveHeightMap("./test/testRN.png", &h)
	h = *heightmap.NewHeightMapDS(129, 80)
	/*
		for i := range h.Pix {
			var row string = ""
			for _, v := range h.Pix[i] {
				row += fmt.Sprint(v) + " "
			}
			fmt.Println(row)
		}
	*/
	heightmap.SaveHeightMap("./test/testDS.png", &h)
}
