package main

import (
	_ "embed"
	"testing"
)

//go:embed input-example
var inputExample []byte

func TestTask1Example(t *testing.T) {
	out := Task1(ToHands(inputExample))
	want := 765*1 + 220*2 + 28*3 + 684*4 + 483*5
	if out != want {
		t.Errorf("Task1(input-example) = %d, want %d", out, want)
	}
}
