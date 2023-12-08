package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input
var input []byte

func main() {
	out1 := Task1(input)
	fmt.Println(out1 == 11309, out1)

}

func Task1(rawLines []byte) int {
	lines := bytes.Split(rawLines, []byte("\n"))

	sequenceRaw := lines[0]
	mapRaw := lines[2:]
	networkMap := make(map[string][]string)

	for _, line := range mapRaw {
		parts := bytes.Split(line, []byte(" = "))
		root := string(parts[0])
		connections := bytes.Split(bytes.Trim(parts[1], "()"), []byte(", "))
		networkMap[root] = []string{
			string(connections[0]),
			string(connections[1]),
		}
	}

	i := 0
	curNode := "AAA"

	for curNode != "ZZZ" {
		i++
		direction := sequenceRaw[(i-1)%len(sequenceRaw)]
		curNode = networkMap[curNode][directionToIndex(direction)]
	}

	return i
}

func directionToIndex(direction byte) int {
	switch direction {
	case 'L':
		return 0
	case 'R':
		return 1
	}
	return -1
}
