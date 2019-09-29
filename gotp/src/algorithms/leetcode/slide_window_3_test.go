package leetcode

import (
	"fmt"
	"math"
	"testing"
)

// 3 滑动窗口
func lengthOfLongestSubstring1(s string) int {
	n := len(s)
	max := 0
	isRepeat := func(s string) bool {
		m := make(map[byte]int)
		for i := 0; i < len(s); i++ {
		m[s[i]]++
		if m[s[i]] > 1 {
		return true
	}
	}
		return false
	}
	for i := n; i > 0; i-- {
		for j := 0; j < i; j++ {
			if !isRepeat(s[j:i]) {
				if i - j > max {
					max= i - j
				}
			}
		}
	}

	return max
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	window := make(map[byte]int, 0)
	n := len(s)
	i := 0
	j := 0
	max := 0
	for j < n {
		if _, ok := window[s[j]]; ok {
			delete(window, s[i])
			i++
		} else {
			window[s[j]] = j
			j++
			if j - i > max {
				max = j - i
			}
		}
	}


	return max
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	window := make(map[byte]int, 0)
	n := len(s)
	max := 0
	for i, j := 0, 0; j < n; j++ {
		if index, ok := window[s[j]]; ok {
			if index > i {
				i = index
			}
		}
		window[s[j]] = j + 1
		if j - i + 1> max {
			max = j - i + 1
		}
	}


	return max
}

// 239
func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	ans := make([]int, 0)
	i := 0
	j := k
	for j <= len(nums){
		max := nums[i]
		for l := i; l < j; l++ {
			if nums[l] > max {
				max = nums[l]
			}
		}
		ans = append(ans, max)
		i++
		j++
	}

	return ans
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	ans := make([]int, 0)
	queue := make([]int, 0)
	for l := 0; l < k; l++ {
		for len(queue) > 0 && nums[l] > nums[queue[len(queue)-1]]{
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, l)
	}
	for j := k - 1; j < len(nums); {
		if j - k + 1 > queue[0] {
			queue = queue[1:]
		}
		ans = append(ans, nums[queue[0]])
		j++
		for j < len(nums) && len(queue) > 0 && nums[j] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, j)
	}

	return ans
}

// 76
func minWindow1(s string, t string) string {
	if s == "" {
		return ""
	}
	n := len(s)
	ans := ""
	min := math.MaxInt64
	contains := func(s, t string) bool {
		set := make(map[byte]int)
		for i := 0; i < len(t); i++ {
			set[t[i]]++
		}
		for i := 0; i < len(s); i++ {
			if _, ok := set[s[i]]; ok {
				set[s[i]]--
				if set[s[i]] == 0 {
					delete(set, s[i])
				}
			}
		}

		return len(set) == 0
	}
	for i := n; i > 0; i-- {
		for j := 0; j < i; j++ {
			if contains(s[j:i], t) {
				if i - j < min {
					ans = s[j:i]
					min = i - j
				}
			}
		}
	}
	return ans
}
// 错误答案
func minWindow2(s string, t string) string {
	seen := make(map[byte]struct{})
	for i := range t {
		seen[t[i]] = struct{}{}
	}
	set := make(map[byte]int)
	min := math.MaxInt64
	ans := ""
	i, j := 0, 0
	for j < len(s) {
		if _, ok := seen[s[j]]; ok {
			if index, ok2 := set[s[j]]; !ok2 {
				set[s[j]] = j
			} else {
				for ; i <= index; i++ {
					if _, ok := set[s[i]]; ok {
						delete(set, s[i])
					}
				}
				for ; i < len(s); i++ {
					if _, ok := set[s[i]]; ok {
						break
					}
				}
			}
			set[s[j]] = j
			j++
			if len(seen) == len(set) {
				if j - i < min {
					ans = s[i:j]
					min = j - i
				}
			}
		} else {
			j++
		}
	}

	return ans
}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(s) < len(t) || len(t) == 0 {
		return ""
	}
	set := make(map[byte]int)
	for i := range t {
		set[t[i]]++
	}

	count := 0
	min := math.MaxInt64
	minLeft := 0
	i, j := 0, 0
	for j < len(s) {
		if _, ok := set[s[j]]; ok {
			set[s[j]]--
			if set[s[j]] >= 0 {
				count++
			}
			for count == len(t) {
				if j - i + 1 < min {
					minLeft = i
					min = j - i + 1
				}
				if _, ok = set[s[i]]; ok {
					set[s[i]]++
					if set[s[i]] > 0 {
						count--
					}
				}
				i++
			}
		}
		j++
	}
	if min > len(s) {
		return ""
	}

	return s[minLeft:minLeft+min]
}

func TestWindow(t *testing.T) {
	fmt.Println(minWindow("a", "b"))
}

