package sort

import "testing"

/** 
* Created by wanjx in 2019/4/6 23:28
**/
 
func TestMergeSort(t *testing.T) {
	a := []int{4, 5, 6,1, 2, 3, 7, 10}
	Show(a)
	MergeSort(a)
	Show(a)
}

func TestMergeSortB2U(t *testing.T) {
	a := []int{4, 5, 6,1, 2, 3, 7, 10}
	Show(a)
	MergeSortB2U(a)
	Show(a)
}


