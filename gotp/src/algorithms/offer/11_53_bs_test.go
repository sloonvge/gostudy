package offer

import (
	"testing"
)

func BinarySearch(a []int, x int) (b bool) {
	if len(a) == 0 {
		return
	}
	i := 0
	j := len(a) - 1

	for i <= j {
		h := (i + j) / 2
		if a[h] == x {
			b = true
			return
		} else if a[h] < x {
			i = h + 1
		} else {
			j = h - 1
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

// 11
func RotateArray(a []int) int {
	if len(a) == 0 {
		panic("invalid parameters")
	}

	i := 0
	j := len(a) - 1
	h := i
	for a[i] >= a[j] {
		if (j - i) == 1 {
			h = j
			break
		}
		h = (i + j) / 2
		if a[i] == a[j] && a[i] == a[h] {
			return inOrder(a, i, j)
		}
		if a[h] <= a[j] {
			j = h
		} else if a[h] >= a[i] {
			i = h
		}
	}
	return a[h]
}
func inOrder(a []int, i, j int) (b int) {
	b = a[i]
	for k := i + 1; k <= j; k++ {
		if a[k] < b {
			b = a[k]
		}
	}
	return
}

func TestRotateArray(t *testing.T) {
	inputs := [][]int{
		{2, 3, 4, 0, 1},
		{0, 1, 2},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1},
	}
	wants := [][]int{
		{0},
		{0},
		{0},
		{0},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := RotateArray(inputs[i]); got != wants[i][j] {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}

// 53 p263
func GetNumberOfK(a []int, k int) int {
	if len(a) == 0 {
		panic("invalid parameters")
	}
	i := getFirstK(a, 0, len(a) - 1, k)
	j := getLastK(a, 0, len(a) - 1, k)
	if i > -1 && j > -1 {
		return j - i + 1
	}
	return 0
}
func getFirstK(a []int, i, j, k int) int {
	h := 0
	for i <= j {
		h = (i + j) / 2
		if a[h] < k {
			i = h + 1
		} else if a[h] > k {
			j = h - 1
		} else {
			if h == 0 || a[h - 1] != k{
				return h
			} else {
				j = h - 1
			}
		}
	}
	return -1
}
func getLastK(a []int, i, j, k int) int {
	h := 0
	for i <= j {
		h = (i + j) / 2
		if a[h] < k {
			i = h + 1
		} else if a[h] > k {
			j = h - 1
		} else {
			if h == len(a) -1  || a[h + 1] != k{
				return h
			} else {
				i = h + 1
			}
		}
	}
	return -1
}

func TestGetNumberOfK(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 3, 4, 5},
		{1},
		{1, 1},
	}
	helper := [][]int{
		{1, 3, 4, 5},
		{1, 2},
		{1, 2},
	}
	wants := [][]int{
		{1, 2, 1, 1},
		{1, 0},
		{2, 0},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := GetNumberOfK(inputs[i],
					helper[i][j]); got != wants[i][j] {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}

func GetMissingNumber(a []int) int {
	i := 0
	j := len(a) - 1
	h := 0
	for i <= j {
		h = (i + j) / 2
		if a[h] == h {
			i = h + 1
		} else {
			if h == 0 || a[h - 1] == h - 1{
				return h
			} else {
				j = h - 1
			}
		}
	}
	return -1
}

func TestGetMissingNumber(t *testing.T) {
	inputs := [][]int{
		{0, 1, 2, 4},
		{0, 1, 3},
		{0, 1, 2},
		{1, 2, 3},
		{0},
	}
	wants := [][]int{
		{3},
		{2},
		{-1},
		{0},
		{-1},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := GetMissingNumber(inputs[i]); got != wants[i][j] {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}

func GetNumberSameAsIndex(a []int) int {
	if len(a) == 0 {
		panic("invalid parameters")
	}
	i := 0
	j := len(a) - 1
	h := 0
	for i <= j {
		h = (i + j) / 2
		if a[h] < h {
			i = h + 1
		} else if a[h] > h {
			j = h - 1
		} else {
			return h
		}
	}
	return -1
}

func TestGetNumberSameAsIndex(t *testing.T) {
	inputs := [][]int{
		{-2, -1, 1, 3, 5},
		{-2, -1, 1},
	}
	wants := [][]int{
		{3},
		{-1},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := GetNumberSameAsIndex(inputs[i]); got != wants[i][j] {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}
