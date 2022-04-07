package main

import (
	"fmt"

	gomap "github.com/diagmatrix/gomap/internal"
)

func main() {
	// 8 octaves 128 amplitude
	for i := 1; i < 9; i++ {
		for j := 255; j > 16; j = j / 2 {
			fmt.Println("Octaves: ", i, "\tAmplitude: ", j)
			h := *gomap.NewHeightMapDS(512, uint8(i), uint8(j))
			gomap.SaveHeightMap("./test/test("+fmt.Sprint(i)+"-"+fmt.Sprint(j)+").png", &h)
			h = *gomap.NewHeightMapDSS(512, uint8(i), uint8(j))
			gomap.SaveHeightMap("./test/testS("+fmt.Sprint(i)+"-"+fmt.Sprint(j)+").png", &h)
		}
	}

}
