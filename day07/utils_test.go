package main

import "testing"

func TestToHands(t *testing.T) {
	out := ToHands(inputExample)
	want := []Hand{
		{hand: "32T3K", bid: 765},
		{hand: "T55J5", bid: 684},
		{hand: "KK677", bid: 28},
		{hand: "KTJJT", bid: 220},
		{hand: "QQQJA", bid: 483},
	}

	if len(out) != len(want) {
		t.Errorf("len(ToHands(input-example)) = %d, want %d", len(out), len(want))
	}

	for i, h := range out {
		if h.bid != want[i].bid || h.hand != want[i].hand {
			t.Errorf("ToHands(input-example)[%d] = %v, want %v", i, h, want[i])
		}
	}
}

func TestIsOnePair(t *testing.T) {
	inputs := []string{
		"A23A4",
		"A23A3",
	}
	outs := []bool{
		true,
		false,
	}

	for i, in := range inputs {
		t.Run(in, func(t *testing.T) {
			out := IsOnePair(in)
			if out != outs[i] {
				t.Errorf("IsOnePair(%s) = %t, want %t", in, out, outs[i])
			}
		})
	}
}

func TestIsTwoPair(t *testing.T) {
	inputs := []string{
		"23432",
		"23332",
		"23512",
	}
	outs := []bool{
		true,
		false,
		false,
	}

	for i, in := range inputs {
		t.Run(in, func(t *testing.T) {
			out := IsTwoPair(in)
			if out != outs[i] {
				t.Errorf("IsTwoPair(%s) = %t, want %t", in, out, outs[i])
			}
		})
	}
}

func TestIsHighCard(t *testing.T) {
	inputs := []string{
		"23432",
		"23332",
		"23512",
		"23514",
	}
	outs := []bool{
		false,
		false,
		false,
		true,
	}

	for i, in := range inputs {
		t.Run(in, func(t *testing.T) {
			out := IsHighCard(in)
			if out != outs[i] {
				t.Errorf("IsHighCard(%s) = %t, want %t", in, out, outs[i])
			}
		})
	}
}

func TestIsThreeOfAKind(t *testing.T) {
	inputs := []string{
		"TTT98",
		"23332",
		"23512",
		"23514",
	}
	outs := []bool{
		true,
		false,
		false,
		false,
	}

	for i, in := range inputs {
		t.Run(in, func(t *testing.T) {
			out := IsThreeOfAKind(in)
			if out != outs[i] {
				t.Errorf("IsThreeOfAKind(%s) = %t, want %t", in, out, outs[i])
			}
		})
	}
}

func TestIsFullHouse(t *testing.T) {
	inputs := []string{
		"TTT98",
		"23332",
		"23512",
		"23514",
	}
	outs := []bool{
		false,
		true,
		false,
		false,
	}

	for i, in := range inputs {
		t.Run(in, func(t *testing.T) {
			out := IsFullHouse(in)
			if out != outs[i] {
				t.Errorf("IsFullHouse(%s) = %t, want %t", in, out, outs[i])
			}
		})
	}
}

func TestIsFourOfAKind(t *testing.T) {
	inputs := []string{
		"AA8AA",
		"TTT98",
		"23332",
		"23512",
		"23514",
	}
	outs := []bool{
		true,
		false,
		false,
		false,
		false,
	}

	for i, in := range inputs {
		t.Run(in, func(t *testing.T) {
			out := IsFourOfAKind(in)
			if out != outs[i] {
				t.Errorf("IsFourOfAKind(%s) = %t, want %t", in, out, outs[i])
			}
		})
	}
}

func TestIsFiveOfAKind(t *testing.T) {
	inputs := []string{
		"AAAAA",
		"AA8AA",
		"TTT98",
		"23332",
		"23512",
		"23514",
	}
	outs := []bool{
		true,
		false,
		false,
		false,
		false,
		false,
	}

	for i, in := range inputs {
		t.Run(in, func(t *testing.T) {
			out := IsFiveOfAKind(in)
			if out != outs[i] {
				t.Errorf("IsFiveOfAKind(%s) = %t, want %t", in, out, outs[i])
			}
		})
	}
}
