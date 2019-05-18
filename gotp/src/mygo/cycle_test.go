package gostudy

import (
	"fmt"
	"testing"
)

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func TestHasCycle(t *testing.T) {
	var a, b, c, d *ListNode
	b = &ListNode{Val:2, Next:nil}
	d = &ListNode{Val:-4, Next:b,}
	c = &ListNode{Val:0,Next:d,}
	*b = ListNode{Val:2, Next:c}
	a = &ListNode{Val:3, Next:b}

	fmt.Println(hasCycle(a))

}

func hasCycle(head *ListNode) bool {
	for head != nil {
		if head == head.Next {
			return true
		}
		fmt.Printf("ori:%v-%v-%v-%v-%v\n", head, head.Next, head.Next.Next, head.Next.Next.Next, head.Next.Next.Next.Next)
		t := head.Next
		head.Next = head
		fmt.Printf("chnage:%v-%v\n", t, head.Next)
		head = t
		fmt.Printf("%v-%v-%v-%v-%v\n", head, head.Next, head.Next.Next, head.Next.Next.Next, head.Next.Next.Next.Next)
		// head, head.Next = head.Next, head
		// Show(head)
		//if head != nil {
			//Show(head.Next)
		//}

	}
	return false
}

func Show(head *ListNode) {
	seen := make(map[*ListNode]int, 0)
	for node := head; node != nil; node = node.Next {
		if _, ok := seen[node]; !ok {
			fmt.Printf("%v\t", node.Val)
			seen[node] = 1
		} else {
			fmt.Println()
			return
		}
	}
	fmt.Println()
}
