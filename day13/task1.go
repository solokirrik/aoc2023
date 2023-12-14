package main

import (
	"bytes"
)

const doDebug = false

func Task1(squares [][]byte) int {
	res := 0

	for i, square := range squares {
		row, col := processSquare(i, bytes.Split(square, []byte("\n")))
		res += col + (100 * row)
	}

	return res
}

func processSquare(i int, square [][]byte) (symmetryRow, symmetryCol int) {
	symmetryRow = chechHorizontal(square)
	symmetryCol = chechVertical(square)

	return addOneIfLegal(symmetryRow), addOneIfLegal(symmetryCol)
}

func addOneIfLegal(in int) int {
	if in == -1 {
		return 0
	}
	return in + 1
}

// find vertical symmetry line = simmetryc horizontally
func chechVertical(square [][]byte) int {
	r := 0
	c := 0
	isSymmetry := false

	for r < len(square) && c < len(square[0]) {
		colLeft := c
		colRight := colLeft + 1

		for colLeft >= 0 && colRight < len(square[r]) {
			if square[r][colLeft] == square[r][colRight] {
				colRight++
				colLeft--
				isSymmetry = true
			} else {
				isSymmetry = false
				break
			}
		}
		if isSymmetry {
			r++
		} else {
			r = 0
			c = (colLeft+colRight)/2 + 1
		}
	}

	if isSymmetry {
		return c
	}

	return -1
}

// find horizontal symmetry line = simmetryc vertically
func chechHorizontal(square [][]byte) int {
	c := 0
	r := 0
	isSymmetry := false

	for c < len(square[0]) && r < len(square) {
		top := r
		bottom := r + 1

		for top >= 0 && bottom < len(square) {
			if square[top][c] == square[bottom][c] {
				top--
				bottom++
				isSymmetry = true
			} else {
				isSymmetry = false
				break
			}
		}
		if isSymmetry {
			c++
		} else {
			c = 0
			r = (top+bottom)/2 + 1
		}
	}

	if isSymmetry {
		return r
	}

	return -1
}
