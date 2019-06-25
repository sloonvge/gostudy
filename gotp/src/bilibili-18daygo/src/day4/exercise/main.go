package main

import (
	"fmt"
	"math/rand"
)

var a []int

func init() {
	n := 10
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
	}
}

func bubbleSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func selectSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[k] {
				k = j
			}
		}
		a[i], a[k] = a[k], a[i]	
	}
}

func insertSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func quickSort(a []int) { 
	// quick(a, 0, len(a)-1)
	quick2(a, 0, len(a)-1)
}

func quick2(a []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	val := a[lo]
	k := lo
	for i := lo + 1; i <= hi; i++ {
		if a[i] < val {
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}
	a[k] = val
	quick2(a, lo, k-1)
	quick2(a, k+1, hi)

}

func quick(a []int, lo int, hi int) {
	if hi <= lo {
		return 
	}
	k := partition(a, lo, hi)
	quick(a, lo, k-1)
	quick(a, k+1, hi)
}

func partition(a []int, lo int, hi int) int {
	k := a[lo]

	i := lo+1
	j := hi
	for {
		for {
			if a[i] <= k {
				i++
			} else {
				break
			}
		}
		for {
			if a[j] > k {
				j--
			} else {
				break
			}
		}
		if (i >= j) {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}


func main() {
	fmt.Println("before: ", a)
	// bubbleSort(a)
	selectSort(a)
	// insertSort(a)
	// quickSort(a)
	fmt.Println("after: ", a)
}
 