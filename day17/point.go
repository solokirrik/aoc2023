package main

import "strconv"

type Point [2]int // Y, X

func (p Point) String() string {
	return strconv.Itoa(p[0]) + "," + strconv.Itoa(p[1])
}

func (p Point) WithDirection(d int) DirectedPoint {
	return DirectedPoint{p[0], p[1], d}
}

type DirectedPoint [3]int // Y, X, D

func (p DirectedPoint) Extract() Point {
	return Point{p[0], p[1]}
}

func (p DirectedPoint) IsEqual(other Point) bool {
	return p[0] == other[0] && p[1] == other[1]
}

func (p DirectedPoint) possibleSteps() []Direction {
	switch p[2] {
	case UP:
		return []Direction{{0, 1}, {0, -1}, {-1, 0}}
	case DOWN:
		return []Direction{{0, 1}, {0, -1}, {1, 0}}
	case LEFT:
		return []Direction{{0, -1}, {1, 0}, {-1, 0}}
	case RIGHT:
		return []Direction{{0, 1}, {1, 0}, {-1, 0}}
	}

	return []Direction{}
}

func (p DirectedPoint) Move(step Direction) DirectedPoint {
	return DirectedPoint{p[0] + step[0], p[1] + step[1], step.encode()}
}

type Direction []int // dY, dX

func (d Direction) encode() int {
	switch {
	case d[0] == 1 && d[1] == 0:
		return DOWN
	case d[0] == -1 && d[1] == 0:
		return UP
	case d[0] == 0 && d[1] == 1:
		return RIGHT
	case d[0] == 0 && d[1] == -1:
		return LEFT
	}

	return 0
}

func getNeighbors(maxRow, maxCol int, curr DirectedPoint) []DirectedPoint {
	out := make([]DirectedPoint, 0, 4)
	possibleSteps := curr.possibleSteps()

	for i := range possibleSteps {
		step := possibleSteps[i]
		if curr[0]+step[0] < 0 || curr[0]+step[0] > maxRow ||
			curr[1]+step[1] < 0 || curr[1]+step[1] > maxCol {
			continue
		}

		out = append(out, curr.Move(step))
	}

	return out
}
