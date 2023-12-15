package main

import (
	"bytes"
	_ "embed"
	"strconv"
	"testing"
)

//go:embed input-example
var inputExample []byte

func TestTask1(t *testing.T) {
	tests := []struct {
		name  string
		parts [][]byte
		want  int
	}{
		{
			name:  "example",
			parts: bytes.Split(inputExample, []byte(",")),
			want:  1320,
		},
		{
			name:  "input",
			parts: bytes.Split(input, []byte(",")),
			want:  509167,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Task1(tt.parts); got != tt.want {
				t.Errorf("got=%d, want=%d", got, tt.want)
			}
		})
	}
}

func TestTask2(t *testing.T) {
	tests := []struct {
		name  string
		parts [][]byte
		want  int
	}{
		{
			name:  "example",
			parts: bytes.Split(inputExample, []byte(",")),
			want:  145,
		},
		{
			name:  "input",
			parts: bytes.Split(input, []byte(",")),
			want:  259333,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Task2(tt.parts); got != tt.want {
				t.Errorf("got=%d, want=%d", got, tt.want)
			}
		})
	}
}

func TestBox(t *testing.T) {
	operations := [][]byte(
		bytes.Split([]byte("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"), []byte(",")),
	)

	tests := []struct {
		name string
		want string
	}{
		{
			name: "1",
			want: "rn.1|cm.2|ot.7|ab.5|pc.6|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := Box{}

			for i := range operations {
				newLense := newLense(operations[i])
				box.addLense(newLense)
			}

			out := serializeBox(&box)
			if out != tt.want {
				t.Errorf("\ngot=%s\nwant=%s", out, tt.want)
			}
		})
	}
}

func serializeBox(b *Box) string {
	out := ""
	curNode := b.head
	for curNode != nil {
		out += curNode.label + "." + strconv.Itoa(curNode.focalLength) + "|"
		curNode = curNode.next
	}

	return out
}
