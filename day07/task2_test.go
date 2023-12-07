package main

import (
	"testing"
)

func TestTask2Example(t *testing.T) {
	out := Task2(ToHands(inputExample))
	want := 5905
	if out != want {
		t.Errorf("Task2(input-example) = %d, want %d", out, want)
	}
}

func TestTask2(t *testing.T) {
	out := Task2(ToHands(input))
	want := 251003917
	if out != want {
		t.Errorf("Task2(input-example) = %d, want %d", out, want)
	}
}

func TestHandToBest(t *testing.T) {
	cases := []string{
		"QJ4J7",
		"25J54",
		"25J52",
		"J5J5J",
		"8AJAA",
		"AAAAA",
		"JAAAA",
		"JJAAA",
		"JJJAA",
		"JJJJA",
		"JJJJJ",
		"AAJKK",
		"AKJJJ",
		"JJJAK",
		"JJAKC",
		"JAAKC",
		"JAKCD",
	}
	wanted := []string{
		"QQ4Q7",
		"25554",
		"25552",
		"55555",
		"8AAAA",
		"AAAAA",
		"AAAAA",
		"AAAAA",
		"AAAAA",
		"AAAAA",
		"AAAAA",
		"AAAKK",
		"AKAAA",
		"AAAAK",
		"AAAKC",
		"AAAKC",
		"AAKCD",
	}

	for i, c := range cases {
		t.Run(c+"-"+wanted[i], func(t *testing.T) {
			out := handToBest(c)
			if out != wanted[i] {
				t.Errorf("handToBest(%s) == %s, wanted %s", c, out, wanted[i])
			}
		})
	}
}

func TestSort2(t *testing.T) {
	inp := `A3JAA 49
A5AAA 720
A7J7J 472
A8JAA 571
AQQJJ 978
AQQQQ 507
AA2AA 126
AA7AA 279
AAATA 230
AAAA8 78
TJTTT 857
J22J2 372
J44J4 220
J555J 777
J9999 431
J9JJ9 454
JJJJJ 131
JJJAA 638
QQJQJ 551
QQJQQ 622
KJJKK 427
KKKKJ 18
AJAAA 405
AAJAJ 959`
	onHands := ToHands([]byte(inp))
	want := []string{
		"A3JAA",
		"A5AAA",
		"A7J7J",
		"A8JAA",
		"AQQJJ",
		"AQQQQ",
		"AA2AA",
		"AA7AA",
		"AAATA",
		"AAAA8",
		"JJJJJ",
		"JJJAA",
		"J22J2",
		"J44J4",
		"J555J",
		"J9JJ9",
		"J9999",
		"TJTTT",
		"QQJQJ",
		"QQJQQ",
		"KJJKK",
		"KKKKJ",
		"AJAAA",
		"AAJAJ",
	}

	got := sortHands2(onHands)
	out := []string{}
	for _, h := range got {
		out = append(out, h.hand)
	}

	for i, w := range want {
		if w != out[i] {
			t.Errorf("sortHands2() == %s, wanted %s", out[i], w)
		}
	}
}
