package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input
var input []byte

//go:embed input-example
var inputExample []byte

func main() {
	mtx := bytes.Fields(inputExample)
	rows := len(mtx)
	cols := len(mtx[0])

	out1 := Task1(mtx, Point{0, 0}, Point{rows - 1, cols - 1})
	fmt.Println(out1 == 866, out1)
}

func Task1(input [][]byte, start, end Point) int {
	return NewBFS(input, start, end).Run()
}
