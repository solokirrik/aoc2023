package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input-example1
var inputExample1 []byte

//go:embed input-example2
var inputExample2 []byte

func TestColor(t *testing.T) {
	t.Logf("some text %s%s%s again gray text \n", OKGREEN, "test-part", ENDC)
}

func TestTask1(t *testing.T) {
	type tCase struct {
		name     string
		input    []byte
		expected int
	}

	cases := []tCase{
		{"example1", inputExample1, 4},
		{"example2", inputExample2, 8},
		{"real", input, 6733},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			inp := bytes.Split(cases[i].input, []byte("\n"))

			out := Task1(inp, false)
			if out != cases[i].expected {
				t.Errorf("Expected %d, got %d", cases[i].expected, out)
			}
		})
	}
}

func TestTask1Matching(t *testing.T) {
	inp := [][]byte{
		{'L', '7', '|', 'F'},
		{'-', 'S', '|', 'L'},
		{'-', '7', 'L', '7'},
	}

	nodes := make([]Node, 0, 3*4)
	for r := range inp {
		for c := range inp[r] {
			nodes = append(nodes, *getNodeWithEnds(inp, r, c))
		}
	}

	wanted := []bool{
		// L
		false, true, false, false, /**/
		false, false, false, false, /**/
		false, false, false, false,
		// 7
		true, false, false, false, /**/
		false, true, false, false, /**/
		false, false, false, false,
		// |
		false, false, false, false, /**/
		false, false, true, false, /**/
		false, false, false, false,
		// F
		false, false, false, false, /**/
		false, false, false, true, /**/
		false, false, false, false,
		// -
		false, false, false, false, /**/
		false, true, false, false, /**/
		false, false, false, false,
		// S
		false, true, false, false, /**/
		true, false, false, false, /**/
		false, false, false, false,
		// |
		false, false, true, false, /**/
		false, false, false, false, /**/
		false, false, true, false,
		// L
		false, false, false, true, /**/
		false, false, false, false, /**/
		false, false, false, false,
		// -
		false, false, false, false, /**/
		false, false, false, false, /**/
		false, true, false, false,
		// 7
		false, false, false, false, /**/
		false, false, false, false, /**/
		true, false, false, false,
		// L
		false, false, false, false, /**/
		false, false, true, false, /**/
		false, false, false, true,
		// 7
		false, false, false, false, /**/
		false, false, false, false, /**/
		false, false, true, false,
	}

	for n := range nodes {
		for sn := range nodes {
			t.Run(fmt.Sprintf("%s -> %s", nodes[n], nodes[sn]), func(t *testing.T) {
				got := matching(nodes[n], nodes[sn])
				if got != wanted[n*len(nodes)+sn] {
					t.Errorf("Expected %s -> %s %v, got %v", nodes[n], nodes[sn], wanted[n*len(nodes)+sn], got)
				}
			})
		}
	}
}
