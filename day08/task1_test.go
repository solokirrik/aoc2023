package main

import (
	_ "embed"
	"testing"
)

//go:embed input-example
var inputExample []byte

func TestTask1Example(t *testing.T) {
	want := 2
	if got := Task1(inputExample); got != want {
		t.Errorf("Task1() = %v, want %v", got, want)
	}
}

func TestTask1(t *testing.T) {
	want := 11309
	if got := Task1(input); got != want {
		t.Errorf("Task1() = %v, want %v", got, want)
	}
}
