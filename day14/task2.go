package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func Task2(mtx [][]byte, cycles int) int {
	start := time.Now()
	roundRocks, _ := readRocks(mtx)
	for i := 0; i < cycles; i++ {
		if i%100000 == 0 {
			since := time.Since(start).String()
			fmt.Println("%", math.Round(float64(i)/float64(cycles)*100*10000)/10000,
				"| time", since[0:min(10, len(since))],
				"| iteration", i, "of", cycles)
		}

		cycleVerticalDirection(north, mtx, roundRocks)
		cycleHorizontalDirection(west, mtx, roundRocks)
		cycleVerticalDirection(south, mtx, roundRocks)
		cycleHorizontalDirection(east, mtx, roundRocks)
	}

	return loadOnNorthBeam(mtx, roundRocks)
}

func cycleVerticalDirection(dir direction, mtx [][]byte, roundRocks map[Rock]struct{}) {
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(len(mtx[0]))

	for x := range mtx[0] {
		go processColumn(dir, roundRocks, mtx, x, wg, mu)
	}

	wg.Wait()
}

func cycleHorizontalDirection(dir direction, mtx [][]byte, roundRocks map[Rock]struct{}) {
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(len(mtx))

	for y := range mtx {
		go processRow(dir, roundRocks, mtx, y, wg, mu)
	}

	wg.Wait()
}

func processRow(
	dir direction,
	roundRocks map[Rock]struct{},
	mtx [][]byte,
	y int,
	wg *sync.WaitGroup,
	mu *sync.Mutex,
) {
	defer wg.Done()

	minX := 0
	maxX := len(mtx[0])
	checkX := func(x int) bool {
		return x < maxX
	}
	if dir.step() > 0 {
		minX = len(mtx[0]) - 1
		maxX = -1

		checkX = func(x int) bool {
			return x > maxX
		}
	}

	for x := minX; checkX(x); x -= dir.step() {
		if mtx[y][x] != 'O' {
			continue
		}

		if (x == 0 && dir.step() < 0) ||
			(x == len(mtx[0])-1 && dir.step() > 0) {
			continue
		}

		rockX := x

		for {
			rockX += dir.step()

			mu.Lock()
			_, isBusy := roundRocks[Rock{Y: y, X: rockX}]
			mu.Unlock()

			isSquareRock := mtx[y][rockX] == '#'
			isBorder := rockX == 0 || rockX == len(mtx[0])-1

			if isSquareRock || isBusy || isBorder {
				mu.Lock()
				delete(roundRocks, Rock{Y: y, X: x})
				mu.Unlock()

				if isSquareRock || isBusy {
					mu.Lock()
					roundRocks[Rock{Y: y, X: rockX - 1*dir.step()}] = struct{}{}
					mu.Unlock()
					break
				}

				if isBorder && !isBusy {
					mu.Lock()
					roundRocks[Rock{Y: y, X: rockX}] = struct{}{}
					mu.Unlock()
					break
				}
			}
		}
	}
}
