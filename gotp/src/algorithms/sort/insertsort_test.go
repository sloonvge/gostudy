package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func InsertSort(a []int) []int {
	n := len(a)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if a[i] < a[j] {
				a[j], a[i] = a[i], a[j]
			}
		}
		fmt.Println("...", a)
	}

	return a
}

func TestInsertSort(t *testing.T) {
	N := 5
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(100)
	}
	a[1] = 1

	fmt.Printf("orig: %v\n", a)
	a = InsertSort(a)
	fmt.Printf("sort: %v\n", a)
	sort.Ints(a)
	fmt.Println("base:", a)
}
