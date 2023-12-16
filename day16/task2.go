package main

import (
	"sort"
	"sync"
)

func Task2(mtx [][]byte) int {
	outs := []int{}
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	rows := len(mtx)
	cols := len(mtx[0])

	wg.Add(2*rows + 2*cols)

	for i := 0; i < rows; i++ {
		go func(i int) {
			out := Task1(mtx, Beam{Direction: [2]int{0, 1}, Pos: Point{i, 0}})

			mu.Lock()
			outs = append(outs, out)
			mu.Unlock()

			wg.Done()
		}(i)

		go func(i int) {
			out := Task1(mtx, Beam{Direction: [2]int{0, -1}, Pos: Point{i, cols - 1}})

			mu.Lock()
			outs = append(outs, out)
			mu.Unlock()

			wg.Done()
		}(i)
	}

	for i := 0; i < cols; i++ {
		go func(i int) {
			out := Task1(mtx, Beam{Direction: [2]int{1, 0}, Pos: Point{0, i}})

			mu.Lock()
			outs = append(outs, out)
			mu.Unlock()

			wg.Done()
		}(i)

		go func(i int) {
			out := Task1(mtx, Beam{Direction: [2]int{-1, 0}, Pos: Point{rows - 1, i}})

			mu.Lock()
			outs = append(outs, out)
			mu.Unlock()

			wg.Done()
		}(i)
	}

	wg.Wait()

	sort.Ints(outs)

	return outs[len(outs)-1]
}
