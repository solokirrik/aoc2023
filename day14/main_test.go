package main

import (
	"bytes"
	"testing"
)

func TestTask1(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "example",
			input: inputExample,
			want:  136,
		},
		{
			name:  "task1",
			input: input,
			want:  106648,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := Task1(bytes.Split(tt.input, []byte("\n")))
			if out != tt.want {
				t.Errorf("Expected %d, got %d", tt.want, out)
			}
		})
	}
}
