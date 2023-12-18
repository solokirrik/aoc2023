package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/solokirrik/aoc2023/ops"
)

//go:embed input
var input []byte

//go:embed input-example
var inputExample []byte

func main() {
	// 200ms
	ops.TimeIt(func() {
		inp := bytes.Split(input, []byte("\n"))
		entries := make([]Entry, 0, len(inp))
		for _, l := range inp {
			parts := bytes.Split(l, []byte(" "))
			dist := int(parts[1][0] - '0')
			if len(parts[1]) > 1 {
				dist = int(parts[1][0]-'0')*10 + int(parts[1][1]-'0')
			}
			e := Entry{
				Dir:  string(parts[0]),
				Dist: dist,
			}

			entries = append(entries, e)
		}

		out1 := Task1(entries)
		fmt.Println(76387 == out1, out1)
	})

	// heat death of the universe
	ops.TimeIt(func() {
		inp := bytes.Split(inputExample, []byte("\n"))
		entries := make([]Entry, 0, len(inp))
		for _, l := range inp {
			part2Info := bytes.Trim(bytes.Split(l, []byte(" "))[2], "()")
			e := Entry{
				Dist: decodeHexIntoDec(string(part2Info[1 : len(part2Info)-1])),
				Dir:  numToDir(int(part2Info[len(part2Info)-1]) - '0'),
			}
			entries = append(entries, e)
		}

		out2 := Task1(entries)
		fmt.Println(952408144115 == out2, out2)
	})
}

type Entry struct {
	Dir  string
	Dist int
}

func Task1(entries []Entry) int {
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

func Task2(entries []Entry) int {
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

func decodeHexIntoDec(hex string) int {
	dec, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		panic(err)
	}
	return int(dec)
}

func buildPerimeter(entries []Entry) ([]Point2D, map[Point2D]struct{}, Point2D, Point2D) {
	perimeter := []Point2D{{0, 0}}
	perimIndex := map[Point2D]struct{}{{0, 0}: {}}

	topLeft, downRight := Point2D{math.Inf(1), math.Inf(1)}, Point2D{math.Inf(-1), math.Inf(-1)}

	for p := range entries {
		for i := 0; i < entries[p].Dist; i++ {
			lastPoint := perimeter[len(perimeter)-1]
			nextPoint := applyDirection(lastPoint, entries[p].Dir, entries[p].Dist)
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

type Point2D struct {
	x, y float64
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

func numToDir(n int) string {
	if n == 0 {
		return "R"
	}
	if n == 1 {
		return "D"
	}
	if n == 2 {
		return "L"
	}
	if n == 3 {
		return "U"
	}

	panic("Unknown direction number" + strconv.Itoa(n))
}

func applyDirection(p Point2D, d string, step int) Point2D {
	if d == "R" {
		return Point2D{p.x + 1, p.y}
	}
	if d == "L" {
		return Point2D{p.x - 1, p.y}
	}
	if d == "D" {
		return Point2D{p.x, p.y + 1}
	}
	if d == "U" {
		return Point2D{p.x, p.y - 1}
	}

	panic("Unknown direction")
}
