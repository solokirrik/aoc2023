package main

import (
	"fmt"
	"sort"
	"strings"
)

func Task2(rawLines []byte) int {
	sequenceRaw, networkMap := parse(rawLines)

	walkers := extractWalkers(networkMap)
	walkers = walkersStepsToFinish2(walkers, sequenceRaw, networkMap)

	rawDivisors := extractDivisors(walkers)
	sort.Slice(rawDivisors, func(i, j int) bool {
		return rawDivisors[i] < rawDivisors[j]
	})

	out := 1
	for _, divisor := range rawDivisors {
		out *= divisor
	}

	return out
}

func extractWalkers(networkMap map[string][]string) []State {
	walkers := make([]State, 0, len(networkMap))

	for node := range networkMap {
		if strings.Contains(node, "A") {
			walkers = append(walkers, State{Position: node})
		}
	}

	return walkers
}

func extractDivisors(walkers []State) []int {
	seenDivisors := make(map[int]bool)
	rawDivisors := []int{}

	for _, walker := range walkers {
		divisors := numberDivisors(walker.StepsToReach)
		for _, divisor := range divisors {
			if seenDivisors[divisor] || walker.StepsToReach == divisor {
				continue
			}
			rawDivisors = append(rawDivisors, divisor)
			seenDivisors[divisor] = true
		}
	}

	return rawDivisors
}

func walkersStepsToFinish2(walkers []State, sequenceRaw []byte, networkMap map[string][]string) []State {
	for i := range walkers {
		walker := walkers[i]
		curNode := walker.Position

		for !strings.Contains(curNode, "Z") {
			walker.StepsToReach++
			direction := sequenceRaw[(walker.StepsToReach-1)%len(sequenceRaw)]
			curNode = networkMap[curNode][directionToIndex(direction)]
		}

		walkers[i] = walker
	}

	return walkers
}

func numberDivisors(x int) []int {
	var divisors []int

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func greatestCommonDivisor(x int, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return abs(x)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func leastCommonMultiple(firstNum, secondNum int) int {
	return firstNum / greatestCommonDivisor(firstNum, secondNum) * secondNum
}

func areOnZ(w []State) bool {
	for _, walker := range w {
		if !strings.Contains(walker.Position, "Z") {
			return false
		}
	}
	return true
}

// 00.003.199.776.350 <- bruteforce interrupted
// 13.740.108.158.591 <- proper answer
func task2FailedBruteforce(sequenceRaw []byte, networkMap map[string][]string, walkers []State) int {
	out := 0

	for !areOnZ(walkers) {
		out++
		if out > 10*1000*1000*1000*1000 {
			fmt.Println("too many iterations")
			break
		}
		direction := sequenceRaw[(out-1)%len(sequenceRaw)]

		for i, walker := range walkers {
			newPosition := networkMap[walker.Position][directionToIndex(direction)]
			if strings.Contains(newPosition, "Z") {
				fmt.Println(out, i, walker.Position, newPosition)
			}

			walkers[i].Position = newPosition
		}
	}

	return out
}
