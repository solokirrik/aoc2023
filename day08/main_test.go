package main

import (
	_ "embed"
	"testing"
)

//go:embed input-example
var inputExample []byte

//go:embed input-example2
var inputExample2 []byte

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

// broken after task2 solution which works with the input, but not the input-example :sadface:
func TestTask2Example(t *testing.T) {
	want := 6
	if got := Task2(inputExample2); got != want {
		t.Errorf("Task2() = %v, want %v", got, want)
	}
}

func TestTask2(t *testing.T) {
	want := 13740108158591
	if got := Task2(input); got != want {
		t.Errorf("Task1() = %v, want %v", got, want)
	}
}
