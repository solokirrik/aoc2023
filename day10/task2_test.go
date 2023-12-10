package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input-example3
var inputExample3 []byte

//go:embed input-example4
var inputExample4 []byte

//go:embed input-example5
var inputExample5 []byte

func TestTask2(t *testing.T) {
	type tCase struct {
		name     string
		input    []byte
		expected int
		doPrint  bool
	}

	doPrint := false

	cases := []tCase{
		{"example3", inputExample3, 4, doPrint},
		{"example4", inputExample4, 8, doPrint},
		{"example5", inputExample5, 10, doPrint},
		{"real", input, 435, doPrint},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			inp := bytes.Split(cases[i].input, []byte("\n"))

			out := Task2FloodFromEdge(inp, pointInPolygonFast, cases[i].doPrint)
			if out != cases[i].expected {
				t.Errorf("Expected %d, got %d", cases[i].expected, out)
			}
		})
	}
}

func TestAngle(t *testing.T) {
	pointA := point{0, 0}
	pointB := point{1, 0}
	pointC := point{0, 1}

	vectorAC := NewVector(pointC, pointA)
	vectorBC := NewVector(pointC, pointB)

	angle := AngleBetweenVectors(vectorAC, vectorBC)

	fmt.Printf("Angle between vectors AC and BC: %.4f radians\n", angle)
}
