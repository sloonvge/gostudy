package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func BubbleSort(a []int) []int {
	n := len(a)
	for i := n - 1; i >= 0; i-- {
		// flag := false
		for j := 0; j < i; j++ {
			if a[j] > a[j + 1] {
				// flag = true
				a[j + 1], a[j] = a[j], a[j + 1]
			}
		}
		// if !flag {
		// 	break
		// }
		fmt.Println("...", a)
	}

	return a
}

func TestBubbleSort(t *testing.T) {
	N := 5
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Printf("orig: %v\n", a)
	a = BubbleSort(a)
	fmt.Printf("sort: %v\n", a)
	sort.Ints(a)
	fmt.Println("base: ", a)
}
