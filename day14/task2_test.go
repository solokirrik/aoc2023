package main

import (
	"bytes"
	"testing"
)

func TestTask2(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		want   int
		cycles int
	}{
		{
			name:   "example-test",
			input:  inputExample,
			want:   69,
			cycles: 1,
		},
		{
			name:   "example-simple",
			input:  inputExample,
			want:   69,
			cycles: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := Task2(bytes.Fields(tt.input), tt.cycles)
			if out != tt.want {
				t.Errorf("Expected %d, got %d", tt.want, out)
			}
		})
	}
}
