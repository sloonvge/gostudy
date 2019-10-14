package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func HeapSort(a []int) []int {
	n := len(a) - 1
	for k := n/2; k >= 1; k-- {
		sink(a, k, n)
	}
	for n > 1 {
		a[1], a[n] = a[n], a[1]
		n--
		sink(a, 1, n)
	}

	return a
}

func sink(a []int, i, j int) {
	for 2*i <= j {
		k := 2*i
		if k < j && a[k] < a[k+1] {
			k++
		}
		if a[i] > a[k] {
			break
		}
		i = k
	}
}

func TestHeapSort(t *testing.T) {
	N := 5
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(100)
	}
	a[0] = 0

	fmt.Printf("orig: %v\n", a)
	a = HeapSort(a)
	fmt.Printf("sort: %v\n", a)
	sort.Ints(a)
	fmt.Println("base:", a)
}