package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)

	for i := range r {
		c := make([]uint8, dx)
		r[i] = c
		for j := range c {
			r[i][j] = uint8(i ^ j)
		}
	}

	return r
}

func main() {
	pic.Show(Pic)
}
