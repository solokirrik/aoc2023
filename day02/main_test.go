package main

import (
	_ "embed"
	"testing"
)

//go:embed input_example
var inputExample string

func TestTask1Example(t *testing.T) {
	want := 8
	got := task1(inputExample)

	t.Log(got)

	if got != want {
		t.Errorf("task1(inputExample) = %d, want %d", got, want)
	}
}

func TestTask1(t *testing.T) {
	want := 2505
	got := task1(input)

	t.Log(got)

	if got != want {
		t.Errorf("task1(inputExample) = %d, want %d", got, want)
	}
}

func TestTask2Example(t *testing.T) {
	want := 2286
	got := task2(inputExample)

	t.Log(got)

	if got != want {
		t.Errorf("task1(inputExample) = %d, want %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	want := 70265
	got := task2(input)

	t.Log(got)

	if got != want {
		t.Errorf("task1(inputExample) = %d, want %d", got, want)
	}
}
