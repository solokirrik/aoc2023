package main

func Task2NonFlood(inp [][]byte, pointInPolygon func(target point, polygon []Node) bool, doPrint bool) int {
	loop, loopMap := extractLoop(inp)

	notInContour := make(map[point]int)
	for r := range inp {
		for c := range inp[r] {
			if _, ok := loopMap[point{y: r, x: c}]; ok {
				continue
			}

			if pointInPolygon(point{y: r, x: c}, loop) {
				notInContour[point{y: r, x: c}] = len(notInContour)
			}
		}
	}

	if doPrint {
		prettyPrint2(inp, loopMap, notInContour)
	}

	return len(notInContour)
}

func Task2FloodFromEdge(inp [][]byte, pointInPolygon func(target point, polygon []Node) bool, doPrint bool) int {
	loop, loopMap := extractLoop(inp)

	edges := make(map[point]int)
	rMax := len(inp) - 1
	cMax := len(inp[0]) - 1

	for r := range inp {
		for c := range inp[r] {
			flood(loopMap, edges, r, c, rMax, cMax)
		}
	}

	for r := rMax; r >= 0; r-- {
		for c := cMax; c >= 0; c-- {
			flood(loopMap, edges, r, c, rMax, cMax)
		}
	}

	for r := rMax; r >= 0; r-- {
		for c := range inp[r] {
			flood(loopMap, edges, r, c, rMax, cMax)
		}
	}

	for r := range inp {
		for c := cMax; c >= 0; c-- {
			flood(loopMap, edges, r, c, rMax, cMax)
		}
	}

	notInContour := make(map[point]int)
	for r := range inp {
		for c := range inp[r] {
			if _, ok := loopMap[point{y: r, x: c}]; ok {
				continue
			}

			if _, ok := edges[point{y: r, x: c}]; ok {
				continue
			}

			isInPol := pointInPolygon(point{y: r, x: c}, loop)
			if isInPol {
				notInContour[point{y: r, x: c}] = len(notInContour)
			}
		}
	}

	if doPrint {
		prettyPrint2(inp, loopMap, notInContour)
	}

	return len(notInContour)
}

func flood(loopMap, edges map[point]int, r, c, rMax, cMax int) {
	if _, ok := loopMap[point{y: r, x: c}]; ok {
		return
	}

	if r == 0 || r == rMax || // top or bottom
		c == 0 || c == cMax { // left or right
		if _, ok := loopMap[point{y: r, x: c}]; !ok {
			edges[point{y: r, x: c}] = len(edges)
		}
		return
	}

	if gotEdgeNeighbour(edges, r, c) {
		edges[point{y: r, x: c}] = len(edges)
	}
}

func gotEdgeNeighbour(edges map[point]int, r, c int) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if _, ok := edges[point{y: r + i, x: c + j}]; ok {
				return true
			}
		}
	}

	return false
}
