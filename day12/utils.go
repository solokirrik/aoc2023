package main

import (
	"fmt"
)

func compress(s string) string {
	compressed := ""
	count := 0
	current := s[0]

	for i := 0; i < len(s); i++ {
		if s[i] == current {
			count++
		} else {
			compressed += fmt.Sprintf("%d%c", count, current)
			count = 1
			current = s[i]
		}
	}

	compressed += fmt.Sprintf("%d%c", count, current)
	return compressed
}

func generateCombinations(n int, prefix string, combinations *[]string) {
	if n == 0 {
		*combinations = append(*combinations, prefix)
	} else {
		generateCombinations(n-1, prefix+".", combinations)
		generateCombinations(n-1, prefix+"#", combinations)
	}
}
