package base

import (
	"fmt"
	"sort"
	"testing"
)

/** 
* Created by wanjx in 2019/3/27 21:54
**/
 
/*
 栈
 */

 type MinStack struct {
 	s []int
}

func NewMinStack() MinStack {
	return MinStack{s: make([]int, 0)}
}

func (m *MinStack) Push(x int) {
	m.s = append(m.s, x)
}

func TestNewMinStack(t *testing.T) {
	ms := NewMinStack()
	ms.s = append(ms.s, 1)
	fmt.Printf("%v\n", ms.s)
	ms.Push(2)
	fmt.Printf("%v\n", ms.s)
	c := make([]int, 0)
	c =append(c, ms.s...)
	c[1] = 3
	fmt.Printf("%v %v\n", c, ms.s)
}


type stack []int

func NewStack() *stack{
	return &stack{}
}

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() int {
	n := len(*s)
	v := (*s)[n - 1]
	*s = (*s)[:n -1]
	return v
}

func (s *stack) Top() int {
	v := (*s)[len(*s) - 1]
	return v
}

func (s *stack) GetMin() int {
	c := make([]int, 0)
	copy(c, *s)
	sort.Ints(c)
	return c[0]
}