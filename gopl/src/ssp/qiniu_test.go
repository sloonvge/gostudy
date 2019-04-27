package ssp

import (
	"fmt"
	"testing"
)

/** 
* Created by wanjx in 2019/3/15 22:07
**/

// go defer在for循环里的情况

func TestForDefer(t *testing.T) {
	for i:=0; i < 5; i++ {
		defer t.Logf("%d ", i)
		// 4 3 2 1 0
		// defer执行顺序是先进后出
	}
}

// return：1. 返回值赋值 2. 调用defer表达式 3. 返回给调用函数
func TestReturn(t *testing.T) {
	got := increase(1)
	t.Logf("%d", got)
}

func increase(x int) (y int) {
	defer func() {
		fmt.Println(y)
		y++
		fmt.Println(y)
	}()
	return x
}
// defer表达式有返回值，将会被丢弃

// defer: 闭包与匿名函数
func TestCloseBag(t *testing.T)  {
	for i:=0; i < 5; i++{
		defer func() {
			t.Logf("%d", i)
			// 5 5 5 5 5
		}()
	}
}

func TestCloseBag2(t *testing.T)  {
	for i:=0; i < 5; i++{
		defer func(i int) {
			t.Logf("%d", i)
			// 4 3 2 1 0
		}(i)
	}
}
// defer 例
func TestDefer( t *testing.T) {
	t.Logf("%d", testDefer())
}
func testDefer() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
