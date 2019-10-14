package container

type InitHeap []int

var N int

func MaxPQ(maxN int) InitHeap{
	return make(InitHeap, maxN + 1)
}

func (h *InitHeap) Insert(x int) {
	N++
	(*h)[N] = x
	h.Swim(N)
}

func (h *InitHeap) DelMax() (x int){
	x = (*h)[1]
	h.Swap(1, N)
	N--
	(*h)[N+1] = 0
	h.Sink(1)
	return
}

func (h InitHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h InitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h InitHeap) Swim(k int)  {
	for k > 1 && h.Less(k/2, k) {
		h.Swap(k, k/2)
		k = k/2
	}
}

func (h InitHeap) Sink(k int) {
	for 2*k <= N {
		j := 2*k
		if j < N && h.Less(j, j+1) {
			j++
		}
		if !h.Less(k, j) {
			break
		}
		h.Swap(k, j)
		k = j
	}
}


