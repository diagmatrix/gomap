package main

import (
	"fmt"

	gomap "github.com/diagmatrix/gomap/internal"
	"github.com/diagmatrix/gomap/internal/smoothing"
)

func main() {
	// 8 octaves 128 amplitude
	for i := 8; i <= 8; i++ {
		for j := 127; j >= 127; j = j / 2 {
			h := *gomap.NewHeightMapDS(256, uint8(i), uint8(j))
			gomap.SaveHeightMap("./test/testO"+fmt.Sprint(i)+"A"+fmt.Sprint(j)+".png", &h)
			nh := h
			var count int = 0
			for k := 10; k > 1; k-- {
				count++
				fmt.Println("Octaves: ", i, "\tAmplitude: ", j, "\t Kernel size: ", 5, "\tPass: ", count)
				nh.Pix = smoothing.SmoothKNN(&nh.Pix, 5)
				gomap.SaveHeightMap("./test/testO"+fmt.Sprint(i)+"A"+fmt.Sprint(j)+"K"+fmt.Sprint(5)+"P"+fmt.Sprint(count)+".png", &nh)
			}
		}
	}

}
