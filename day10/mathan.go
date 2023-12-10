package main

import "math"

type Point2D struct {
	x, y float64
}

/*
Without Flood: 162.995708ms
With Flood: 40.265083ms
Per iteration:
- 2 sum
- 4 sub
- 1 mult
- 1 div
- 1 sign
*/
func pointInPolygonFast(target point, polygon []Node) bool {
	if len(polygon) < 3 {
		return false
	}

	qPatt := [][]int{{0, 1}, {3, 2}}
	end := len(polygon) - 1
	prevPoint := Point2D{float64(polygon[end].pos.x), float64(polygon[end].pos.y)}
	prevPoint.x -= float64(target.x)
	prevPoint.y -= float64(target.y)
	prevQ := qPatt[boolToInt(prevPoint.y < 0)][boolToInt(prevPoint.x < 0)]

	w := 0

	for i := 0; i <= end; i++ {
		curPoint := Point2D{float64(polygon[i].pos.x), float64(polygon[i].pos.y)}

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

/*
Without Flood: 1.306766458s
With Flood: 157.84875ms
Per iteration:
- 4 sum
- 2 sub
- 8 mult
- 2 div
- 2 arctg
*/
func pointInPolygonSlow(target point, polygon []Node) bool {
	if len(polygon) < 3 {
		return false
	}

	end := len(polygon) - 1
	prevPoint := Point2D{float64(polygon[end].pos.x), float64(polygon[end].pos.y)}
	prevPoint.x -= float64(target.x)
	prevPoint.y -= float64(target.y)

	sum := 0.0

	for i := 0; i <= end; i++ {
		curPoint := Point2D{float64(polygon[i].pos.x), float64(polygon[i].pos.y)}
		curPoint.x -= float64(target.x)
		curPoint.y -= float64(target.y)

		del := prevPoint.x*curPoint.y - curPoint.x*prevPoint.y
		xy := curPoint.x*prevPoint.x + curPoint.y*prevPoint.y

		sum += (math.Atan((prevPoint.x*prevPoint.x+prevPoint.y*prevPoint.y-xy)/del) +
			math.Atan((curPoint.x*curPoint.x+curPoint.y*curPoint.y-xy)/del))

		prevPoint = curPoint
	}

	return math.Abs(sum) > 5
}

func boolToInt(b bool) int {
	if b {
		return 1
	}

	return 0
}
