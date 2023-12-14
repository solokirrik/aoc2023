package main

import (
	"bytes"
)

type Platform [][]byte
type Direction uint8

const (
	SQUARE = '#'
	ROUND  = 'O'
	EMPTY  = '.'

	NORTH Direction = 1
	WEST  Direction = 2
	SOUTH Direction = 3
	EAST  Direction = 4
)

func Task2v3(input [][]byte) int {
	p := Platform(input)

	cache := map[string]int{}
	cycles := 1000000000

	for i := 0; i < cycles; i++ {
		p.Tilt(NORTH)
		p.Tilt(WEST)
		p.Tilt(SOUTH)
		p.Tilt(EAST)
		h := bytes.Join(p, []byte{})
		if _, ok := cache[string(h)]; ok {
			i = cycles - (cycles-i)%(i-cache[string(h)])
		}
		cache[string(h)] = i
	}

	return LoadOnNorthBeamV3(p)
}

func (p Platform) Tilt(d Direction) {
	switch d {
	case NORTH:
		for x := 0; x < len(p[0]); x++ {
			freeRow := 0
			for y := 0; y < len(p); y++ {
				switch p[y][x] {
				case SQUARE:
					freeRow = y + 1
				case ROUND:
					if freeRow < y {
						p[freeRow][x] = ROUND
						p[y][x] = EMPTY
					}
					freeRow++
				}
			}
		}
	case WEST:
		for y := 0; y < len(p); y++ {
			freeCol := 0
			for x := 0; x < len(p[y]); x++ {
				switch p[y][x] {
				case SQUARE:
					freeCol = x + 1
				case ROUND:
					if freeCol < x {
						p[y][freeCol] = ROUND
						p[y][x] = EMPTY
					}
					freeCol++
				}
			}
		}
	case SOUTH:
		for x := 0; x < len(p[0]); x++ {
			freeRow := len(p) - 1
			for y := len(p) - 1; y >= 0; y-- {
				switch p[y][x] {
				case SQUARE:
					freeRow = y - 1
				case ROUND:
					if freeRow > y {
						p[freeRow][x] = ROUND
						p[y][x] = EMPTY
					}
					freeRow--
				}
			}
		}
	case EAST:
		for y := 0; y < len(p); y++ {
			freeCol := len(p[y]) - 1
			for x := len(p[y]) - 1; x >= 0; x-- {
				switch p[y][x] {
				case SQUARE:
					freeCol = x - 1
				case ROUND:
					if freeCol > x {
						p[y][freeCol] = ROUND
						p[y][x] = EMPTY
					}
					freeCol--
				}
			}
		}
	}
}

func LoadOnNorthBeamV3(p Platform) int {
	out := 0

	rows := len(p)
	cols := len(p[0])
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if p[y][x] == ROUND {
				out += rows - y
			}
		}
	}

	return out
}
