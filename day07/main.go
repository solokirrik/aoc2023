package main

import (
	_ "embed"
	"fmt"
)

var power1 = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
var power2 = []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}
var hands = []string{"FiveOfAKind", "FourOfAKind", "FullHouse", "ThreeOfAKind", "TwoPair", "OnePair", "HighCard"}
var handPower = map[string]func(h string) bool{
	"FiveOfAKind":  IsFiveOfAKind,
	"FourOfAKind":  IsFourOfAKind,
	"FullHouse":    IsFullHouse,
	"ThreeOfAKind": IsThreeOfAKind,
	"TwoPair":      IsTwoPair,
	"OnePair":      IsOnePair,
	"HighCard":     IsHighCard,
}

//go:embed input
var input []byte

/*
- AAAAA - Five of a kind - where all five cards have the same label
- AA8AA - Four of a kind - where four cards have the same label and one card has a different label
- 23332 - Full house - where three cards have the same label, and the remaining two cards share a different label
- TTT98 - Three of a kind - where three cards have the same label, and the remaining two cards are each different from any other card in the hand
- 23432 - Two pair - where two cards share one label, two other cards share a second label, and the remaining card has a third label
- A23A4 - One pair - where two cards share one label, and the other three cards have a different label from the pair and each other
- 23456 - High card - where all cards' labels are distinct
*/

func main() {
	onHands := ToHands(input)

	out := Task1(onHands)
	fmt.Println(251029473 == out, out)

	out2 := Task2(onHands)
	fmt.Println(44218354 == out2, out2)
}
