package main

import "fmt"

func Task1(inp [][]byte, doPrint bool) int {
	loop, loopMap := extractLoop(inp)

	if doPrint {
		prettyPrint1(inp, loopMap)
	}

	return (len(loop) + 1) / 2
}

func extractLoop(inp [][]byte) (loop []Node, loopMap map[point]int) {
	start := findStart(inp)

	sNode := getNodeWithEnds(inp, start.y, start.x)
	loop = []Node{*sNode, *sNode.NR}
	loopMap = make(map[point]int)

	loopMap[start] = 0
	loopMap[sNode.NR.pos] = 1

	node := sNode.NR

	for {
		pos := findRightNode(node.NL.pos.y, node.NL.pos.x, node.ends)
		if pos == nil {
			fmt.Println("no right node", node)
			return loop, loopMap
		}

		newNode := getNodeWithEnds(inp, pos.y, pos.x)
		node.NR = newNode
		newNode.NL = node
		node = newNode

		if node.pipeType == "S" {
			break
		}

		loop = append(loop, *newNode)
		loopMap[newNode.pos] = len(loopMap)
	}

	return loop, loopMap
}

func findRightNode(nlPosR, nlPosC int, ends []point) *point {
	for i := range ends {
		if ends[i].y != nlPosR || ends[i].x != nlPosC {
			return &ends[i]
		}
	}

	return nil
}

func findStart(inp [][]byte) point {
	for r := range inp {
		for c := range inp[r] {
			if inp[r][c] == 'S' {
				return point{y: r, x: c}
			}
		}
	}

	return point{}
}
