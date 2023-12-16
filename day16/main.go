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
	out1 := 0
	ops.TimeIt(
		func() {
			out1 = Task1(bytes.Fields(input), Beam{Direction: [2]int{0, 1}, Pos: Point{0, 0}})
		})
	fmt.Println(out1 == 6740, out1)

	out2 := 0
	ops.TimeIt(
		func() {
			out2 = Task2(bytes.Fields(input))
		})
	fmt.Println(out2 == 7041, out2)

}
