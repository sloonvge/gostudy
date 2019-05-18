package sort


/**
* Created by wanjx in 2019/4/6 23:28
**/
 
func MergeSort(a []int) {
	lo := 0
	hi := len(a) - 1
	aux := make([]int, len(a))
	mergeSort(a, aux, lo, hi)
}


func MergeSortB2U(a []int) {
	N := len(a)
	aux := make([]int, N)
	for sz := 1; sz < N; sz += sz {
		for lo := 0; lo < N - sz; lo += sz * 2 {
			merge(a, aux, lo, lo + sz - 1, min(lo+sz*2-1, N))
		}
	}
}

func mergeSort(a []int, aux[]int, lo int, hi int) {
	 if hi <= lo {
	 	return
	 }
	mid := (lo + hi) / 2
	mergeSort(a, aux, lo, mid)
	mergeSort(a, aux, mid + 1, hi)

	merge(a, aux, lo, mid, hi)
}

func merge(a []int, aux []int, lo int, mid int, hi int) {
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
		} else if aux[i] < aux[j] {
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