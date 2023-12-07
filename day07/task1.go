package main

import (
	"slices"
	"strings"
)

func Task1(onHands []Hand) int {
	out := 0

	onHands = sortHands1(onHands)
	for i := 1; i <= len(hands); i++ {
		// fmt.Println(hands[i-1], i)
		out += onHands[i-1].bid * i
	}

	return out
}

func handToPower(hand string) int {
	for i, p := range hands {
		if handPower[p](hand) {
			return i
		}
	}

	return -1
}

func sortHands1(onHands []Hand) []Hand {
	slices.SortFunc(onHands, func(a, b Hand) int {
		aRate := handToPower(a.hand)
		bRate := handToPower(b.hand)

		if aRate < bRate {
			return 1
		}
		if aRate > bRate {
			return -1
		}

		for i := 0; i < len(a.hand); i++ {
			aCharRate := strings.IndexAny(strings.Join(power1, ""), string([]byte(a.hand)[i]))
			bCharRate := strings.IndexAny(strings.Join(power1, ""), string([]byte(b.hand)[i]))

			if aCharRate < bCharRate {
				return 1
			}

			if aCharRate > bCharRate {
				return -1
			}
		}

		return 0
	})

	return onHands
}
