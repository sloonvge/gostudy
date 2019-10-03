package base

import (
	"fmt"
	"testing"
)

type CircularList struct {
	array []int
	head int
	tail int
	len int
	empty bool
}

func New(n int) *CircularList {
	array := make([]int, n + 1)
	return &CircularList{array:array, len:n + 1, empty:true}
}

func (c *CircularList) Enqueue(x int) {
	if (c.tail + 1) % c.len == c.head {
		c.head = (c.head + 1) % c.len
	}
	c.array[c.tail] = x
	c.tail = (c.tail + 1) % c.len
}

func (c *CircularList) Dequeue() (x int) {
	if c.head == c.tail {
		c.empty = true
		return
	}
	x = c.array[c.head]
	c.array[c.head] = 0
	c.head = (c.head + 1) % c.len
	return x
}

func (cl *CircularList) Len() int {
	return cl.len - 1
}

func TestQueue(t *testing.T) {
	t.Run("circular list", func(t *testing.T) {
		cl := New(5)
		n := cl.Len()
		for i := 0; i < n + 3; i++ {
			cl.Enqueue(i)
			fmt.Printf("enqueue: head=%d, tail=%d\n", cl.head, cl.tail)
		}

		for i := 0; i < n; i++ {
			x := cl.Dequeue()
			fmt.Printf("dequeue: head=%d, tail=%d\n", cl.head, cl.tail)
			fmt.Println(x)
		}
	})
}
