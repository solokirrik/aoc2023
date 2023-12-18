package main

import (
	"bytes"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

func Task1(inp [][]byte) int {
	entries := make([]Instruction, 0, len(inp))
	for _, l := range inp {
		parts := bytes.Split(l, []byte(" "))
		dist := int(parts[1][0] - '0')
		if len(parts[1]) > 1 {
			dist = int(parts[1][0]-'0')*10 + int(parts[1][1]-'0')
		}
		e := Instruction{
			Dir:  string(parts[0]),
			Dist: dist,
		}

		entries = append(entries, e)
	}

	perimetr, perimIndex, topLeft, downRight := buildPerimeter(entries)
	fmt.Println(topLeft, downRight)

	cubics := atomic.Uint64{}
	wg := sync.WaitGroup{}

	for y := topLeft.y; y < downRight.y+1; y++ {
		for x := topLeft.x; x < downRight.x+1; x++ {
			wg.Add(1)
			go func(x, y float64, wg *sync.WaitGroup) {
				defer wg.Done()
				if _, ok := perimIndex[Point2D{x, y}]; ok {
					cubics.Add(1)
					return
				}
				if PointInPolygonFast(Point2D{x, y}, perimetr) {
					cubics.Add(1)
				}
			}(x, y, &wg)
		}
	}

	wg.Wait()

	return int(cubics.Load())
}

func Task1V2(inp [][]byte) uint64 {
	instructions := make([]Instruction, 0, len(inp))
	for _, l := range inp {
		parts := bytes.Split(l, []byte(" "))
		dist := int(parts[1][0] - '0')
		if len(parts[1]) > 1 {
			dist = int(parts[1][0]-'0')*10 + int(parts[1][1]-'0')
		}

		instructions = append(instructions, Instruction{
			Dir:  string(parts[0]),
			Dist: dist,
		})
	}

	return uint64(GaussArea(extractPoints(instructions)))
}

func buildPerimeter(instr []Instruction) ([]Point2D, map[Point2D]struct{}, Point2D, Point2D) {
	perimeter := []Point2D{{0, 0}}
	perimIndex := map[Point2D]struct{}{{0, 0}: {}}

	topLeft, downRight := Point2D{math.Inf(1), math.Inf(1)}, Point2D{math.Inf(-1), math.Inf(-1)}

	for p := range instr {
		for i := 0; i < instr[p].Dist; i++ {
			lastPoint := perimeter[len(perimeter)-1]
			nextPoint := applyMove(lastPoint, instr[p].Dir, 1)
			perimIndex[nextPoint] = struct{}{}
			if nextPoint.x > downRight.x {
				downRight.x = nextPoint.x
			}
			if nextPoint.y > downRight.y {
				downRight.y = nextPoint.y
			}
			if nextPoint.x < topLeft.x {
				topLeft.x = nextPoint.x
			}
			if nextPoint.y < topLeft.y {
				topLeft.y = nextPoint.y
			}
			perimeter = append(perimeter, nextPoint)
		}
	}

	return perimeter, perimIndex, topLeft, downRight
}

func PointInPolygonFast(target Point2D, polygon []Point2D) bool {
	if len(polygon) < 3 {
		return false
	}

	qPatt := [][]int{{0, 1}, {3, 2}}
	end := len(polygon) - 1
	prevPoint := Point2D{float64(polygon[end].x), float64(polygon[end].y)}
	prevPoint.x -= float64(target.x)
	prevPoint.y -= float64(target.y)
	prevQ := qPatt[boolToInt(prevPoint.y < 0)][boolToInt(prevPoint.x < 0)]

	w := 0

	for i := 0; i <= end; i++ {
		curPoint := polygon[i]

		curPoint.x -= float64(target.x)
		curPoint.y -= float64(target.y)

		q := qPatt[boolToInt(curPoint.y < 0)][boolToInt(curPoint.x < 0)]

		switch q - prevQ {
		case -3:
			w++
		case 3:
			w--
		case -2:
			if prevPoint.x*curPoint.y >= prevPoint.y*curPoint.x {
				w++
			}
		case 2:
			if !(prevPoint.x*curPoint.y >= prevPoint.y*curPoint.x) {
				w--
			}
		}

		prevPoint = curPoint
		prevQ = q
	}

	return w != 0
}

func boolToInt(b bool) int {
	if b {
		return 1
	}

	return 0
}

func applyMove(p Point2D, dir string, step float64) Point2D {
	switch dir {
	case "R":
		return Point2D{p.x + step, p.y}
	case "L":
		return Point2D{p.x - step, p.y}
	case "D":
		return Point2D{p.x, p.y + step}
	case "U":
		return Point2D{p.x, p.y - step}
	}

	panic("Unknown direction")
}
