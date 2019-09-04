package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

/** 
* Created by wanjx in 2019/9/4 23:27
**/
 
func MergeSort(a []int) []int{
	n := len(a)
	if n <= 1 {
		return a
	}
	b := make([]int, n / 2)
	c := make([]int, n - n / 2)
	copy(b, a[:len(b)])
	copy(c, a[len(b):])
	return merge(MergeSort(b), MergeSort(c))
}

func merge(a, b []int) []int {
	c := make([]int, len(a) + len(b))
	var i, j int
	for k := 0; k < len(c); k++ {
		if i >= len(a) {
			c[k] = b[j]
			j++
		} else if j >= len(b) {
			c[k] = a[i]
			i++
		} else if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
	}

	return c
}

func TestMergeSort(t *testing.T) {
	N := 5
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Printf("origin:%v\n", a)
	a = MergeSort(a)
	fmt.Printf("sort:%v\n", a)
	sort.Ints(a)
	t.Run("", func(t *testing.T) {
		a := []int{0, 1, 2, 3, 4}
		b := []int{5, 6}
		copy(b, a[:2])
		fmt.Println(b)
		a[0] = -1
		fmt.Println(b)
	})
}