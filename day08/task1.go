package main

func Task1(rawLines []byte) int {
	i := 0
	curNode := "AAA"
	sequenceRaw, networkMap := parse(rawLines)

	for curNode != "ZZZ" {
		i++
		direction := sequenceRaw[(i-1)%len(sequenceRaw)]
		curNode = networkMap[curNode][directionToIndex(direction)]
	}

	return i
}
