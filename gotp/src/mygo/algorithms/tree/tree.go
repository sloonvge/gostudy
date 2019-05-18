package tree

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func (n *ListNode) Add(node *ListNode) {
	var t *ListNode
	for t = n; t != nil; t = n.Next {
		break
	}
	fmt.Println(t)
	t.Next = node

}

func (n *ListNode) String() string {
	s := "["
	for node := n; node != nil; node = node.Next {
		s += fmt.Sprintf("%d,", node.Val)
	}
	s += "]"
	return s
}
