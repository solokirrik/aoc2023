package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input
var input []byte

func main() {
	inp := bytes.Split(input, []byte("\n"))

	out1 := Task1(inp, false)
	fmt.Println(out1 == 6733, out1)

	// 162.995708ms
	out21 := Task2NonFlood(inp, pointInPolygonFast, false)
	fmt.Println(out21 == 435, out21)

	// 40.265083ms
	out22 := Task2FloodFromEdge(inp, pointInPolygonFast, false)
	fmt.Println(out22 == 435, out22)
}
