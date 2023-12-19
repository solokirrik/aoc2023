package main

import (
	"bytes"
	"fmt"
	"strconv"
)

const ACCEPTED = "A"
const REJECTED = "R"

func Task1(inpParts [][]byte) int {
	partsLines := bytes.Split(inpParts[1], []byte("\n"))
	parts := make([]part, len(partsLines))
	for i, line := range partsLines {
		parts[i] = newPart(line)
	}

	rules := make(map[string]rule)
	rulesLines := bytes.Split(inpParts[0], []byte("\n"))

	for _, line := range rulesLines {
		r := paseRule(line)
		rules[r.name] = r
	}

	rating := 0
	for _, part := range parts {
		ruleName := "in"
	RULE:
		for {
			rule := rules[ruleName]
		CH:
			for _, check := range rule.checks {
				if check.fieldName == "" ||
					check.comp(part.fieldValue(check.fieldName), check.bVal) {
					switch check.then {
					case ACCEPTED:
						rating += part.rating()
						break RULE
					case REJECTED:
						break RULE
					default:
						ruleName = check.then
						break CH
					}
				}
			}
		}
	}

	return rating
}

func paseRule(inp []byte) rule {
	name, checksPart, _ := bytes.Cut(inp, []byte("{"))
	checksPart = bytes.TrimRight(checksPart, "}")
	checks := bytes.Split(checksPart, []byte(","))

	r := rule{
		name: string(name),
	}

	for _, ch := range checks {
		if !bytes.Contains(ch, []byte(":")) {
			r.checks = append(r.checks, check{then: string(ch)})
			continue
		}

		r.checks = append(r.checks, newChech(ch))
	}

	return r
}

type rule struct {
	name   string
	checks []check
}

type check struct {
	fieldName string
	comp      func(a int, b int) bool
	bVal      int
	then      string
}

func newChech(inp []byte) check {
	fieldComp, then, _ := bytes.Cut(inp, []byte(":"))
	sep := []byte("<")
	exec := isLess

	if bytes.Contains(fieldComp, []byte("<")) {
		sep = []byte("<")
		exec = isLess
	} else {
		sep = []byte(">")
		exec = isMore
	}

	field, bValStr, _ := bytes.Cut(fieldComp, sep)
	bVal, _ := strconv.Atoi(string(bValStr))
	return check{
		fieldName: string(field),
		comp:      exec,
		bVal:      bVal,
		then:      string(then),
	}
}

func isLess(a, b int) bool { return a < b }

func isMore(a, b int) bool { return a > b }

type part struct {
	x, m, a, s int
}

func (p part) rating() int {
	return p.x + p.m + p.a + p.s
}

func (p part) fieldValue(fieldName string) int {
	switch fieldName {
	case "x":
		return p.x
	case "m":
		return p.m
	case "a":
		return p.a
	case "s":
		return p.s
	}

	return -1000
}

func newPart(inp []byte) part {
	p := part{}
	_, err := fmt.Sscanf(string(inp), "{x=%d,m=%d,a=%d,s=%d}", &p.x, &p.m, &p.a, &p.s)
	if err != nil {
		panic(err)
	}

	return p
}
