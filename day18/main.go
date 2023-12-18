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
	// Part1 V1
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

	// 64.167µs
	// Part1 V2
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

		out1 := GaussArea(EntriesToLines(entries))
		fmt.Println(76387 == out1, out1)
	})

	// 93.917µs
	// Part2
	ops.TimeIt(func() {
		inp := bytes.Split(input, []byte("\n"))
		entries := make([]Entry, 0, len(inp))
		for _, l := range inp {
			part2Info := bytes.Trim(bytes.Split(l, []byte(" "))[2], "()")
			e := Entry{
				Dist: decodeHexIntoDec(string(part2Info[1 : len(part2Info)-1])),
				Dir:  numToDir(int(part2Info[len(part2Info)-1]) - '0'),
			}
			entries = append(entries, e)
		}

		out22 := GaussArea(EntriesToLines(entries))
		fmt.Println(250022188522074 == uint64(out22), uint64(out22))
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

// heat death of the universe
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

func EntriesToLines(entries []Entry) []Line {
	lines := make([]Line, 0, len(entries))
	p1x, p1y, p2x, p2y := 0.0, 0.0, 0.0, 0.0

	for _, e := range entries {
		switch e.Dir {
		case "R":
			p2x = p1x + float64(e.Dist)
		case "L":
			p2x = p1x - float64(e.Dist)
		case "D":
			p2y = p1y + float64(e.Dist)
		case "U":
			p2y = p1y - float64(e.Dist)
		}

		lines = append(lines, Line{
			p1: Point2D{p1x, p1y},
			p2: Point2D{p2x, p2y},
		})

		p1x, p1y = p2x, p2y
	}

	return lines
}

type Line struct {
	p1 Point2D
	p2 Point2D
}

func (l Line) Len() float64 {
	if l.p1.y == l.p2.y {
		return math.Abs(l.p1.x - l.p2.x)
	}
	return math.Abs(l.p1.y - l.p2.y)
}

func GaussArea(lines []Line) float64 {
	sum := 0.0

	for i := 0; i < len(lines); i++ {
		sum += (lines[i].p1.y + lines[i].p2.y) * (lines[i].p1.x - lines[i].p2.x)
		sum += lines[i].Len()
	}

	return sum/2 + 1
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
			nextPoint := applyDirection(lastPoint, entries[p].Dir, 1)
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

func buildPerimeterV2(entries []Entry) ([]Point2D, map[Point2D]struct{}, Point2D, Point2D) {
	perimeter := []Point2D{{0, 0}}
	perimIndex := map[Point2D]struct{}{{0, 0}: {}}

	topLeft, downRight := Point2D{math.Inf(1), math.Inf(1)}, Point2D{math.Inf(-1), math.Inf(-1)}

	for p := range entries {
		lastPoint := perimeter[len(perimeter)-1]
		nextPoint := applyDirection(lastPoint, entries[p].Dir, float64(entries[p].Dist))
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

func applyDirection(p Point2D, d string, step float64) Point2D {
	if d == "R" {
		return Point2D{p.x + step, p.y}
	}
	if d == "L" {
		return Point2D{p.x - step, p.y}
	}
	if d == "D" {
		return Point2D{p.x, p.y + step}
	}
	if d == "U" {
		return Point2D{p.x, p.y - step}
	}

	panic("Unknown direction")
}
