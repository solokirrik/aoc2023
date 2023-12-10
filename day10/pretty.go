package main

import (
	"fmt"
)

const (
	HEADER    = "\033[95m"
	OKBLUE    = "\033[94m"
	OKCYAN    = "\033[96m"
	OKGREEN   = "\033[92m"
	YELLOW    = "\033[93m"
	RED       = "\033[91m"
	BOLD      = "\033[1m"
	UNDERLINE = "\033[4m"
	ENDC      = "\033[0m"
)

func prettyPrint1(inp [][]byte, loopMap map[point]int) {
	for r := range inp {
		for c := range inp[r] {
			if idx, ok := loopMap[point{y: r, x: c}]; ok {
				if inp[r][c] == 'S' {
					fmt.Print(RED, string(inp[r][c]), ENDC)
					continue
				}

				if idx > 0 && idx < len(loopMap)-2 {
					fmt.Print(OKGREEN, string(inp[r][c]), ENDC)
					continue
				} else {
					fmt.Print(YELLOW, string(inp[r][c]), ENDC)
					continue
				}
			}

			fmt.Print(string(inp[r][c]))

		}
		fmt.Println()
	}
}

func prettyPrint2(inp [][]byte, loopMap, internal map[point]int) {
	for r := range inp {
		for c := range inp[r] {
			if idx, ok := loopMap[point{y: r, x: c}]; ok {
				if idx > 0 && idx < len(loopMap)-2 {
					fmt.Print(OKGREEN, string(inp[r][c]), ENDC)
					continue
				} else {
					fmt.Print(YELLOW, string(inp[r][c]), ENDC)
					continue
				}
			}

			if _, ok := internal[point{y: r, x: c}]; ok {
				fmt.Print(RED, string(inp[r][c]), ENDC)
				continue
			}

			fmt.Print(OKBLUE, string(inp[r][c]), ENDC)
		}
		fmt.Println()
	}
}
