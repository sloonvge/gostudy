package gostudy

import "testing"

/** 
* Created by wanjx in 2019/4/6 23:28
**/
 
func TestMergeSort(t *testing.T) {
	a := []int{4, 5, 6, 1, 2}
	Show(a)
	MergeSort(a)
	Show(a)
}


