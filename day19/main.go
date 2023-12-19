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
	ops.TimeIt(func() {
		out1 := Task1(bytes.Split(input, []byte("\n\n")))
		fmt.Println(out1 == 420739, out1)
	})

	ops.TimeIt(func() {
		out2 := Task2(bytes.Split(bytes.Split(input, []byte("\n\n"))[0], []byte("\n")))
		fmt.Println(out2 == 167409079868000, out2)
	})
}
