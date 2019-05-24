package main

import (
	"fmt"
	"math/rand"
	"testing"
)

/**
* Created by wanjx in 2019/4/27 23:23
**/
 
func QuickSort(a []int) {
	quickSort(a, 0, len(a) - 1)
}

func quickSort(a []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	j := partition(a, lo, hi)
	quickSort(a, lo, j - 1)
	quickSort(a, j + 1, hi)
}

func partition(a []int, lo int, hi int) int {
	i := lo + 1
	j := hi

	v := a[lo]
	for {
		for ; a[i] < v; i++ {
			if i == hi {
				break
			}
		}

		for ; a[j] > v; j-- {
			if j == lo {
				break
			}
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
	}

	a[lo], a[i] = a[i], a[lo]
	return i
}

func TestQuickSort(t *testing.T) {
	N := 5
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Printf("origin:%v\n", a)
	QuickSort(a)
	fmt.Printf("sort:%v\n", a)
}