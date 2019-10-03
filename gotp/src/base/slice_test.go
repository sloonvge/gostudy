package base

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	nums :=[]int{1, 2, 3}
	changeSlice(nums)
	fmt.Printf("out func: %p %v\n", &nums, nums)
	fmt.Printf("out func: after append: %p len: %d cap: %d\n",
		&nums, len(nums), cap(nums))
}

func changeSlice(nums []int) {
	if len(nums) < 3{
		panic("slice len is small than 3")
	}

	nums[2] = -1
	fmt.Printf("before append: %p len: %d cap: %d\n",
		&nums, len(nums), cap(nums))
	nums = append(nums, -1)
	fmt.Printf("append: %v\n", nums)
	fmt.Printf("after append: %p len: %d cap: %d\n",
		&nums, len(nums), cap(nums))
}

func TestSlicePtr(t *testing.T) {
	nums :=[]int{1, 2, 3}
	fmt.Printf("out func: %p %v\n", &nums, nums)
	changeSlicePtr(&nums)
	fmt.Printf("out func: after append: %p len: %d cap: %d\n",
		&nums, len(nums), cap(nums))
}

func changeSlicePtr(nums *[]int) {
	old := *nums
	if len(*nums) < 3{
		panic("slice len is small than 3")
	}

	old[2] = -1
	fmt.Printf("before append: %p len: %d cap: %d\n",
		&nums, len(old), cap(old))
	*nums = append(old, -1)
	fmt.Printf("append: %v\n", *nums)
	fmt.Printf("after append: %p len: %d cap: %d\n",
		&nums, len(*nums), cap(*nums))
}

func  TestDoubleSlice(t *testing.T)  {
	nums := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	changeDoubleSlice(nums)
	fmt.Printf("%v\n", nums)
}

func changeDoubleSlice(nums [][]int) {
	if len(nums) < 3{
		panic("slice len is small than 3")
	}

	nums[2][2] = -1
}

func TestSliceAppend(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		var s []int
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
		s = append(s, 1, 2, 3, 4, 5)
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	})
	t.Run("1", func(t *testing.T) {
		s := []int{1, 2}
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
		s = append(s, 4)
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
		s = append(s, 5)
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
		s = append(s, 6)
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	})
	t.Run("1", func(t *testing.T) {
		s := []int{1, 2, 3}
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
		s = append(s, 4)
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
		s = append(s, 5)
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
		s = append(s, 6)
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
		s = append(s, 7)
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	})
	t.Run("2", func(t *testing.T) {
		s := []int{1, 2}
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
		s = append(s, 4, 5, 6)
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	})
	t.Run("2", func(t *testing.T) {
		s := []int{1, 2, 3}
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
		s = append(s, 4, 5, 6, 7)
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	})
	t.Run("3", func(t *testing.T) {
		s := make([]int, 0)

		oldCap := cap(s)

		for i := 0; i < 2048; i++ {
			s = append(s, i)

			newCap := cap(s)

			if newCap != oldCap {
				fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
				oldCap = newCap
			}
		}
	})
}
