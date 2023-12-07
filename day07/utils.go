package main

import (
	"bytes"
	"strconv"
	"strings"
)

type Hand struct {
	hand string
	bid  int
}

func ToHands(input []byte) []Hand {
	lines := bytes.Split(input, []byte("\n"))
	out := make([]Hand, 0, len(lines))

	for _, line := range lines {
		ll := bytes.Split(line, []byte(" "))
		bid, err := strconv.Atoi(string(ll[1]))
		panicOnError(err)

		out = append(out, Hand{hand: string(ll[0]), bid: bid})
	}

	return out
}

func IsFiveOfAKind(hand string) bool {
	for _, p := range strings.Split(hand, "") {
		if strings.Count(hand, p) == 5 {
			return true
		}
	}
	return false
}

func IsFourOfAKind(hand string) bool {
	for _, p := range strings.Split(hand, "") {
		if strings.Count(hand, p) == 4 {
			return true
		}
	}
	return false
}

func IsFullHouse(hand string) bool {
	isTriple, char := isOneTriple(hand)

	if !isTriple {
		return false
	}

	isPair, _ := isExactOnePair(hand, char)

	return isTriple && isPair
}

func IsThreeOfAKind(hand string) bool {
	isTwo := false
	isThree := false

	for _, p := range strings.Split(hand, "") {
		if !isThree && strings.Count(hand, p) == 3 {
			isThree = true
		}
		if !isTwo && strings.Count(hand, p) == 2 {
			isTwo = true
		}
	}

	return isThree && !isTwo
}

func IsTwoPair(hand string) bool {
	gotPair, pairCard := isExactOnePair(hand, "")
	if !gotPair {
		return false
	}

	goSecondPair, _ := isExactOnePair(hand, pairCard)
	if goSecondPair {
		return true
	}

	return false
}

func isExactOnePair(hand, except string) (bool, string) {
	for _, p := range strings.Split(hand, "") {
		if p == except {
			continue
		}
		if strings.Count(hand, p) == 2 {
			return true, p
		}
	}

	return false, ""
}

func isOneTriple(hand string) (bool, string) {
	for _, p := range strings.Split(hand, "") {
		if strings.Count(hand, p) == 3 {
			return true, p
		}
	}
	return false, ""
}

func IsOnePair(hand string) bool {
	gotPair, pairCard := isExactOnePair(hand, "")
	if !gotPair {
		return false
	}

	gotSecondPair, _ := isExactOnePair(hand, pairCard)
	if !gotSecondPair {
		return true
	}

	return false
}

func IsHighCard(hand string) bool {
	for _, p := range strings.Split(hand, "") {
		if strings.Count(hand, p) != 1 {
			return false
		}
	}
	return true
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
