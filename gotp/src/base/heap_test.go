package base

import (
	"testing"
)

type Item struct {
	value int
	priority int
	index int
}

type PQ []*Item

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PQ) Swim(i int) {
	for i > 1 && pq.Less(i/2, i) {
		pq.Swap(i/2, i)
		i = i/2
	}
}

func (pq PQ) Sink(i int) {
	for 2*i <= pq.Len() {
		j := 2*i
		if j < pq.Len() && pq.Less(j, j+1) {
			j++
		}
		if !pq.Less(i, j) {
			break
		}
		pq.Swap(i, j)
		i = j
	}
}

func (pq *PQ) Insert(v int) {
	*pq = append(*pq, &Item{
		value: v,
		index: pq.Len()+1,
	})
	pq.Swim(pq.Len())
}

func (pq *PQ) Pop() int {
	max := (*pq)[1]
	pq.Swap(1, pq.Len()-1)
	*pq = (*pq)[:pq.Len()-1]
	pq.Sink(1)
	return max.value

}

func TestPriorityQueue(t *testing.T) {
	t.Run("base", func(t *testing.T) {
	})
	t.Run("demo", func(t *testing.T) {
		
	})
}

