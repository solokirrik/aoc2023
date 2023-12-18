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

type Instruction struct {
	Dir  string
	Dist int
}

type Point2D struct {
	x, y float64
}

func main() {
	inp := bytes.Split(input, []byte("\n"))

	// 200ms
	// Part1 V1
	ops.TimeIt(func() {
		out1 := Task1(inp)
		fmt.Println(76387 == out1, out1)
	})

	// 64.167µs
	// Part1 V2
	ops.TimeIt(func() {
		out1 := Task1V2(inp)
		fmt.Println(76387 == out1, out1)
	})

	// 86.833µs
	// Part2
	ops.TimeIt(func() {
		out23 := Task2V2(inp)
		fmt.Println(250022188522074 == uint64(out23), uint64(out23))
	})
}
