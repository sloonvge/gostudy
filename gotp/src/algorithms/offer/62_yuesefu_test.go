package offer

import (
	"container/list"
	"fmt"
	"testing"
)

func LastRemaining(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	}
	numbers := list.New()
	for i := 0; i < n; i++ {
		numbers.PushBack(i)
	}
	current := numbers.Front()
	for numbers.Len() > 1 {
		for i := 1; i < m; i++ {
			current = current.Next()
			if current == nil {
				current = numbers.Front()
			}
		}
		numbers.Remove(current)
	}

	return current.Value.(int)
}

func TestLastRemaining(t *testing.T) {
	fmt.Println(LastRemaining(4, 3))
}