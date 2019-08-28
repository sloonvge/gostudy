package offer

import (
	"fmt"
	"testing"
)

type ListNode struct {
	value int
	next *ListNode
}

func GenerateListNode(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode{value:a[0]}
	var tmp *ListNode
	tmp = head
	for i := 1; i < len(a); i++ {
		node := &ListNode{value:a[i]}
		tmp.next = node
		tmp = node
	}
	return head
}

func PrintListNode(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Printf("%d ", head.value)
	PrintListNode(head.next)
}

func EqualListNode(head1, head2 *ListNode) bool {
	if head1 == nil && head2 == nil {
		return true
	}
	if head1 == nil || head2 == nil {
		return false
	}
	var equal bool
	if head1.value == head2.value {
		equal =  true
	}
	return equal && EqualListNode(head1.next, head2.next)
}


// 25
func MergeSortedLinkedList(head1, head2 *ListNode) *ListNode{
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}
	var mergedHead *ListNode
	if head1.value < head2.value {
		mergedHead = head1
		mergedHead.next = MergeSortedLinkedList(head1.next, head2)
	} else {
		mergedHead = head2
		mergedHead.next = MergeSortedLinkedList(head1, head2.next)
	}
	return mergedHead
}

func MergeSortedLinkedList2(head1, head2 *ListNode) *ListNode{
	var mergedNode, cur *ListNode
	mergedNode = &ListNode{
		value:-1,
	}
	cur = mergedNode
	for head1 != nil && head2 != nil {
		if head1.value < head2.value {
			cur.next = head1
			head1 = head1.next
		} else {
			cur.next = head2
			head2 = head2.next
		}
		cur = cur.next
	}
	if head1 != nil {
		cur.next = head1
	} else {
		cur.next = head2
	}
	return mergedNode.next
}

func TestMergeSortedLinkedList(t *testing.T) {
	inputs := []*ListNode{
		GenerateListNode([]int{1,3,5,7}),
		GenerateListNode([]int{1}),
	}
	helper := [][]*ListNode {
		{
		GenerateListNode([]int{1, 3, 5, 7}),
		},
		{
			GenerateListNode([]int{2}),
		},
	}
	wants := [][]*ListNode{
		{
		GenerateListNode([]int{1,1,3,3,5,5,7,7}),
		},
		{
			GenerateListNode([]int{1, 2}),
		},
	}
	t.Run("0", func(t *testing.T) {
		head := GenerateListNode([]int{1,3,5,7})
		fmt.Println(EqualListNode(head, GenerateListNode([]int{1,3,5,7, 9})))
	})
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := MergeSortedLinkedList(inputs[i],
					helper[i][j]); !EqualListNode(got, wants[i][j]){
					PrintListNode(wants[i][j])
					fmt.Println()
					PrintListNode(got)
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
	t.Run("2", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := MergeSortedLinkedList2(inputs[i],
					helper[i][j]); !EqualListNode(got, wants[i][j]){
					PrintListNode(wants[i][j])
					fmt.Println()
					PrintListNode(got)
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}
