// Adapted from the examples in Go Docs

package advent_helpers

import "container/heap"

type Record struct {
	Value    int
	Distance int // The priority of the item in the queue.

	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Record

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Record)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *Record, distance int) {
	item.Distance = distance
	heap.Fix(pq, item.index)
}
