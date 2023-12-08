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

	out2 := Task2(input)
	fmt.Println(out2 == 13740108158591, out2)
}

func parse(rawLines []byte) (sequenceRaw []byte, networkMap map[string][]string) {
	lines := bytes.Split(rawLines, []byte("\n"))

	sequenceRaw = lines[0]
	mapRaw := lines[2:]
	networkMap = make(map[string][]string)

	for _, line := range mapRaw {
		if len(line) == 0 {
			continue
		}
		parts := bytes.Split(line, []byte(" = "))
		root := string(parts[0])
		connections := bytes.Split(bytes.Trim(parts[1], "()"), []byte(", "))
		networkMap[root] = []string{
			string(connections[0]),
			string(connections[1]),
		}
	}

	return sequenceRaw, networkMap
}

type State struct {
	Position     string
	StepsToReach int
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
