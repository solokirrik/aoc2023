package main

import (
	"bytes"
	_ "embed"
	"log/slog"
)

//go:embed input
var fInput []byte

func main() {
	slog.Info("Result1", "sum", task1(fInput))
	slog.Info("Result2", "sum", task2(fInput))
}

func task1(input []byte) int {
	sum := 0
	lines := bytes.Split(input, []byte("\n"))

	for _, line := range lines {
		foundDigits := []int{}

		for _, char := range line {
			if char >= '1' && char <= '9' {
				foundDigits = append(foundDigits, int(char-48))
			}
		}

		if len(foundDigits) > 0 {
			sum += foundDigits[0] * 10
			sum += foundDigits[len(foundDigits)-1]
		}
	}

	return sum
}

var digits = map[string]int{
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

var firstDigLetters = map[byte]interface{}{
	'o': nil,
	'n': nil,
	'e': nil,
	't': nil,
	'f': nil,
	's': nil,
}

func task2(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	sum := 0

	for _, line := range lines {
		foundDigits := []int{}

		for c, char := range line {
			if char >= '1' && char <= '9' {
				foundDigits = append(foundDigits, int(char-48))
				continue
			}

			if _, ok := firstDigLetters[char]; ok {
				for digitWord, digit := range digits {
					if bytes.HasPrefix(line[c:], []byte(digitWord)) {
						foundDigits = append(foundDigits, digit)
						break
					}
				}

				continue
			}
		}

		if len(foundDigits) > 0 {
			sum += foundDigits[0] * 10
			sum += foundDigits[len(foundDigits)-1]
		}
	}

	return sum
}

var bnumbers = map[string]byte{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

// didn't worked because:
// 1. map keys are not ordered
// 2. digits words can be part of other number words
// combination of 1 and 2 makes it impossible to replace words in order without an
// extra try-call to ReplaceFirst for suitable digit word
func task2fail(input []byte) int {
	sum := 0
	lines := bytes.Split(input, []byte("\n"))

	for _, line := range lines {
		for k, v := range bnumbers {
			line = bytes.ReplaceAll(line, []byte(k), []byte{v})
		}
		sum += task1(line)
	}

	return sum
}
