package main

import "fmt"

type point struct {
	y, x int
}

type Node struct {
	pos      point
	NL, NR   *Node
	pipeType string
	ends     []point
}

func (n Node) String() string {
	return fmt.Sprintf("%s[%d;%d]", string(n.pipeType), n.pos.y, n.pos.x)
}

func matching(n1, n2 Node) bool {
	if n1.pipeType == "." || n2.pipeType == "." {
		return false
	}

	if n1.pos.y == n2.pos.y && n1.pos.x == n2.pos.x {
		return false
	}

	if n1.pipeType == "S" && (samePoint(n1.pos, n2.ends[0]) || samePoint(n1.pos, n2.ends[1])) ||
		n2.pipeType == "S" && (samePoint(n2.pos, n1.ends[0]) || samePoint(n2.pos, n1.ends[1])) {
		return true
	}

	if n1.pipeType != "S" && n2.pipeType != "S" &&
		(samePoint(n1.ends[0], n2.pos) || samePoint(n1.ends[1], n2.pos)) &&
		(samePoint(n2.ends[0], n1.pos) || samePoint(n2.ends[1], n1.pos)) {
		return true
	}

	return false
}

func samePoint(p1, p2 point) bool {
	return p1.y == p2.y && p1.x == p2.x
}

func getNodeWithEnds(mtx [][]byte, r, c int) *Node {
	pipeType := string(mtx[r][c])
	maxR, maxCol := len(mtx)-1, len(mtx[0])-1

	if r < 0 || r > maxR || c < 0 || c > maxCol {
		return &Node{pipeType: "."}
	}

	switch mtx[r][c] {
	case '|': // is a vertical pipe connecting north and south.
		return &Node{
			pos:      point{y: r, x: c},
			pipeType: pipeType,
			ends:     []point{{y: max(r-1, 0), x: c}, {y: r + 1, x: c}},
		}
	case '-': // is a horizontal pipe connecting east and west.
		return &Node{
			pos:      point{y: r, x: c},
			pipeType: pipeType,
			ends:     []point{{y: r, x: max(c-1, 0)}, {y: r, x: min(c+1, maxCol)}},
		}
	case 'L': // is a 90-degree bend connecting north and east.
		return &Node{
			pos:      point{y: r, x: c},
			pipeType: pipeType,
			ends:     []point{{y: max(r-1, 0), x: c}, {y: r, x: min(c+1, maxCol)}},
		}
	case 'J': // is a 90-degree bend connecting north and west.
		return &Node{
			pos:      point{y: r, x: c},
			pipeType: pipeType,
			ends:     []point{{y: max(r-1, 0), x: c}, {y: r, x: max(c-1, 0)}},
		}
	case '7': // is a 90-degree bend connecting south and west.
		return &Node{
			pos:      point{y: r, x: c},
			pipeType: pipeType,
			ends:     []point{{y: r, x: max(c-1, 0)}, {y: min(r+1, maxR), x: c}},
		}
	case 'F': // is a 90-degree bend connecting south and east.
		return &Node{
			pos:      point{y: r, x: c},
			pipeType: pipeType,
			ends:     []point{{y: r, x: min(c+1, maxCol)}, {y: min(r+1, maxR), x: c}},
		}
	case '.': // is ground; there is no pipe in this tile.
		return &Node{pipeType: "."}
	case 'S': // is the starting position of the animal
		sPairs := startNeighbours(mtx, r, c)
		sNode := &Node{
			pos:      point{y: r, x: c},
			NL:       &sPairs[0],
			NR:       &sPairs[1],
			pipeType: pipeType,
			ends:     []point{sPairs[0].pos, sPairs[1].pos},
		}
		sPairs[0].NR = sNode
		sPairs[1].NL = sNode

		return sNode
	}

	return &Node{pipeType: "."}
}

func startNeighbours(mtx [][]byte, r, c int) []Node {
	out := make([]Node, 0, 2)
	maxR, maxCol := len(mtx)-1, len(mtx[0])-1

	cases := [][]int{
		{r, c - 1},
		{r, c + 1},
		{r - 1, c},
		{r + 1, c},
	}

	for _, cs := range cases {
		if cs[0] < 0 || cs[0] > maxR || cs[1] < 0 || cs[1] > maxCol {
			continue
		}

		if node := getNodeWithEnds(mtx, cs[0], cs[1]); node.pipeType != "." {
			sNode := Node{pos: point{y: r, x: c}, pipeType: "S"}
			if matching(sNode, *node) {
				out = append(out, *node)
			}
		}
	}

	return out
}
