package main

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/solokirrik/aoc2023/ops"
)

//go:embed input
var input []byte

//go:embed input-example
var inputExample []byte

func main() {
	out1 := Task1(bytes.Fields(input))
	fmt.Println(106648 == out1, out1)

	const cycles = 1000000000

	// out2 := Task2(bytes.Fields(inputExample), cycles)
	// fmt.Println(64 == out2, out2)

	// out2 := Task2v2(bytes.Fields(input), cycles)
	// fmt.Println(87700 == out2)

	out2 := 0
	ops.TimeIt(func() {
		out2 = Task2v3(bytes.Fields(input))
	})
	fmt.Println(87700 == out2)
}

func printField(yMax, xMax int, roundRocks map[Rock]struct{}, squareRocks map[Rock]struct{}) {
	for y := 0; y < yMax; y++ {
		row := make([]byte, 0, xMax)
		for x := 0; x < xMax; x++ {
			if _, ok := squareRocks[Rock{Y: y, X: x}]; ok {
				row = append(row, '#')
				continue
			}
			if _, ok := roundRocks[Rock{Y: y, X: x}]; ok {
				row = append(row, 'O')
				continue
			}

			row = append(row, '.')
		}

		fmt.Println(string(row))
	}
	fmt.Println()
}
