package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
)

// heat death of the universe
func Task2(entries []Instruction) int {
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

func Task2V2(inp [][]byte) uint64 {
	instructions := make([]Instruction, 0, len(inp))
	for _, l := range inp {
		instructionRaw := bytes.Trim(bytes.Split(l, []byte(" "))[2], "()")
		last := len(instructionRaw) - 1

		instructions = append(instructions, Instruction{
			Dist: decodeHexIntoDec(instructionRaw[1:last]),
			Dir:  numToDir(byteToInt(instructionRaw[last])),
		})
	}

	return uint64(GaussArea(extractPoints(instructions)))
}

func GaussArea(points []Point2D) float64 {
	sum := 0.0
	last := len(points) - 1

	for i := 0; i < len(points)-1; i++ {
		sum += (points[i].y + points[i+1].y) * (points[i].x - points[i+1].x)
		sum += Len(points[i], points[i+1])
	}

	sum += (points[last].y + points[0].y) * (points[last].x - points[0].x)
	sum += Len(points[last], points[0])

	return sum/2 + 1
}

func Len(p1, p2 Point2D) float64 {
	if p1.y == p2.y {
		return math.Abs(p1.x - p2.x)
	}
	return math.Abs(p1.y - p2.y)
}

func decodeHexIntoDec(hex []byte) int {
	dec, err := strconv.ParseInt(string(hex), 16, 64)
	if err != nil {
		panic(err)
	}
	return int(dec)
}

func extractPoints(instr []Instruction) []Point2D {
	points := make([]Point2D, 0, len(instr))
	points = append(points, Point2D{0, 0})

	for p := range instr {
		lastPoint := points[len(points)-1]
		nextPoint := applyMove(lastPoint, instr[p].Dir, 1)
		points = append(points, nextPoint)
	}

	return points
}

func byteToInt(b byte) int {
	return int(b - '0')
}

func numToDir(n int) string {
	switch n {
	case 0:
		return "R"
	case 1:
		return "D"
	case 2:
		return "L"
	case 3:
		return "U"
	}

	panic("Unknown direction number" + strconv.Itoa(n))
}
