package main

import "fmt"

/** 
* Created by wanjx in 2019/4/27 16:21
**/

// Solution 1. 移动n, 复数会造成死循环
func NumberOf1(n int) int {
	fmt.Printf("%b\t%b\n", n, n >> 1)
	var ret int

	for x := n; x & 1 == 1; x >>= 1 {
		ret += 1
	}
	return ret
}

func NumberOf12(n int) int {
	fmt.Printf("%b\n", n)
	var ret int

	var x uint
	for x = 1; x != 0; x <<= 1 {
		if (uint(n) & x) != 0 {
			ret += 1
		}
	}
	return ret
}

func NumberOf13(n int) int {
	var r int
	fmt.Printf("%b %b\n", n, n -1)
	for x := n; x != 0; x &= x - 1 {
		r += 1
	}
	return r
}

func main() {
	v := NumberOf13(0x80000000)

	fmt.Println(v)
}
