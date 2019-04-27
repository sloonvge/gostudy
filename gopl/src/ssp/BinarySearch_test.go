package ssp

import "testing"

/** 
* Created by wanjx in 2019/3/19 11:58
**/
 
func TestBinarySearch(t *testing.T) {
	a := []int{ 1, 4, 5, 6, 7}
	t.Logf("out: %d", BinarySearch(a, 0))
}

func BinarySearch(a []int, t int) int {
	lo := 0
	hi := len(a) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		if a[mid] == t {
			return mid
		} else if a[mid] > t {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}