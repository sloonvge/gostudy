package main

import (
	"testing"
	"unsafe"
)

// 获取slice内存地址
func TestSlicePointer(t *testing.T) {
	a := []int{ 1, 2, 3}
	b := a[:2]
	c := &a
	t.Logf("%v", &a)
	t.Logf("a: %v b: %v c: %v %x",
		unsafe.Pointer(&a), unsafe.Pointer(&b), unsafe.Pointer(&c), c)
}

// append函数扩容原理
func appendInt(x []int, y ...int) []int{
	var z []int
	// zlen := len(x) + 1
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	// z[len(x)] = y
	copy(z[len(x):], y)
	return z
}

func TestSliceAppend(t *testing.T) {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		t.Logf("%d%d len=%d\tcap=%d\t%v\n",
			unsafe.Pointer(&y), i, len(y), cap(y), y)
		x = y
	}
}

// 去除空字符串
func noEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}

	}
	return strings[:i]
}

func noEmpty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func TestSliceNoEmpty(t *testing.T) {
	s := []string{"a", "", "b"}
	t.Logf("s=%q", s)
	s = noEmpty(s)
	t.Logf("s=%q", s)

}

// slice实现栈
type StackInt struct {
	S []int
}

func (s *StackInt) Push(i int) {
	s.S = append(s.S, i)
}

func (s *StackInt) Pop() int {
	top := s.S[len(s.S) - 1]
	s.S = s.S[:len(s.S) - 1]
	return top
}

// 删除数组中某一元素
func remove(s []int, i int) []int{
	copy(s[i:], s[i+1:])
	return s[:len(s) - 1]
}

func remove2(s []int, i int) []int {
	s[i] = s[len(s) - 1]
	return s[:len(s) - 1]
}

// slice 测试结构体
type SliceAdd []int

func (s *SliceAdd) Add(v int) {
	*s = append(*s, v)
}

func (s SliceAdd) swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestSliceAdd( t *testing.T) {
	s := SliceAdd{}
	for i:=0;i<5;i++ {
		s.Add(i)
	}
	t.Logf("%v", s)
	s.swap(0, 4)
	t.Logf("%v", s)
}