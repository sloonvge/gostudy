package algorithms

import "testing"

/*
堆
 */
type Heap []int

func NewHeap() (h Heap) {
	h = Heap{}
	h = append(h, -1)
	return
}


func (h *Heap) Insert(v int) {
	*h = append(*h, v)
	h.swim()
}

func (h Heap) swim() {
	N := len(h)
	i := N - 1
	for {
		if i > 1 && (h)[i] > (h)[ i / 2] {
			h.swap(i ,i / 2)
			i = i / 2
		} else {
			break
		}
	}
}

func (h Heap) swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

//func (h Heap) max() int {
//
//}

func (h *Heap) DelMax() int {
	m := (*h)[1]
	h.swap(1, len(*h) - 1)
	*h = (*h)[:len(*h) - 1]
	h.sink()
	return m
}

func (h Heap) sink() {
	s := 1
	for {
		if h[s] < h[2 * s] && s < len(h) {
			h.swap(s, 2 * s)
			s = 2 * s
		} else {
			break
		}
	}
}

func (h Heap) IsEmpty() bool{
	return len(h) == 0
}

func (h Heap) Size() int {
	return len(h)
}

func TestHeap(t *testing.T) {
	s := []int{9, 10, 56, 1, 43, 298, 100}
	h := NewHeap()
	for _, v := range s {
		h.Insert(v)
	}
	t.Logf("%v", h)
}