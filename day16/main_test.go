package main

// import (
// 	"bytes"
// 	"sort"
// 	"testing"
// )

// func TestTask1(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		inp  [][]byte
// 		want int
// 	}{
// 		{
// 			name: "example",
// 			inp:  bytes.Fields(inputExample),
// 			want: 46,
// 		},
// 		{
// 			name: "input",
// 			inp:  bytes.Fields(input),
// 			want: 6740,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Task1(tt.inp, Beam{Direction: [2]int{0, 1}, Pos: Point{0, 0}}); got != tt.want {
// 				t.Errorf("got=%d, want=%d", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestTask2(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		inp  [][]byte
// 		want int
// 	}{
// 		{
// 			name: "example",
// 			inp:  bytes.Fields(inputExample),
// 			want: 51,
// 		},
// 		{
// 			name: "input",
// 			inp:  bytes.Fields(input),
// 			want: 7041,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Task2(tt.inp); got != tt.want {
// 				t.Errorf("got=%d, want=%d", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestToBeams(t *testing.T) {
// 	t.Run("keep-direction", func(t *testing.T) {
// 		beam := Beam{Direction: [2]int{0, 1}, Pos: Point{0, 0}}
// 		positions := filterBeams(10, 10, moveBeams(Apply('.', beam, Point{0, 1})))
// 		if len(positions) != 1 {
// 			t.Errorf("got=%d, want=%d", len(positions), 1)
// 		}
// 	})

// 	t.Run("up-and-down", func(t *testing.T) {
// 		beam := Beam{Direction: [2]int{0, 1}, Pos: Point{0, 0}}
// 		positions := filterBeams(10, 10, moveBeams(Apply('|', beam, Point{2, 5})))
// 		if len(positions) != 2 {
// 			t.Errorf("got=%d, want=%d", len(positions), 2)
// 		}

// 		sort.Slice(positions, func(i, j int) bool {
// 			return positions[i].Pos.Y < positions[j].Pos.Y
// 		})

// 		if positions[0].Pos != (Point{1, 5}) {
// 			t.Errorf("got=%v, want=%v", positions[0], Point{1, 5})
// 		}
// 		if positions[1].Pos != (Point{3, 5}) {
// 			t.Errorf("got=%v, want=%v", positions[1], Point{3, 5})
// 		}
// 	})

// 	t.Run("right", func(t *testing.T) {
// 		beam := Beam{Direction: [2]int{0, 1}, Pos: Point{2, 5}}
// 		positions := filterBeams(10, 10, moveBeams(Apply('-', beam, Point{2, 5})))
// 		if len(positions) != 1 {
// 			t.Errorf("got=%d, want=%d", len(positions), 1)
// 		}
// 		if positions[0].Pos.Y != 2 {
// 			t.Errorf("got=%d, want=%d", positions[0].Pos.Y, 2)
// 		}
// 		if positions[0].Pos.X != 6 {
// 			t.Errorf("got=%d, want=%d", positions[0].Pos.X, 6)
// 		}
// 	})
// }
