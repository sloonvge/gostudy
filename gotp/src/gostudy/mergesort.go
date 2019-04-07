package gostudy

import "fmt"

/**
* Created by wanjx in 2019/4/6 23:28
**/
 
func MergeSort(a []int) {
	lo := 0
	hi := len(a) - 1
	aux := make([]int, len(a))
	sort(a, aux, lo, hi)
}

func sort(a []int, aux[]int, lo int, hi int) {
	// if hi <= lo {
	// 	return
	// }
	mid := (lo + hi) / 2
	fmt.Println(mid)
	// sort(a, aux, lo, mid)
	// sort(a, aux, mid + 1, hi)
	// Show(a)
	merge(a, aux, lo, mid, hi)
}

func merge(a []int, aux []int, lo int, mid int, hi int) {
	fmt.Println(isSorted(a, lo, mid))
	fmt.Println(isSorted(a, mid+1, hi))
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++{
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if a[i] < a[j] {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}
}

func isSorted(a []int, lo int, hi int) bool {
	for i := lo; i <= hi - 1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}

	return true
}