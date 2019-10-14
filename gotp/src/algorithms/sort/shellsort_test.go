package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

// func ShellSort(a []int) []int {
// 	n := len(a)
// 	h := 1
// 	for h < n/3 {
// 		h = 3*h +1 // 1 4 13 40 121
// 	}
// 	for h >= 1 {
// 		for i := h; i < n; i++ {
// 			for j := i; j >= h && a[j] < a[j-h];j -= h {
// 				a[j], a[j-h] = a[j-h], a[j]
// 			}
// 		}
// 		h /= 3
// 	}
//
// 	return a
// }

func ShellSort(a []int) []int {
	n := len(a)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
		h /= 3
	}

	return a
}

func TestShellSort(t *testing.T) {
	N := 5
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Printf("orig: %v\n", a)
	a = ShellSort(a)
	fmt.Printf("sort: %v\n", a)
	sort.Ints(a)
	fmt.Println("base: ", a)
}
