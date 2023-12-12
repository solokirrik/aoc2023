package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input
var input []byte

func main() {
	lines := bytes.Split(input, []byte("\n"))

	var cache1 = make(map[string]int)
	var cache2 = make(map[string]int)

	out := Task1(cache1, lines)
	fmt.Println(7169 == out, out)

	out = Task2(cache2, lines)
	fmt.Println(1738259948652 == out, out)
}

func Task1(cache map[string]int, lines [][]byte) int {
	out := 0

	for i := range lines {
		lineParts := bytes.Split(lines[i], []byte(" "))
		hashes := CountIntoHashes(lineParts[1])
		out += countArrangements(cache, lineParts[0], hashes)
	}

	return out
}

func Task2(cache map[string]int, lines [][]byte) int {
	const copies = 5

	out := 0

	for i := range lines {
		hashes := []int{}
		springParts := make([][]byte, 0, copies)
		lineParts := bytes.Split(lines[i], []byte(" "))

		for i := 0; i < copies; i++ {
			hashes = append(hashes, CountIntoHashes(lineParts[1])...)
			springParts = append(springParts, lineParts[0])
		}

		out += countArrangements(cache, bytes.Join(springParts, []byte("?")), hashes)
	}

	return out
}

func CountIntoHashes(c []byte) []int {
	hashes := []int{}

	for _, hash := range bytes.Split(c, []byte(",")) {
		hashInt := int(hash[0] - 48)
		if len(hash) > 1 {
			hashInt = hashInt*10 + int(hash[1]-48)
		}

		hashes = append(hashes, hashInt)
	}

	return hashes
}

func countArrangements(cache map[string]int, template []byte, hashes []int) int {
	if len(template) == 0 && len(hashes) == 0 {
		return 1
	}

	if len(template) == 0 {
		return 0
	}

	key := string(template)
	for _, group := range hashes {
		key += strconv.Itoa(group) + ","
	}

	if count, ok := cache[key]; ok {
		return count
	}

	if bytes.HasPrefix(template, []byte("?")) {
		return countArrangements(cache, bytes.Replace(template, []byte("?"), []byte("."), 1), hashes) +
			countArrangements(cache, bytes.Replace(template, []byte("?"), []byte("#"), 1), hashes)
	}

	if bytes.HasPrefix(template, []byte(".")) {
		cache[key] = countArrangements(cache, bytes.TrimPrefix(template, []byte(".")), hashes)
		return cache[key]
	}

	if bytes.HasPrefix(template, []byte("#")) {
		if len(hashes) == 0 ||
			len(template) < hashes[0] ||
			bytes.Contains(template[0:hashes[0]], []byte(".")) {
			cache[key] = 0
			return 0
		}

		if len(hashes) > 1 {
			if len(template) < hashes[0]+1 || template[hashes[0]] == '#' {
				cache[key] = 0
				return 0
			}

			cache[key] = countArrangements(cache, template[hashes[0]+1:], hashes[1:])
			return cache[key]
		}

		cache[key] = countArrangements(cache, template[hashes[0]:], hashes[1:])
		return cache[key]
	}

	return 0
}
