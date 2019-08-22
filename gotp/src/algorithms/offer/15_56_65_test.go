package offer

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func NumberOf1(n int) (ans int) {
	s := fmt.Sprintf("%b", n)
	for _, v := range s {
		if v == '1' {
			ans++
		}
	}
	return
}

func NumberOf12(n int) (ans int) {
	for i := n; i > 0; i = i >> 1 {
		if i & 1 == 1 {
			ans++
		}
	}
	return
}

func NumberOf13(n int) (ans int) {
	var flag = 1
	for flag != 0 {
		if n & flag != 0 {
			ans++
		}
		flag = flag << 1
	}
	return
}

func NumberOf14(n int) (ans int) {
	for n != 0 {
		ans++
		n &= n -1
	}
	return
}

func TestNumberOf1(t *testing.T) {
	inputs := []int{0, 1, 3, 0x7FFFFFFF, 0x80000000, 0xFFFFFFFF}
	wants := []int{0, 1, 2, 31, 1, 32}
	t.Run("1", func(t *testing.T) {
		for i, input := range inputs {
			got := NumberOf1(input)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
	// go print problem
	t.Run("2", func(t *testing.T) {
		for i, input := range inputs {
			got := NumberOf12(input)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
	t.Run("3", func(t *testing.T) {
		for i, input := range inputs {
			got := NumberOf13(input)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
	t.Run("4", func(t *testing.T) {
		for i, input := range inputs {
			got := NumberOf14(input)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
}

func TestRight(t *testing.T)  {
	t.Run("move right", func(t *testing.T) {
		var i int32 = -1
		fmt.Printf("int32:%#x\n", i)
		i >>= 1
		fmt.Printf(">>1:%#x\n", i)
		var u = uint32(i)
		fmt.Printf("uint32:%#x %d\n", u, u)
		u >>= 1
		fmt.Printf("uint32:%#x %d\n", u, u)


		var n2 = -1
		fmt.Printf("%b\n", n2)
		n2 = n2 >> 1
		fmt.Printf("%b\n", n2)
	})
}

func TestXOR(t *testing.T) {
	fmt.Printf("%d\n", 2 ^ 1 ^ 1)
	fmt.Printf("%d\n", 2^0)
}

// 56 p275
func NumberAppearOnce(data []int) (nums []int) {
	if len(data) == 0 {
		return
	}
	xorData := XORData(data)
	index := indexOf1(xorData)
	data1, data2 := splitData(data, index)
	num1 := XORData(data1)
	num2 := XORData(data2)
	nums = append(nums, num1)
	nums = append(nums, num2)
	return
}

func XORData(data []int) (n int){
	for i := 0; i < len(data); i++ {
		n ^= data[i]
	}
	return n
}
func indexOf1(n int) (index int) {
	var flag int = 1
	for {
		if n & flag != 0{
			break
		}
		flag = flag << 1
		index++
	}
	return
}

func splitData(data []int, index int) (data1, data2 []int) {
	for _, n := range data {
		if (n & (1 << uint(index))) != 0{
			data1 = append(data1, n)
		} else {
			data2 = append(data2, n)
		}
	}
	return
}

func TestNumbersAppearOnce(t *testing.T) {
	inputs := [][]int{
		{2, 4, 3, 6, 3, 2, 5, 5},
		{1, 2},
		{2, -7, 6, -3, -3, 2, 5, 5},
	}
	wants := [][]int{
		{4, 6},
		{1, 2},
		{-7, 6},
	}
	t.Run("1", func(t *testing.T) {
		for i, input := range inputs {
			got := NumberAppearOnce(input)
			sort.Ints(got)
			sort.Ints(wants[i])
			if !reflect.DeepEqual(got, wants[i]) {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
}

func NumbersOnlyOneAppearOnce(numbers []int) (num int) {
	if len(numbers) == 0 {
		return
	}
	bitSum := make([]int, 32)
	for _, number := range numbers {
		var bitMask = 1
		for j := 31; j >= 0; j-- {
			bit := number & bitMask
			if bit != 0 {
				bitSum[j]++
			}
			bitMask = bitMask << 1
		}
	}
	fmt.Printf("%v\n", bitSum)
	for i := 0; i < 32; i++ {
		num = num << 1
		num += bitSum[i] % 3
	}
	return
}

func TestNumbersOnlyOneAppearOnce(t *testing.T) {
	inputs := [][]int{
		{1, 3, 3, 3, 5, 5, 5},
		{0, 3, 3, 3, 5, 5, 5},
		{-1, 3, 3, 3, 5, 5, 5},
	}
	wants := []int{
		1,
		0,
		-1,
	}
	t.Run("1", func(t *testing.T) {
		for i, input := range inputs {
			got := NumbersOnlyOneAppearOnce(input)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
}

// 65
func Add(num1, num2 int) int {
	n1 := num1 ^ num2
	n2 := num1 & num2 << 1
	if n2 == 0 {
		return n1
	}
	return Add(n1, n2)
}

func TestAdd(t *testing.T) {
	fmt.Printf("2 + 3 = %d\n", Add(2, 3))
	fmt.Printf("2 + -1 = %d\n", Add(2, -1))
	fmt.Printf("2 + 0 = %d\n", Add(2, 0))
}