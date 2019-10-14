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

func GenerateCycleListNode(a []int, b []int) *ListNode {
	if len(a) == 0 && len(b) == 0{
		return nil
	}
	head1 := &ListNode{value:a[0]}
	var tmp *ListNode
	tmp = head1
	for i := 1; i < len(a); i++ {
		node := &ListNode{value:a[i]}
		tmp.next = node
		tmp = node
	}
	head2 := &ListNode{value:b[0]}
	tmp = head2
	for i := 1; i < len(b); i++ {
		node := &ListNode{value:b[i]}
		tmp.next = node
		tmp = node
		if i == len(b) - 1 {
			tmp.next = head2
		}
	}
	head1.next = head2
	return head1
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

func EqualOneValue(head1, head2 *ListNode) bool {
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
	return equal
}

// 6
func PrintListReversingly(head *ListNode) {
	if head == nil {
		return
	}
	PrintListReversingly(head.next)
	fmt.Println(head.value)
}

func TestPrintListReversingly(t *testing.T) {
	head := GenerateListNode([]int{1,3,5,7})
	PrintListReversingly(head)
}

// 18
func DeleteNode(head *ListNode, node *ListNode) (ok bool) {
	if head == nil || node == nil {
		return false
	}
	if node.next != nil {
		next := node.next
		node.value = next.value
		node.next = next.next
		next = nil
	} else if head == node {
		node = nil
		head = nil
	} else {
		n := head
		for ; n.next.value != node.value;  {
			n = n.next
		}
		n.next = nil
		node = nil
	}

	return ok
}

func TestDeleteNode(t *testing.T) {
	inputs := []*ListNode{
		GenerateListNode([]int{1,3,5,7}),
	}
	helper := [][]*ListNode {
		{
			GenerateListNode([]int{3}),
		},
	}
	wants := [][]*ListNode{
		{
			GenerateListNode([]int{1,5,7}),
		},

	}
	t.Run("0", func(t *testing.T) {
	})
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := DeleteNode(inputs[i],
					helper[i][j]); !EqualListNode(inputs[i], wants[i][j]){
						PrintListNode(inputs[i])
						fmt.Println()
						PrintListNode(wants[i][j])
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}

// 22
func FindKthToTail(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	i := 0
	return findKthToTail(head, k, &i)
}
func findKthToTail(head *ListNode, k int, i *int) *ListNode {
	if head == nil {
		return nil
	}
	var node *ListNode
	node = findKthToTail(head.next, k, i)
	*i++
	if k == *i {
		return head
	}
	return node
}

func FindKthToTail2(head *ListNode, k uint) *ListNode {
	if head == nil || k == 0{
		return nil
	}
	var ahead, behind *ListNode
	ahead = head
	behind = head
	for i := 0; uint(i) < k - 1; i++ {
		ahead = ahead.next
		if ahead == nil {
			return nil
		}
	}
	for ahead.next != nil {
		ahead = ahead.next
		behind = behind.next
	}
	return behind
}

func TestFindKthToTail(t *testing.T) {
	inputs := []*ListNode{
		GenerateListNode([]int{1,3,5,7}),
	}
	helper := [][]int{
		{1, 3, 4, 5},
	}
	wants := [][]*ListNode{
		{
			GenerateListNode([]int{7}),
			GenerateListNode([]int{3}),
			GenerateListNode([]int{1}),
			GenerateListNode([]int{}),
		},
	}
	t.Run("0", func(t *testing.T) {
	})
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := FindKthToTail(inputs[i],
					helper[i][j]); !EqualOneValue(got, wants[i][j]){
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
	t.Run("2", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := FindKthToTail2(inputs[i],
					uint(helper[i][j])); !EqualOneValue(got, wants[i][j]){
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}

// 23
func EntryNodeOfLoop(head *ListNode) *ListNode {
	meet := meetingNode(head)
	if meet == nil {
		return nil
	}
	nodesInLoop := 1
	for node := meet; node.next != meet ; node = node.next {
		nodesInLoop++
	}
	var first, second *ListNode
	first = head
	second = head
	for i := 0; i < nodesInLoop; i++ {
		first = first.next
	}
	for first != second {
		first = first.next
		second = second.next
	}
	return first
}
func meetingNode(head *ListNode) *ListNode {
	var slow, fast, meet *ListNode
	slow = head
	fast = head
	for slow != nil && fast != nil {
		slow = slow.next
		fast = fast.next
		if fast != nil {
			fast = fast.next
		}
		if slow == fast {
			meet = slow
			break
		}

	}
	return meet
}

func TestEntryNodeOfLoop(t *testing.T) {
	inputs := []*ListNode{
		GenerateCycleListNode([]int{1,2,},[]int{3,4,5,6}),
		GenerateListNode([]int{1,3,5,7}),
	}
	wants := [][]*ListNode{
		{
			GenerateListNode([]int{3}),
		},
		{
			GenerateListNode([]int{}),
		},
	}
	t.Run("0", func(t *testing.T) {
	})
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := EntryNodeOfLoop(inputs[i]); !EqualListNode(got, wants[i][j]){
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}

// 24
func ReverseList(head *ListNode) *ListNode {
	var pre, cur, next *ListNode
	cur = head
	for cur != nil  {
		next = cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	return pre
}

func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.next == nil  {
		return head
	}
	var recvHead *ListNode
	recvHead = ReverseList2(head.next)
	head.next.next = head
	head.next = nil
	return recvHead
}

func TestReverseList(t *testing.T) {
	inputs := []*ListNode{
		GenerateListNode([]int{1,3,5,7}),
		GenerateListNode([]int{1}),
		GenerateListNode([]int{}),
	}
	wants := [][]*ListNode{
		{
			GenerateListNode([]int{7,5,3,1}),
		},
		{
			GenerateListNode([]int{1}),
		},
		{
			GenerateListNode([]int{}),
		},
	}
	t.Run("0", func(t *testing.T) {
	})
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := ReverseList(inputs[i]); !EqualListNode(got, wants[i][j]){
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
				if got := ReverseList2(inputs[i]); !EqualListNode(got, wants[i][j]){
					PrintListNode(wants[i][j])
					fmt.Println()
					PrintListNode(got)
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
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

// 52
func FindFirstCommonNode(head1, head2 *ListNode) *ListNode {
	len1 := 0
	len2 := 0
	for node := head1; node != nil ; node = node.next {
		len1++
	}
	for node := head2; node != nil; node = node.next {
		len2++
	}
	var longList, shortList *ListNode
	longList = head2
	shortList = head1
	diff := len2 - len1
	if len1 > len2 {
		longList = head1
		shortList = head2
		diff = len1 - len2
	}

	for i := 0; i < diff; i++ {
		longList = longList.next
	}
	for longList != nil && shortList != nil && longList != shortList {
		shortList = shortList.next
		longList = longList.next
	}
	return longList
}