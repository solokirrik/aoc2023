package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func readFile(t *testing.T, path string) []byte {
	f, err := os.Open(path)

	defer f.Close()

	if err != nil {
		t.Fatal(err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	return data
}

func TestTask1(t *testing.T) {
	tests := []struct {
		name  string
		input [][]byte
		want  int
	}{
		{
			name:  "example1",
			input: bytes.Split(readFile(t, "./inputs/input-example-1"), []byte("\n\n")),
			want:  5,
		},
		{
			name:  "example2",
			input: bytes.Split(readFile(t, "./inputs/input-example-2"), []byte("\n\n")),
			want:  4 * 100,
		},
		{
			name:  "example",
			input: bytes.Split(readFile(t, "./inputs/input-example"), []byte("\n\n")),
			want:  4*100 + 5,
		},
		{
			name:  "input-test-1",
			input: bytes.Split(readFile(t, "./inputs/input-test-1"), []byte("\n\n")),
			want:  2 * 100,
		},
		{
			name:  "input-test-10",
			input: bytes.Split(readFile(t, "./inputs/input-test-10"), []byte("\n\n")),
			want:  6 * 100,
		},
		{
			name:  "input-test-12",
			input: bytes.Split(readFile(t, "./inputs/input-test-12"), []byte("\n\n")),
			want:  4 * 100,
		},
		{
			name:  "input-test-21",
			input: bytes.Split(readFile(t, "./inputs/input-test-21"), []byte("\n\n")),
			want:  2 * 100,
		},
		{
			name:  "input-test-r-1",
			input: bytes.Split(readFile(t, "./inputs/input-test-r-1"), []byte("\n\n")),
			want:  12,
		},
		{
			name:  "input-test-r-2",
			input: bytes.Split(readFile(t, "./inputs/input-test-r-2"), []byte("\n\n")),
			want:  1,
		},
		{
			name:  "input",
			input: bytes.Split(input, []byte("\n\n")),
			want:  28651,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out1 := Task1(tt.input)
			if out1 != tt.want {
				t.Errorf("%s got=%d, expected=%d", tt.name, out1, tt.want)
			}
		})
	}
}

func TestTask2(t *testing.T) {
	tests := []struct {
		name  string
		input [][]byte
		want  int
	}{
		{
			name:  "example1",
			input: bytes.Split(readFile(t, "./inputs/input-example-1"), []byte("\n\n")),
			want:  300,
		},
		{
			name:  "example2",
			input: bytes.Split(readFile(t, "./inputs/nput-example-2"), []byte("\n\n")),
			want:  100,
		},
		{
			name:  "example",
			input: bytes.Split(readFile(t, "./inputs/input-example"), []byte("\n\n")),
			want:  4*100 + 100,
		},
		{
			name:  "input",
			input: bytes.Split(input, []byte("\n\n")),
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out1 := Task2(tt.input)
			if out1 != tt.want {
				t.Errorf("%s got=%d, expected=%d", tt.name, out1, tt.want)
			}
		})
	}
}
