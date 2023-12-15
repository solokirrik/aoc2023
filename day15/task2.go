package main

import (
	"bytes"
	"strconv"
)

var (
	EQUAL = []byte("=")
	MINUS = []byte("-")
)

type System []Box

func newLense(lense []byte) *Lense {
	sep := EQUAL
	if bytes.Contains(lense, MINUS) {
		sep = MINUS
	}

	parts := bytes.Split(lense, sep)
	hashVal := hash(parts[0])
	focalLength := 0

	if bytes.Equal(sep, EQUAL) {
		focalLength, _ = strconv.Atoi(string(parts[1]))
	}

	return &Lense{
		hash:        hashVal,
		sep:         sep,
		label:       string(parts[0]),
		focalLength: focalLength,
	}
}

func (s System) focusingPower() int {
	out := 0
	for i := range s {
		out += s[i].focusingPower()
	}

	return out
}

type Box struct {
	head *Lense
}

type Lense struct {
	hash        int
	label       string
	sep         []byte
	focalLength int
	prev        *Lense
	next        *Lense
}

func (b *Box) addLense(l *Lense) {
	switch {
	case bytes.Equal(l.sep, MINUS):
		b.minus(l)
	case bytes.Equal(l.sep, EQUAL):
		b.upsert(l)
	}
}

func (b *Box) minus(l *Lense) {
	// replace all the heading lenses with the same label
	if b.head == nil {
		return
	}

	for {
		if b.head == nil {
			return
		}
		if b.head.label != l.label {
			break
		}
		if b.head.label == l.label {
			if b.head.next != nil {
				b.head.next.prev = nil
			}
			b.head = b.head.next
		}
	}

	curLense := b.head.next
	for curLense != nil {
		if curLense.label == l.label {
			if curLense.next == nil {
				curLense.prev.next = nil
				return
			}

			curLense.prev.next = curLense.next
			curLense.next.prev = curLense.prev
		}
		curLense = curLense.next
	}
}

func (b *Box) upsert(l *Lense) {
	if b.head == nil {
		b.head = l
		return
	}

	if b.head.label == l.label {
		b.head.focalLength = l.focalLength
		return
	}

	var prevLense *Lense

	curLense := b.head
	for curLense != nil {
		if curLense.label == l.label {
			if curLense.next != nil {
				curLense.next.prev = l
				l.next = curLense.next
			}
			if curLense.prev != nil {
				curLense.prev.next = l
				l.prev = curLense.prev
			} else {
				b.head = l
			}
			return
		}
		prevLense = curLense
		curLense = curLense.next
	}

	prevLense.next = l
	l.prev = prevLense
}

func (b *Box) focusingPower() int {
	fp := 0

	slot := 1
	curLense := b.head
	for curLense != nil {
		fp += (b.head.hash + 1) * slot * curLense.focalLength
		slot++
		curLense = curLense.next
	}

	return fp
}
