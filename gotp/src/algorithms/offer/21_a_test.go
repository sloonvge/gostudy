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
