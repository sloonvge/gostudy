package offer

import (
	"fmt"
	"sort"
	"testing"
)

// 61
func IsContinuous(numbers []int) bool {
	if len(numbers) == 0 {
		return false
	}
	sort.Ints(numbers)
	numberOf0 := 0
	numberOfGap := 0
	for i := 0; i < len(numbers) && numbers[i] == 0; i++ {
		numberOf0++
	}
	small := numberOf0
	big := small + 1
	for big < len(numbers) {
		if numbers[small] == numbers[big] {
			return false
		}
		numberOfGap += numbers[big] - numbers[small] - 1
		small, big = big, big + 1
	}

	if numberOfGap > numberOf0 {
		return false
	}
	return true
}


func TestIsContinuous(t *testing.T) {
	v := IsContinuous([]int{1, 2, 3, 4, 6})
	fmt.Println(v)
}