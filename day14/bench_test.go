package main

import (
	"bytes"
	"sync"
	"testing"
)

func BenchmarkConcurTask1(b *testing.B) {
	mtx := bytes.Split(input, []byte("\n"))
	roundRocks, _ := readRocks(mtx)

	b.Run("Concur", func(b *testing.B) {
		mu := &sync.Mutex{}
		wg := &sync.WaitGroup{}
		wg.Add(len(mtx[0]))
		direction := north

		for x := range mtx[0] {
			go processColumn(direction, roundRocks, mtx, x, wg, mu)
		}

		wg.Wait()
	})

	b.Run("NonConcur", func(b *testing.B) {
		mu := &sync.Mutex{}
		wg := &sync.WaitGroup{}
		wg.Add(len(mtx[0]))
		direction := north

		for x := range mtx[0] {
			processColumn(direction, roundRocks, mtx, x, wg, mu)
		}

		wg.Wait()
	})
}
