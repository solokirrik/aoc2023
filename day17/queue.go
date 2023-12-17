package main

import (
	"fmt"
	"slices"
)

type Queue struct {
	q []Waiter
}

type Waiter struct {
	parent         Point
	pos            DirectedPoint
	heatLoss       int
	straightNumber int
}

func (q *Waiter) String() string {
	return fmt.Sprintf("Y:%d X:%d H:%d D:%d S:%d", q.pos[0], q.pos[1], q.heatLoss, q.pos[2], q.straightNumber)
}

func (q *Queue) Push(w []Waiter) {
	for i := range w {
		q.q = append(q.q, w[i])
	}

	slices.SortStableFunc[[]Waiter](q.q, func(a, b Waiter) int {
		return a.heatLoss - b.heatLoss
	})
}

func (q *Queue) Pop() Waiter {
	w := q.q[0]
	q.q = q.q[1:]
	return w
}

func (q *Queue) IsEmpty() bool {
	return len(q.q) == 0
}
