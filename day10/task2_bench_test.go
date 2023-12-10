package main

import (
	"bytes"
	"testing"
)

func BenchmarkTask2FloodFast(b *testing.B) {
	inp := bytes.Split(input, []byte("\n"))

	for i := 0; i < b.N; i++ {
		Task2FloodFromEdge(inp, pointInPolygonFast, false)
	}
}

func BenchmarkTask2FloodSlow(b *testing.B) {
	inp := bytes.Split(input, []byte("\n"))

	for i := 0; i < b.N; i++ {
		Task2FloodFromEdge(inp, pointInPolygonSlow, false)
	}
}

func BenchmarkTask2NonFloodFast(b *testing.B) {
	inp := bytes.Split(input, []byte("\n"))

	for i := 0; i < b.N; i++ {
		Task2NonFlood(inp, pointInPolygonFast, false)
	}
}

func BenchmarkTask2NonFloodSlow(b *testing.B) {
	inp := bytes.Split(input, []byte("\n"))

	for i := 0; i < b.N; i++ {
		Task2NonFlood(inp, pointInPolygonSlow, false)
	}
}
