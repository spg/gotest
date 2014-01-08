package main

import (
	"code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	table := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		table[i] = make([]uint8, dx)
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			table[i][j] = uint8((i + j) / 2)
		}
	}

	return table
}

func main() {
	pic.Show(Pic)
}
