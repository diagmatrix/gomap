package main

import (
	"math/rand"
	"time"

	gomap "github.com/diagmatrix/gomap/internal"
)

func main() {
	var h gomap.HeightMap = *gomap.NewHeightMapRN(1024, 1024)
	gomap.SaveHeightMap("./test/testRN.png", &h)
	h = *gomap.NewHeightMapPN(1024, 1024, 2, 2, 1, rand.NewSource(time.Now().Unix()))
	gomap.SaveHeightMap("./test/testPN.png", &h)
}
