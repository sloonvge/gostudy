package main

import "fmt"

type queue []int
func NewQueue() *queue {
	return &queue{}
}

func (s *queue) Push(x int) {
	*s = append(*s, x)
}
func (s *queue) Pop() int{
	old := *s
	x := old[0]
	*s = old[1:]
	return x
}
func (s queue) Len() int {
	return len(s)
}

type stack []int

func New() *stack {
	return &stack{}
}

func (s *stack) Push(x int) {
	*s = append(*s, x)
}
func (s *stack) Pop() int{
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func (s stack) Len() int {
	return len(s)
}

func reverse(expr string) {
	count := 0
	s1 := New()
	s2 := NewQueue()
	for i, s := range expr {
		if s == '(' {
			count++
			s1.Push(i)
		} else if s == ')' {
			count--
			s2.Push(i)
		}
	}
	if count != 0 {
		return
	}
	ret := []byte(expr)
	for s1.Len() != 0 {
		i1 := s1.Pop()
		i2 := s2.Pop()

		copy(ret[i1+1:i2], r(ret[i1+1:i2]))
		// fmt.Println(string(ret[i1+1:i2]), string(r(ret[i1+1:i2])))
	}
	fmt.Println(string(ret))
}

func r(s []byte) []byte {
	var ret []byte
	for i := 0; i < len(s); i++ {
		ret = append(ret, s[len(s) - i - 1])
	}
	return ret
}

func main() {
	reverse("((ur)oi)")
}