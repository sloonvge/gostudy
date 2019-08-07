package base

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	nums :=[]int{1, 2, 3}
	changeSlice(nums)
	fmt.Printf("out func: %p %v\n", &nums, nums)
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
