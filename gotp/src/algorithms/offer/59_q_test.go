package offer

import (
	"fmt"
	"testing"
)

// 59
func MaxInWindows(a []int, k int) []int {
	ans := make([]int, 0)
	s := make([]int, 0)
	max := a[0]
	maxIndex := 0
	for i := 1; i < k; i++ {
		if a[i] > max {
			max = a[i]
			maxIndex = i
		}
	}
	s = append(s, maxIndex)
	ans = append(ans, a[s[0]])
	for i := 1; i < len(a) - k + 1; i++ {
		cur := a[i + k - 1]
		if i > s[0] {
			s = s[1:]
		}
		for len(s) != 0 && a[s[len(s)-1]] < cur {
			s = s[:len(s) - 1]
		}
		s = append(s, i + k - 1)
		ans = append(ans, a[s[0]])
	}

	return ans
}


func TestMaxInWindows(t *testing.T) {
	fmt.Println(MaxInWindows([]int{2,3,4,2,6,2,5,1}, 3))
}