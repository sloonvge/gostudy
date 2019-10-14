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
 
// func MergeSort(a []int) []int{
// 	n := len(a)
// 	if n <= 1 {
// 		return a
// 	}
// 	b := make([]int, n / 2)
// 	c := make([]int, n - n / 2)
// 	copy(b, a[:len(b)])
// 	copy(c, a[len(b):])
// 	return merge(MergeSort(b), MergeSort(c))
// }
//
// func merge(a, b []int) []int {
// 	c := make([]int, len(a) + len(b))
// 	var i, j int
// 	for k := 0; k < len(c); k++ {
// 		if i >= len(a) {
// 			c[k] = b[j]
// 			j++
// 		} else if j >= len(b) {
// 			c[k] = a[i]
// 			i++
// 		} else if a[i] < b[j] {
// 			c[k] = a[i]
// 			i++
// 		} else {
// 			c[k] = b[j]
// 			j++
// 		}
// 	}
//
// 	return c
// }

// func MergeSort(a []int) []int{
// 	if len(a) <= 1 {
// 		return a
// 	}
// 	n := len(a)
// 	b := make([]int, n >> 1)
// 	c := make([]int, n-n >> 1)
//
// 	copy(b, a[:n>>1])
// 	copy(c, a[n>>1:])
// 	return merge(MergeSort(b), MergeSort(c))
// }
// func merge(b, c []int) []int{
// 	a := make([]int, len(b) + len(c))
// 	var i, j int
// 	for k := 0; k < len(a); k++ {
// 		if i >= len(b) {
// 			a[k] = c[j]
// 			j++
// 		} else if j >= len(c) {
// 			a[k] = b[i]
// 			i++
// 		} else if b[i] < c[j] {
// 			a[k] = b[i]
// 			i++
// 		} else {
// 			a[k] = c[j]
// 			j++
// 		}
// 	}
// 	return a
// }

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	n := len(a)
	b := make([]int, n >> 1)
	c := make([]int, n - n >> 1)
	copy(b, a[:n >> 1])
	copy(c, a[n >> 1:])
	return merge(MergeSort(b), MergeSort(c))
}
func merge(b, c []int) (a []int) {
	a = make([]int, len(b) + len(c))
	j := 0
	k := 0
	for i := 0; i < len(a); i++ {
		if j >= len(b) {
			a[i] = c[k]
			k++
		} else if k >= len(c) {
			a[i] = b[j]
			j++
		} else if b[j] < c[k] {
			a[i] = b[j]
			j++
		} else {
			a[i] = c[k]
			k++
		}
	}

	return
}

func TestMergeSort(t *testing.T) {
	N := 5
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Printf("orig: %v\n", a)
	a = MergeSort(a)
	fmt.Printf("sort: %v\n", a)
	sort.Ints(a)
	fmt.Println("base:", a)
}