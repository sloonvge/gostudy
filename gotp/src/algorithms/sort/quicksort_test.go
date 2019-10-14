package main

import (
	"fmt"
	"sort"
	"testing"
)

/**
* Created by wanjx in 2019/4/27 23:23
**/
 
// func QuickSort(a []int) {
// 	quickSort(a, 0, len(a) - 1)
// }
//
// func quickSort(a []int, lo int, hi int) {
// 	if lo >= hi {
// 		return
// 	}
//
// 	j := partition(a, lo, hi)
// 	quickSort(a, lo, j - 1)
// 	quickSort(a, j + 1, hi)
// }
//
// func partition(a []int, lo int, hi int) int {
// 	i := lo + 1
// 	j := hi
//
// 	v := a[lo]
// 	for {
// 		for ; a[i] < v; i++ {
// 			if i == hi {
// 				break
// 			}
// 		}
//
// 		for ; a[j] > v; j-- {
// 			if j == lo {
// 				break
// 			}
// 		}
//
// 		if i >= j {
// 			break
// 		}
//
// 		a[i], a[j] = a[j], a[i]
// 	}
//
// 	a[lo], a[i] = a[i], a[lo]
// 	return i
// }

// func QuickSort(a []int) {
// 	quickSort(a , 0, len(a) - 1)
// }
// func quickSort(a []int, lo, hi int) {
// 	if lo >= hi {
// 		return
// 	}
// 	s := partition(a, lo, hi)
// 	quickSort(a, lo, s - 1)
// 	quickSort(a, s + 1, hi)
// }
// func partition(a []int, lo, hi int) int {
// 	i := lo + 1
// 	j := hi
// 	v := a[lo]
// 	for {
// 		for a[i] <= v {
// 			if i == hi {
// 				break
// 			}
// 			i++
// 		}
// 		for a[j] >= v {
// 			if j == lo {
// 				break
// 			}
// 			j--
// 		}
// 		if i >= j {
// 			break
// 		}
// 		a[i], a[j] = a[j], a[i]
// 	}
//
// 	a[lo], a[j] = a[j], a[lo]
// 	return j
// }

// func QuickSort(a []int) {
// 	if len(a) == 0 {
// 		return
// 	}
// 	quickSort(a, 0, len(a) - 1)
// }
// func quickSort(a []int, i, j int) {
// 	if i >= j {
// 		return
// 	}
// 	k := partion(a, i, j)
// 	quickSort(a, 0, k - 1)
// 	quickSort(a, k + 1, j)
// }
// func partion(a []int, start, end int) int {
// 	k := a[start]
// 	var i, j int
// 	i = start + 1
// 	j = end
// 	for {
// 		for ; i < end && a[i] <= k; i++{
// 		}
// 		for ; j > start && a[j] >= k; j--{
// 		}
// 		if i >= j {
// 			break
// 		}
// 		a[i], a[j] = a[j], a[i]
// 		i++
// 		j--
// 	}
// 	a[start], a[j] = a[j], a[start]
//
// 	return j
// }

func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	quickSort(a, 0, len(a) - 1)
	return a
}
func quickSort(a []int, i, j int) {
	if i >= j {
		return
	}
	p := partition(a, i, j)
	quickSort(a, i, p - 1)
	quickSort(a, p + 1, j)
}
func partition(a []int, lo, hi int) int {
	k := a[lo]
	i := lo + 1
	j := hi
	for {
		for ; i < hi && a[i] <= k; i++ {
		}
		for ; j > lo && a[j] >= k; j-- {
		}
		if i >= j { break }
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	a[j], a[lo] = a[lo], a[j]

	return j
}

func TestQuickSort(t *testing.T) {
	// N := 5
	// a := make([]int, N)
	// for i := 0; i < N; i++ {
	// 	a[i] = rand.Intn(100)
	// }
	a := []int{1,1,2,2,1}

	fmt.Printf("origin:%v\n", a)
	QuickSort(a)
	fmt.Printf("sort:%v\n", a)
	sort.Ints(a)
}