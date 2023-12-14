package main

import (
	"bytes"
	"testing"
)

func TestTask2v2(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		want   int
		cycles int
	}{
		{
			name:   "example-test",
			input:  inputExample,
			want:   87,
			cycles: 1,
		},
		{
			name:   "example-simple",
			input:  inputExample,
			want:   69,
			cycles: 3,
		},
		// {
		// 	name:   "example-full",
		// 	input:  inputExample,
		// 	want:   64,
		// 	cycles: 1000000000,
		// },
		// {
		// 	name:   "task2-full",
		// 	input:  input,
		// 	want:   106648,
		// 	cycles: 1000000000,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := Task2v2(bytes.Fields(tt.input), tt.cycles)
			if out != tt.want {
				t.Errorf("Expected %d, got %d", tt.want, out)
			}
		})
	}
}

func TestGetNextVertical(t *testing.T) {
	colIdx := 4
	firstRockInRow := Rock{Y: 1, X: 4}
	defaultRock := Rock{Y: 3, X: 4}
	middleRockInRow := Rock{Y: 5, X: 4}
	lastRockInRow := Rock{Y: 9, X: 4}
	defaultIndex := map[int][]*Rock{
		4: {
			&firstRockInRow,
			&defaultRock,
			&middleRockInRow,
			&lastRockInRow,
		},
	}

	tests := []struct {
		name         string
		want         *Rock
		rock         *Rock
		rocksIndex   map[int][]*Rock
		isNextLowest bool
	}{
		{
			name:         "next-lowest",
			rock:         &defaultRock,
			want:         &firstRockInRow,
			rocksIndex:   defaultIndex,
			isNextLowest: true,
		},
		{
			name:         "next-highiest",
			rock:         &defaultRock,
			want:         &middleRockInRow,
			rocksIndex:   defaultIndex,
			isNextLowest: false,
		},
		{
			name:         "next-highiest-last-in-row",
			rock:         &lastRockInRow,
			want:         nil,
			rocksIndex:   defaultIndex,
			isNextLowest: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			next := getNextVertical(colIdx, tt.rock, tt.rocksIndex, tt.isNextLowest)
			if tt.want != nil && next == nil {
				t.Fatalf("Expected next to be not nil")
			}
			if tt.want == nil && next != nil {
				t.Fatalf("Expected next to be nil, got %d", next.Y)
			}
			if tt.want != nil && next != nil && next.Y != tt.want.Y {
				t.Fatalf("Expected %d, got %d", tt.want.Y, next.Y)
			}
		})
	}
}

func TestGetNextHorizontal(t *testing.T) {
	colIdx := 4
	rock := Rock{Y: 4, X: 3}
	columnsSquareRocks := map[int][]*Rock{
		4: {
			&Rock{Y: 4, X: 2},
			&Rock{Y: 4, X: 5},
			&Rock{Y: 4, X: 9},
		},
	}

	tests := []struct {
		name         string
		wantX        int
		isNextLowest bool
	}{
		{
			name:         "next-lowest",
			wantX:        2,
			isNextLowest: true,
		},
		{
			name:         "next-highiest",
			wantX:        5,
			isNextLowest: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			next := getNextHorizontal(colIdx, &rock, columnsSquareRocks, tt.isNextLowest)
			if next == nil {
				t.Fatalf("Expected next to be not nil")
			}
			if next.X != tt.wantX {
				t.Fatalf("Expected %d, got %d", tt.wantX, next.X)
			}
		})
	}
}
