package main

import (
	gomap "github.com/diagmatrix/gomap/internal"
)

func main() {
	var h gomap.HeightMap = *gomap.NewHeightMap(1000, 1000)
	gomap.SaveHeightMap("./test/test.png", &h)
}
