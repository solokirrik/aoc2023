package main

import (
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	slog.Info("Task1", "sum", task1(input))
	slog.Info("Task2", "sum", task2(input))
}

func task1(inp string) int {
	lines := strings.Split(inp, "\n")
	sum := 0

	for _, line := range lines {
		sum += getValidGame(line)
	}

	return sum
}

func task2(inp string) int {
	lines := strings.Split(inp, "\n")
	mult := 0

	for _, line := range lines {
		mult += gameMinMult(line)
	}

	return mult
}

func gameMinMult(line string) int {
	minRed, minGreen, minBlue := 0, 0, 0
	_, colors := gameData(line)

	for _, balls := range colors {
		if balls.Color == "red" && (balls.Count) > minRed {
			minRed = (balls.Count)
			continue
		}

		if balls.Color == "green" && (balls.Count) > minGreen {
			minGreen = (balls.Count)
			continue
		}

		if balls.Color == "blue" && (balls.Count) > minBlue {
			minBlue = (balls.Count)
			continue
		}
	}

	return minRed * minGreen * minBlue
}

type BallsData struct {
	Count int
	Color string
}

func getValidGame(line string) int {
	line = strings.TrimLeft(line, "Game ")

	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)

	failed := false
	id, colors := gameData(line)

	for _, balls := range colors {
		if (balls.Color == "red" && balls.Count > maxRed) ||
			(balls.Color == "green" && balls.Count > maxGreen) ||
			(balls.Color == "blue" && balls.Count > maxBlue) {
			failed = true
			break
		}
	}

	if failed {
		return 0
	}

	return id
}

func gameData(line string) (int, []BallsData) {
	parts := strings.Split(line, ": ")
	id, _ := strconv.Atoi(string(parts[0]))
	data := strings.Split(strings.ReplaceAll(parts[1], ";", ","), ", ")

	out := make([]BallsData, len(data))
	for i, d := range data {
		parts := strings.Split(d, " ")
		count, _ := strconv.Atoi(parts[0])
		out[i] = BallsData{Count: count, Color: parts[1]}
	}

	return id, out
}
