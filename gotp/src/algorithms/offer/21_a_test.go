package offer

import (
	"reflect"
	"testing"
)

// 21
func ReorderOddEven(a []int) []int {
	var i, j int
	i = 0
	j = len(a) - 1
	for i < j {
		for i < j && a[i] & 0x1 == 1 {
			i++
		}
		for i < j && a[j] & 0x1 == 0 {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	return a
}

func TestReorderOddEven(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4, 5},
		{1, 3, 5},
		{2, 4, 6},
		{},
		{2, 4, 1, 3, 5},
		{1, 3, 5, 2, 4},
	}
	wants := [][]int{
		{1, 5, 3, 4, 2},
		{1, 3, 5},
		{2, 4, 6},
		{},
		{5, 3, 1, 4, 2},
		{1, 3, 5, 2, 4},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			if got := ReorderOddEven(inputs[i]); !reflect.DeepEqual(got, wants[i]){
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
}

// 39
func MoreThanHalfNum(a []int, n int) int{
	if len(a) == 0 {
		return -1
	}
	result := a[0]
	times := 1
	for i := 0; i < n; i++ {
		if times == 0 {
			result = a[i]
			times = 1
		} else if a[i] == result{
			times++
		} else {
				times--
		}
	}
	return result
}

func MoreThanHalfNum2(a []int) int {
	n := len(a)
	start := 0
	end := n - 1
	mid := n >> 1
	index := partition(a, n, start, end)
	for index != mid {
		if index > mid {
			end = index - 1
			index = partition(a, n, start, end)
		} else if index < mid {
			start = index + 1
			index = partition(a, n, start, end)
		}
	}

	return a[mid]
}
func partition(a []int, n, start, end int) int{
	if len(a) == 0 || start < 0 || end > n {
		panic("invalid parameters")
	}

	index := 0
	a[index], a[end] = a[end], a[index]
	small := start - 1
	for index = start; index < end; index++ {
		if a[index] < a[end] {
			small++
			if small != index {
				a[index], a[small] = a[small], a[index]
			}
		}
	}
	small++
	a[small], a[end] = a[end], a[small]

	return small
}

func TestMoreThanHalfNum(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 2, 2},
	}
	wants := []int{
		2,
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			if got := MoreThanHalfNum(inputs[i], len(inputs[i])); got != wants[i]{
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
	t.Run("2", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			if got := MoreThanHalfNum2(inputs[i]); got != wants[i]{
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
}

// 40
func GetLeastNumbers(a []int, k int) []int {
	lo := 0
	hi := len(a) - 1
	index := partition(a, len(a), lo, hi)

	for index != k - 1 {
		if index > k - 1 {
			hi = index - 1
			index = partition(a, len(a), lo, hi)
		} else {
			lo = index + 1
			index = partition(a, len(a), lo, hi)
		}
	}

	return a[:k]
}

func GetLeastNumbers2(a []int, k int) []int {
	lo := 0
	hi := len(a) - 1
	index := partition(a, len(a), lo, hi)

	for index != k - 1 {
		if index > k - 1 {
			hi = index - 1
			index = partition(a, len(a), lo, hi)
		} else {
			lo = index + 1
			index = partition(a, len(a), lo, hi)
		}
	}

	return a[:k]
}

func TestGetLeastNumbers(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 2, 2},
	}
	wants := [][]int{
		{1, 2},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			if got := GetLeastNumbers(inputs[i], 2); reflect.DeepEqual(got, wants[i]){
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
	// t.Run("2", func(t *testing.T) {
	// 	for i := 0; i < len(inputs); i++ {
	// 		if got := GetLeastNumbers2(inputs[i]); got != wants[i]{
	// 			t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
	// 		}
	// 	}
	// })
}
