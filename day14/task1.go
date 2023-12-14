package main

import (
	"sync"
)

type Rock struct {
	Y, X         int
	alreadyMoved bool
}

type direction int

func (d direction) step() int {
	return int(d)
}

const (
	north direction = -1
	west  direction = -1
	south direction = 1
	east  direction = 1
)

func Task1(mtx [][]byte) int {
	roundRocks, _ := readRocks(mtx)

	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(len(mtx[0]))
	direction := north

	for x := range mtx[0] {
		// processColumn(direction, roundRocks, mtx, x, wg, mu)
		go processColumn(direction, roundRocks, mtx, x, wg, mu)
	}

	wg.Wait()

	return loadOnNorthBeam(mtx, roundRocks)
}

func loadOnNorthBeam(mtx [][]byte, roundRocks map[Rock]struct{}) int {
	out := 0

	lenMtx := len(mtx)
	for r := range roundRocks {
		out += lenMtx - r.Y
	}

	return out
}

func readRocks(mtx [][]byte) (map[Rock]struct{}, map[Rock]struct{}) {
	roundRocks := make(map[Rock]struct{})
	squareRocks := make(map[Rock]struct{})
	for y := range mtx {
		for x := range mtx[y] {
			if mtx[y][x] == '#' {
				squareRocks[Rock{Y: y, X: x}] = struct{}{}
			}
			if mtx[y][x] == 'O' {
				roundRocks[Rock{Y: y, X: x}] = struct{}{}
			}
		}
	}

	return roundRocks, squareRocks
}

func processColumn(
	dir direction,
	roundRocks map[Rock]struct{},
	mtx [][]byte,
	x int,
	wg *sync.WaitGroup,
	mu *sync.Mutex,
) {
	defer wg.Done()

	minY := 0
	maxY := len(mtx[0])
	checkY := func(y int) bool {
		return y < maxY
	}
	if dir.step() > 0 {
		minY = len(mtx[0]) - 1
		maxY = -1

		checkY = func(y int) bool {
			return y > maxY
		}
	}

	for y := minY; checkY(y); y -= dir.step() {
		if mtx[y][x] != 'O' {
			continue
		}
		if (y == 0 && dir.step() < 0) ||
			(y == len(mtx[0])-1 && dir.step() > 0) {
			continue
		}

		rockY := y

		for {
			rockY += dir.step()
			mu.Lock()
			_, isBusy := roundRocks[Rock{Y: rockY, X: x}]
			mu.Unlock()

			isSquareRock := mtx[rockY][x] == '#'
			isBorder := rockY == 0 || rockY == len(mtx)-1

			if isSquareRock || isBusy || isBorder {
				mu.Lock()
				delete(roundRocks, Rock{Y: y, X: x})
				mu.Unlock()

				if isSquareRock || isBusy {
					mu.Lock()
					roundRocks[Rock{Y: rockY - 1*dir.step(), X: x}] = struct{}{}
					mu.Unlock()
					break
				}

				if isBorder && !isBusy {
					mu.Lock()
					roundRocks[Rock{Y: rockY, X: x}] = struct{}{}
					mu.Unlock()
					break
				}
			}
		}
	}
}
