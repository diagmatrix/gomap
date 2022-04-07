package heightmap

import (
	"math"
	"math/rand"
)

func DiamondSquare(s int, o uint8, m *[][]uint8) {
	var squareSize int = s
	var halfSquare int = squareSize / 2
	// Initial corners
	(*m)[0][0] = 128
	(*m)[0][s-1] = 128
	(*m)[s-1][0] = 128
	(*m)[s-1][s-1] = 128

	var offset uint8 = o

	for squareSize > 2 {
		// 1. Diamond
		/*
			c0 --- c1
			| \   /|
			|  cd  |
			| /   \|
			c3 --- c2
		*/
		for x := halfSquare; x < s; x += (squareSize - 1) {
			for y := halfSquare; y < s; y += (squareSize - 1) {
				// Corners
				var xmin, xmax int = x - halfSquare, x + halfSquare
				var ymin, ymax int = y - halfSquare, y + halfSquare
				var c0, c1, c2, c3 uint8 = (*m)[xmin][ymin], (*m)[xmax][ymin], (*m)[xmax][ymax], (*m)[xmin][ymax]
				var cd uint8 = (c0+c1+c2+c3)/4 + uint8(rand.Int31n(int32(offset)*2+1)-int32(offset))
				(*m)[x][y] = cd
			}
		}
		// 2. Square
		/*
						   c01
						   |
					c0 --- s0 --- c1
					|      |      |
			c30 --- s3 --- cd --- s1 --- c12
					|      |      |
					c3 --- s2 --- c2
						   |
						   c23
		*/
		for x := 0; x < s; x += halfSquare {
			for y := (x + halfSquare) % (squareSize - 1); y < s; y += (squareSize - 1) {
				// Corners
				var cN, cE, cS, cW uint8 = 0, 0, 0, 0 // c01,c1,cd,c0 in the picture for s0
				var count uint8 = 0                   // Number of valid corners {3,4}
				if y-halfSquare > 0 {
					cN = (*m)[x][y-halfSquare]
					count += 1
				}
				if x+halfSquare < s {
					cE = (*m)[x+halfSquare][y]
					count += 1
				}
				if y+halfSquare < s {
					cS = (*m)[x][y+halfSquare]
					count += 1
				}
				if x-halfSquare > 0 {
					cW = (*m)[x-halfSquare][y]
					count += 1
				}
				(*m)[x][y] = (cN+cE+cS+cW)/count + uint8(rand.Int31n(int32(offset)*2+1)-int32(offset))
			}
		}
		squareSize = int(math.Ceil(float64(squareSize) / 2))
		halfSquare = squareSize / 2
		offset = offset / 2
		if offset == 0 {
			offset = 1
		}
	}
}
