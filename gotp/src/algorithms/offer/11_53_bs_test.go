package offer

import "testing"

func BinarySearch(a []int, b int) (isExist bool) {
	if len(a) == 0 {
		return
	}
	lo := 0
	hi := len(a) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		if a[mid] == b {
			isExist = true
			return
		} else if a[mid] < b {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return
}

func TestBinarySearch(t *testing.T) {
	inputs := [][]int{
		{0, 1, 2, 3, 4},
		{0},
	}
	helper := [][]int{
		{2, 5, 0, 4},
		{0, 1},
	}
	wants := [][]bool{
		{true, false, true, true},
		{true, false},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := BinarySearch(inputs[i],
					helper[i][j]); got != wants[i][j] {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}
