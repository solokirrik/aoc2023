package main

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/solokirrik/aoc2023/ops"
)

//go:embed input
var input []byte

func main() {
	parts := bytes.Split(input, []byte(","))

	out1 := Task1(parts)
	fmt.Println(out1 == 509167, out1)

	ops.TimeIt(func() {
		out2 := Task2(parts)
		fmt.Println(out2 == 259333, out2)
	})
}

func Task1(parts [][]byte) int {
	out := 0
	for i := 0; i < len(parts); i++ {
		hashVal := hash(parts[i])
		out += hashVal
	}

	return out
}

func Task2(parts [][]byte) int {
	system := make(System, 256)

	for i := 0; i < len(parts); i++ {
		newLense := newLense(parts[i])
		system[newLense.hash].addLense(newLense)
	}

	return system.focusingPower()
}

func hash(label []byte) int {
	curVal := 0
	for i := 0; i < len(label); i++ {
		curVal += int(label[i])
		curVal *= 17
		curVal %= 256
	}

	return curVal
}
