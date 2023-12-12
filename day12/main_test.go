package main

import (
	"bytes"
	_ "embed"
	"testing"
)

//go:embed input-example
var inputExample []byte

func TestTask1(t *testing.T) {
	cases := []struct {
		name string
		in   []byte
		want int
	}{
		{
			name: "example",
			in:   inputExample,
			want: 21,
		},
		{
			name: "input",
			in:   input,
			want: 7169,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			lines := bytes.Split(c.in, []byte("\n"))

			var cache = make(map[string]int)

			out := Task1(cache, lines)
			if out != c.want {
				t.Errorf("Expected %v, got %v", c.want, out)
			}
		})
	}
}

func TestTask2(t *testing.T) {
	cases := []struct {
		name string
		in   []byte
		want int
	}{
		{
			name: "example",
			in:   inputExample,
			want: 525152,
		},
		{
			name: "input",
			in:   input,
			want: 1738259948652,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			lines := bytes.Split(c.in, []byte("\n"))

			var cache = make(map[string]int)

			out := Task2(cache, lines)
			if out != c.want {
				t.Errorf("Expected %v, got %v", c.want, out)
			}
		})
	}
}

func TestCompressStringWithSymbolCount(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "example1",
			in:   "##.#.",
			out:  "2#1.1#1.",
		},
		{
			name: "example2",
			in:   ".###...##.#.",
			out:  "1.3#3.2#1.1#1.",
		},
		{
			name: "example3",
			in:   ".###.##...#.",
			out:  "1.3#1.2#3.1#1.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := compress(test.in)
			if actual != test.out {
				t.Errorf("Expected %v, got %v", test.out, actual)
			}
		})
	}
}
