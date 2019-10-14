package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func SelectSort(a []int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
		fmt.Println("...", a)
	}

	return a
}

func TestSelectSort(t *testing.T) {
	N := 5
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Printf("orig: %v\n", a)
	a = SelectSort(a)
	fmt.Printf("sort: %v\n", a)
	sort.Ints(a)
	fmt.Println("base:", a)
}
