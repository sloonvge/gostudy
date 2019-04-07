package sort

import (
	"testing"
	"sort"
	"math/rand"
	"fmt"
)

//
// 快速排序
//
// 1. 原地排序，需要很小的一个辅助栈。 2. 复杂度O(NlgN)
// 3. 是一种分治的算法
// 4. 随机选取一个切分元素，使得切分元素左边的元素都不大于该切分元素；右边的都不大于该切分元素。

func QuickSort(a []int) {
	quickSort(a, 0, len(a) - 1)
}

func quickSort(a []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	j := partition(a, lo, hi)
	fmt.Printf("%v\n", a)
	quickSort(a, lo, j - 1)
	fmt.Printf("left:%v %d %d\n", a, lo, j - 1)
	quickSort(a, j + 1, hi)
	fmt.Printf("right:%v %d %d len:%d\n", a, j + 1, hi, len(a))
}

func partition(a []int, lo int, hi int) int {
	i := lo + 1
	j := hi

	v := a[lo]
	for  {
		fmt.Printf("+++%d %d\n", i, a[i])
		for ;a[i] <= v; i++ {

			if i == hi {
				break
			}
		}

		for a[j] >= v  {
			j--
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}

	a[lo], a[j] = a[j], a[lo]
	return j
}

func TestSortInts(t *testing.T) {
	sort.Ints([]int{2, 3, 1})
}

func TestQuickSort(t *testing.T) {
	N := 5
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Printf("origin:%v\n", a)
	QuickSort(a)
	fmt.Printf("sort:%v\n", a)
}