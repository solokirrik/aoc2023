package main

type Node struct {
	name      string
	Childrens []*Node
	checks    []check
}

func Task2(rulesLines [][]byte) int {
	rules := make(map[string]rule)

	for _, line := range rulesLines {
		r := paseRule(line)
		rules[r.name] = r
	}

	rating := 0
	return rating
}

/*

	1 <= X <= 4000
	2090 < M <= 4000
	2006 <= A <= 4000
	1 <= S <= 1351
	9 781 240 000

	1 <= X <= 1416
	1 <= M <= 4000
	1 <= A < 2006
	1 <= S < 1351
	15 331 032 000 000

	2662 < X <= 4000
	1 <= M <= 4000
	1 <= A < 2006
	1 <= S < 1351
	14 486 526 000 000

	1 <= X <= 4000
	1 <= M <= 4000
	1 <= A < 2006
	3448 <= S <= 4000
	17 740 240 000 000

	1 <= X <= 4000
	1 <= M <= 4000
	1 <= A <= 2006
	1351 <= S < 3448
	67 305 312 000 000

	1 <= X <= 4000
	838 < M < 1801
	1 <= A <= 4000
	1351 <= S <= 4000
	40 746 400 000 000

	1 <= X <= 4000
	1 <= M <= 838
	1 <= A <= 1716
	1351 <= S <= 4000
	15 242 884 800 000



	->167 409 079 868 000
	!!170 862 176 040 000
*/
