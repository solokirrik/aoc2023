package main

import (
	"fmt"
)

const (
	UP    = 1
	DOWN  = 2
	LEFT  = 3
	RIGHT = 4
)

type BFS struct {
	field   [][]byte
	visited map[string]int
	queue   Queue
	start   Point
	end     Point
}

func NewBFS(mtx [][]byte, start, end Point) *BFS {
	return &BFS{
		field:   mtx,
		start:   start,
		end:     end,
		visited: make(map[string]int),
		queue:   Queue{},
	}
}

func (b *BFS) lossAt(y, x int) int {
	return int(b.field[y][x] - '0')
}

func (b *BFS) Run() int {
	connections := make(map[Point]Point)

	b.queue.Push([]Waiter{
		{pos: b.start.WithDirection(RIGHT)},
		{pos: b.start.WithDirection(DOWN)},
	})

	for !b.queue.IsEmpty() {
		candidate := b.queue.Pop()
		pos := candidate.pos.Extract()
		dir := candidate.pos[2]

		heat := b.lossAt(pos[0], pos[1]) + candidate.heatLoss
		if candidate.pos.IsEqual(b.end) {
			b.visited[pos.String()] = heat
			connections[pos] = candidate.parent
			break
		}

		if cacheHeat, ok := b.visited[pos.String()]; ok {
			if cacheHeat < heat {
				continue
			}
		}

		b.visited[pos.String()] = heat
		connections[pos] = candidate.parent

		waiters := b.sheduleOptions(candidate, heat, dir)
		b.queue.Push(waiters)
	}

	printField(b.field, connections, b.end)

	return b.visited[b.end.String()]
}

func sign(a, b Point) byte {
	if a[0] == b[0] {
		if a[1] < b[1] {
			return '>'
		}
		return '<'
	}
	if a[0] < b[0] {
		return 'v'
	}
	return '^'
}

func (b *BFS) sheduleOptions(parent Waiter, heat, dir int) []Waiter {
	maxStraight := 3
	maxRow := len(b.field) - 1
	maxCol := len(b.field[0]) - 1

	options := getNeighbors(maxRow, maxCol, parent.pos)
	waiters := make([]Waiter, 0, len(options))
	allowStraight := parent.straightNumber < maxStraight

	for i := range options {
		if !allowStraight && dir == options[i][2] {
			continue
		}

		newWaiter := Waiter{
			parent:         parent.pos.Extract(),
			pos:            options[i],
			heatLoss:       heat,
			straightNumber: 1,
		}
		if dir == options[i][2] {
			newWaiter.straightNumber = parent.straightNumber + 1
		}
		waiters = append(waiters, newWaiter)
	}

	return waiters
}

func printField(field [][]byte, connections map[Point]Point, end Point) {
	path := []Point{}
	index := make(map[Point]Point)
	dest := end
	for {
		path = append(path, dest)
		orig, ok := connections[dest]
		if ok {
			index[dest] = orig
		}
		if !ok || (orig[0] == 0 && orig[1] == 0) {
			fmt.Println("not found", dest.String())
			break
		}
		dest = orig
	}

	newField := field[:]
	for r := range newField {
		for c := range newField[r] {
			origin, ok := index[Point{r, c}]
			if !ok {
				newField[r][c] = '.'
				continue
			}
			newField[r][c] = sign(origin, Point{r, c})
		}
		row := string(newField[r])
		fmt.Println(row)
	}
}
