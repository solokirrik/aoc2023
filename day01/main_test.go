package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestTask1Example(t *testing.T) {
	input := []byte(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)
	expectedSum := 12 + 38 + 15 + 77
	sum := task1(input)

	t.Log(sum)

	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}
}

func TestTask1(t *testing.T) {
	input := fInput
	expectedSum := 54388

	sum := task1(input)
	t.Log(sum)
	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}
}

func TestTask2Example(t *testing.T) {
	input := []byte(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)
	expectedSum := 29 + 83 + 13 + 24 + 42 + 14 + 76
	sum := task2(input)

	t.Log(sum)

	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}
}

// 53515 -> not 53519|53576|53597|53648|53652
func TestTask2(t *testing.T) {
	input := fInput
	expectedSum := 53515
	sum := task2(input)

	t.Log(sum)

	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}
}

func TestReplace(t *testing.T) {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	inp := "four3threezdszkzkbhnsqpmsninebq"
	exp := "433zdszkzkbhnsqpms9bq"

	for k, v := range numbers {
		inp = strings.ReplaceAll(inp, k, strconv.Itoa(v))
	}

	if inp != exp {
		t.Errorf("Expected %s, but got %s", exp, inp)
	}
}
