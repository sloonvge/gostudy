package tree

import (
	"testing"
	"fmt"
)

func TestListNode_String(t *testing.T) {
	var head, node *ListNode
	head = &ListNode{Val: 0}
	for i := 1; i < 5; i++ {
		node = &ListNode{Val: i}
		head.Add(node)
	}

	fmt.Println(head)
}
