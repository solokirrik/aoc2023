package main

import (
	"slices"
	"sort"
	"strings"
)

func Task2(onHands []Hand) int {
	out := 0

	onHands = sortHands2(onHands)
	for i := 1; i <= len(onHands); i++ {
		out += onHands[i-1].bid * i
	}

	return out
}

type CardCount struct {
	card  string
	count int
}

func handToBest(hand string) string {
	if !strings.Contains(hand, "J") {
		return hand
	}

	if strings.Count(hand, "J") == 5 {
		return "AAAAA"
	}

	cards := []CardCount{}
	cardsIndex := map[string]int{}

	// count cards
	for _, p := range strings.Split(hand, "") {
		if p == "J" {
			continue
		}

		idx, ok := cardsIndex[p]
		if !ok {
			cardsIndex[p] = len(cards)
			cards = append(cards, CardCount{card: p, count: 1})
			continue
		}

		cards[idx].count++
	}

	// sort by power
	sort.Slice(cards, func(i, j int) bool {
		// sort by power if count is equal
		if cards[i].count == cards[j].count {
			iCharRate := strings.IndexAny(strings.Join(power2, ""), cards[i].card)
			jCharRate := strings.IndexAny(strings.Join(power2, ""), cards[j].card)

			return iCharRate < jCharRate
		}

		return cards[i].count > cards[j].count
	})

	// pick best card to replace joker
	if bestCard := cards[0].count; bestCard > 1 {
		newHand := strings.ReplaceAll(hand, "J", cards[0].card)
		return newHand
	}

	// replace joker with most powerful card among all single cards
	return strings.ReplaceAll(hand, "J", cards[0].card)

}

func sortHands2(onHands []Hand) []Hand {
	slices.SortFunc(onHands, func(a, b Hand) int {
		aRate := handToPower(handToBest(a.hand))
		bRate := handToPower(handToBest(b.hand))

		if aRate < bRate {
			return 1
		}

		if aRate > bRate {
			return -1
		}

		for i := 0; i < len(a.hand); i++ {
			aCharRate := strings.IndexAny(strings.Join(power2, ""), string([]byte(a.hand)[i]))
			bCharRate := strings.IndexAny(strings.Join(power2, ""), string([]byte(b.hand)[i]))

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
