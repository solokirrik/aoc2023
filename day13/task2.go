package main

import "bytes"

// Does not work
func Task2(squares [][]byte) int {
	res := 0

	for i, square := range squares {
		row, col := processSquare(i, bytes.Split(square, []byte("\n")))
		res += col + (100 * row)
	}

	return res
}
