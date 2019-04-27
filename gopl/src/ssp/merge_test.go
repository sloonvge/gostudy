package ssp

import "testing"

/** 
* Created by wanjx in 2019/3/18 21:15
**/

func TestMerge(t *testing.T) {
	a := []int{4, 5, 6, 1, 2, 3}
	Merge(a, 0, 2, len(a))
	t.Logf("Merge: %v", a)
}
func Merge(a []int, lo int, mid int, hi int) {
	b := make([]int, 0)
	copy(b, a)
	i := lo
	j := mid + 1

	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = b[j]
			j++
		} else if j > hi {
			a[k] = a[i]
			i++
		} else if a[i] <= a[j] {
			a[k] = a[i]
			i++
		} else {
			a[k] = a[j]
			j++
		}
	}
}
// 匿名函数
/*
变量的生命周期不由它的作用域决定
 */

func square() func() int{
	var x int
	return func() int {
		x++
		return x*x
	}
}

func TestFuncZuoyy(t *testing.T) {
	f := square() // 函数值不仅仅是一行代码，也记录了状态
	t.Logf("%d\t%d", f(), f())
}