package algorithms

import (
	"testing"
	"fmt"
	"strings"
)

//
//链表
//

type Node struct {
	Item int
	Next *Node

}

// 在表头插入节点

func TestListedFirst(t *testing.T) {
	node := Node{
		Item: 1,
	}

	nodeNew := Node{
		Item: 2,
		Next: &node,
	}

	fmt.Printf("%v", nodeNew)
	var s *string
	*s = "44"
	print(strings.Contains(*s, "4"))
}

type Queue struct {
	first *Node
	last *Node
	N int
}

func (q *Queue) Enqueue(x int) {
	oldlast := q.last
	q.last = &Node{}
	q.last.Item = x
	if q.first == nil {
		q.first = q.last
	} else {
		oldlast.Next = q.last
	}
	q.N++
}

func (q *Queue) dequeue() *Node{
	oldfirst := q.first
	q.first = q.first.Next
	if q.N == 0 {
		q.last = nil



	}
	q.N--

	return oldfirst
}


