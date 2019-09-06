package main
//
// import "fmt"
//
// type ListNode struct {
// 	value int
// 	next *ListNode
// }
//
// func GenerateListNode(a []int) *ListNode {
// 	if len(a) == 0 {
// 		return nil
// 	}
// 	head := &ListNode{value:a[0]}
// 	var tmp *ListNode
// 	tmp = head
// 	for i := 1; i < len(a); i++ {
// 		node := &ListNode{value:a[i]}
// 		tmp.next = node
// 		tmp = node
// 	}
// 	return head
// }
//
// func PrintListNode(head *ListNode) {
// 	if head == nil {
// 		return
// 	}
// 	fmt.Printf("%d ", head.value)
// 	PrintListNode(head.next)
// }
//
// func sort(node *ListNode, m int) *ListNode {
// 	if node == nil {
// 		return nil
// 	}
// 	var pre, cur, next *ListNode
// 	cur = node
// 	for ; cur.next != nil; {
// 		next = cur.next
// 		if cur.value <= m {
// 			pre, cur = cur, pre
// 		}
// 		cur = next
// 	}
//
// 	return node
// }
//
// func main() {
// 	node := GenerateListNode([]int{9, 6, 3, 7, 6, 5})
// 	PrintListNode(sort(node, 4))
// }
