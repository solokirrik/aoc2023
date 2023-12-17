package main

import (
	"bytes"
	"testing"
)

func TestTask1(t *testing.T) {
	in := [2]int{0, 0}
	tests := []struct {
		name string
		inp  []byte
		out  [2]int
		want int
	}{
		{"example", inputExample, in, 102},
		// {"input", input, in, 866},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mtx := bytes.Fields(tt.inp)
			rows := len(mtx)
			cols := len(mtx[0])

			if out1 := Task1(mtx, in, [2]int{rows - 1, cols - 1}); out1 != tt.want {
				t.Errorf("got=%d, want=%d", out1, tt.want)
			}
		})
	}
}
