package main

import (
	"bytes"
	_ "embed"
)

//go:embed input
var input []byte

func main() {
	squares := bytes.Split(input, []byte("\n\n"))

	out1 := Task1(squares)
	println(out1 == 28651, out1)
}
