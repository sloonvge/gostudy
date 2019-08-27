package jd

import (
	"fmt"
	"testing"
)

// 在英文的输入中，我们经常会遇到大小写切换的问题，频繁切换大小写会增加我们的按键次数，
// 也会降低我们的打字效率。 众所周知，切换大小写有两种方式，一种是按下“caps locks”，
// 也就是大写锁定键，这样一来，之后的输入模式都会被切换。
// 另一种是同时按下shift和需要打印的字母，可以临时切换大小写(算作按下两个键)。
// 已知初始状态下，打字模式是小写，现在给出需要打印的字符串(区分大小写)，
// 请你计算出最少需按键多少次才能打印出来。

// 输入第一行仅包含一个正整数n，表示字符串的长度(1<=n<=1000000)。
//
// 输入第二行包含一个长度为n的字符串，仅包含大小写字母。

// 输出仅包含一个正整数，即最少的按键次数

// input:
// 6
// AaAAAA
// output:
// 8

func jd1(input string) int {
	// tem := make([]int, 0)
	// if input[0] < 'a' {
	// 	tem = append(tem, 0)
	// }
	// for i := 0; i < len(input) - 1; i++ {
	// 	if input[i] < 'a' && input[i+1] >= 'a'||
	// 		input[i] >= 'a' && input[i+1] < 'a' {
	// 		tem = append(tem, i + 1)
	// 	}
	// }
	// fmt.Printf("%v\n", tem)
	c := make([]int, len(input))
	s := make([]int, len(input))
	ca := 'a'
	sa := 'a'
	for i := 0; i < len(input); i++ {
		if i >= 1 {
			c[i] += c[i - 1]
			s[i] += s[i - 1]
		}
		if input[i] < 'a' { // A
			if ca == 'a' {
				c[i] += 2
				ca = 'A'
			} else if ca == 'A' {
				c[i] += 1
			}
			if sa == 'a' {
				s[i] += 2
			} else if sa == 'A' {
				s[i] += 1
			}
		}
		if 'a' <= input[i]{ // a
			if ca == 'a' {
				c[i] += 1
			} else if ca == 'A' {
				c[i] += 2
				ca = 'a'
			}
			if sa == 'a' {
				s[i] += 1
			} else if sa == 'A' {
				s[i] += 2
			}
		}
		if s[i] < c[i] {
			c[i] = s[i]
		}
		if s[i] > c[i] {
			s[i] = c[i]
			sa = ca
		}
	}
	fmt.Printf("%v %v\n", c, s)
	return s[len(input) - 1]
}

func TestJd1(t *testing.T) {
	fmt.Println(jd1("AaAAAA"))
	fmt.Println(jd1("AaAAa"))
	fmt.Println(jd1("AA"))
}




