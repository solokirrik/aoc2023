package main

type Point struct {
	Y, X int
}

type DirectedPoint struct {
	Y, X, D int
}

const (
	UP    = 1
	DOWN  = 2
	LEFT  = 3
	RIGHT = 4
)

func dirToD(dir [2]int) int {
	switch {
	case dir[0] == 1 && dir[1] == 0:
		return DOWN
	case dir[0] == -1 && dir[1] == 0:
		return UP
	case dir[0] == 0 && dir[1] == 1:
		return RIGHT
	case dir[0] == 0 && dir[1] == -1:
		return LEFT
	}

	return 0
}

func Task1(mtx [][]byte, start Beam) int {
	maxRow := len(mtx) - 1
	maxCol := len(mtx[0]) - 1
	curBeam := start
	positionsQueue := []Beam{curBeam}
	energized := make(map[Point]struct{})
	visited := make(map[DirectedPoint]struct{})

	for len(positionsQueue) > 0 {
		curBeam = positionsQueue[0]
		energized[curBeam.Pos] = struct{}{}
		positionsQueue = positionsQueue[1:]

		mirror := mtx[curBeam.Pos.Y][curBeam.Pos.X]
		newBeams := moveBeams(maxRow, maxCol, curBeam.Apply(mirror))

		if len(newBeams) == 0 {
			continue
		}

		for i := range newBeams {
			key := DirectedPoint{
				Y: newBeams[i].Pos.Y,
				X: newBeams[i].Pos.X,
				D: dirToD(newBeams[i].Direction),
			}
			if _, ok := visited[key]; ok {
				continue
			}
			visited[key] = struct{}{}
			positionsQueue = append(positionsQueue, newBeams[i])
		}
	}

	return len(energized)
}

type Beam struct {
	Direction [2]int // Y, X sign
	Pos       Point
}

func (b *Beam) Sign() string {
	switch {
	case b.Direction[0] == 1 && b.Direction[1] == 0:
		return "v"
	case b.Direction[0] == -1 && b.Direction[1] == 0:
		return "^"
	case b.Direction[0] == 0 && b.Direction[1] == 1:
		return ">"
	case b.Direction[0] == 0 && b.Direction[1] == -1:
		return "<"
	}

	return "?"
}

func (b Beam) Apply(next byte) []Beam {
	to := b.Pos

	switch next {
	case '.':
		return []Beam{{Direction: b.Direction, Pos: to}}
	case '|':
		beamUp := Beam{Direction: [2]int{-1, 0}, Pos: to}
		beamDown := Beam{Direction: [2]int{+1, 0}, Pos: to}

		if b.Direction[1] == 0 { // X == 0
			if b.Direction[0] > 0 { // TOP to BOTTOM
				return []Beam{beamDown}
			}
			return []Beam{beamUp}
		}

		if b.Direction[0] == 0 { // Y == 0
			return []Beam{beamUp, beamDown}
		}
	case '-':
		beamLeft := Beam{Direction: [2]int{0, -1}, Pos: to}
		beamRight := Beam{Direction: [2]int{0, +1}, Pos: to}

		if b.Direction[0] == 0 { // Y == 0
			if b.Direction[1] > 0 { // LEFT to RIGHT
				return []Beam{beamRight}
			}
			return []Beam{beamLeft}
		}

		if b.Direction[1] == 0 { // X == 0
			return []Beam{beamLeft, beamRight}
		}
	case '/':
		newBeam := b
		newBeam.Pos = to

		// beam was vertical
		if b.Direction[0] != 0 {
			newBeam.Direction = [2]int{0, -1 * b.Direction[0]}
		}

		// beam was horizontal
		if b.Direction[1] != 0 {
			newBeam.Direction = [2]int{-1 * b.Direction[1], 0}
		}
		return []Beam{newBeam}
	case '\\':
		newBeam := b
		newBeam.Pos = to

		// beam was vertical
		if b.Direction[0] != 0 {
			newBeam.Direction = [2]int{0, 1 * b.Direction[0]}
		}

		// beam was horizontal
		if b.Direction[1] != 0 {
			newBeam.Direction = [2]int{1 * b.Direction[1], 0}
		}

		return []Beam{newBeam}
	}

	return nil
}

func moveBeams(maxRow, maxCol int, beams []Beam) []Beam {
	out := make([]Beam, 0, 1)

	for i := range beams {
		newY := beams[i].Pos.Y + beams[i].Direction[0]
		newX := beams[i].Pos.X + beams[i].Direction[1]

		if newY < 0 || newY > maxRow || newX < 0 || newX > maxCol {
			continue
		}
		beams[i].Pos = Point{Y: newY, X: newX}
		out = append(out, beams[i])
	}

	return beams
}
