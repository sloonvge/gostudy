package offer

import (
	"fmt"
	"testing"
)

// 49
func isUgly(n int) bool {
	for n % 2 == 0 {
		n /= 2
	}
	for n % 3 == 0 {
		n /= 3
	}
	for n % 5 == 0 {
		n /= 5
	}
	if n == 1 {
		return true
	}
	return false
}

func GetUglyNumber(index int) int{
	if index <= 0 {
		return 0
	}
	n := 1
	i := 0
	for {
		if isUgly(n) {
			i++
		}
		if i == index {
			break
		}
		n++
	}

	return n
}

func GetUglyNumber2(index int) int{
	if index <= 0 {
		return 0
	}
	nums := make([]int, index)
	nums[0] = 1
	next := 1
	m2 := 0
	m3 := 0
	m5 := 0

	for next < index {
		minUgly := min(nums[m2] * 2, nums[m3] * 3, nums[m5] * 5)
		nums[next] = minUgly
		for nums[m2] * 2 <= minUgly {
			m2++
		}
		for nums[m3] * 3 <= minUgly {
			m3++
		}
		for nums[m5] * 5 <= minUgly {
			m5++
		}
		next++
	}

	return nums[next - 1]
}
func min(n1, n2, n3 int) int{
	var m int
	if n1 < n2 {
		m = n1
	} else {
		m = n2
	}
	if m < n3 {
		return m
	}
	return n3
}

func TestGetUglyNumber(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		fmt.Println(isUgly(8))
	})
	t.Run("1", func(t *testing.T) {
		fmt.Println(GetUglyNumber(1000))
	})
	t.Run("2", func(t *testing.T) {
		fmt.Println(GetUglyNumber2(1000))
	})
}
