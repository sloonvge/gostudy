package offer

import (
	"fmt"
	"log"
	"testing"
)

// 两个栈实现队列
type Stack []int

func (s *Stack) Push(a int) {
	if s == nil {
		*s = make([]int, 0)
	}
	*s = append(*s, a)
}

func (s *Stack) Pop() (v int) {
	if *s == nil {
		panic("stack is nil")
		return
	}
	if len(*s) == 0 {
		panic("stack is already empty.")
		return
	}
	v = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}


type Queue struct {
	stack1 *Stack
	stack2 *Stack
}

func NewQueue() *Queue {
	return &Queue{
		stack1: new(Stack),
		stack2: new(Stack),
	}
}

func (q *Queue) AppendTail(a int) {
	q.stack1.Push(a)
}

func (q *Queue) DeleteHead() (v int) {
	if len(*q.stack2) != 0 {
		v = q.stack2.Pop()
		return
	}
	for len(*q.stack1) != 0 {
		a := q.stack1.Pop()
		q.stack2.Push(a)
	}
	v = q.stack2.Pop()
	return
}


func TestStack(t *testing.T) {
	t.Run("push", func(t *testing.T) {
		var s Stack
		log.Printf("before: %v\n", s)
		s.Push(1)
		s.Push(2)
		log.Printf("after push: %v\n", s)
		v := s.Pop()
		log.Printf("after pop: %v pop %d\n", s, v)
	})
	t.Run("pop nil", func(t *testing.T) {
		var s Stack
		s.Pop()
	})
}

func TestQueueWithTwoStack(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		q := NewQueue()
		for i := 0; i < 5; i++ {
			fmt.Printf("append %d\n", i)
			q.AppendTail(i)
		}
		for i := 0; i < 5; i++ {
			v := q.DeleteHead()
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	})
	t.Run("exist", func(t *testing.T) {
		q := NewQueue()
		q.stack1.Push(1)
		q.stack2.Push(0)

		for i := 0; i < 2; i++ {
			v := q.DeleteHead()
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	})
}


