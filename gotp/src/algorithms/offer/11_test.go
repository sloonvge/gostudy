package offer

import (
	"fmt"
	"math/rand"
	"mygo/algorithms/sort"
	"testing"
)

func QuickSort(a []int) {
	quickSort(a, 0, len(a) - 1)

}

func quickSort(a []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	k := partition(a, lo, hi)
	quickSort(a, lo, k - 1)
	quickSort(a, k + 1, hi)
}

func partition(a []int, lo int, hi int) int {
	i := lo + 1
	j := hi
	k := a[lo]
	for {
		for ; a[i] <= k; i++ {
			if i == hi {
				break
			}
		}
		for a[j] >= k{
			j--
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}

func TestSort(t *testing.T) {
	N := 5
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Printf("origin:%v\n", a)
	t.Run("quick base", func(t *testing.T) {
		sort.MergeSort(a)
		fmt.Printf("sort:%v\n", a)
	})
	t.Run("quick", func(t *testing.T) {
		QuickSort(a)
		fmt.Printf("sort:%v\n", a)
	})

}
