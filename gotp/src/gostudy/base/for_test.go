package base

import (
	"testing"
	"fmt"
)

func TestFor(t *testing.T) {
	n := 5
	i := 0
	for ;i < n; i++{
		fmt.Println(i)
	}
}

func TestForLoop(t *testing.T) {
	n := 5
	i := 0
	for i < n{

		i = 6
		fmt.Println(i)
		i++
	}
	fmt.Printf("%d\n", i)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func TestConvert(t *testing.T) {
	root := &TreeNode{
		Val:2,
		Left:&TreeNode{
			Val:1,
		},
		Right:&TreeNode{
			Val:3,
		},
	}
	convertBST(root)
}

func convertBST(root *TreeNode) *TreeNode {
	stack := NewStack()
	total := 0
	node := root
	for node != nil || len(stack.S) != 0 {
		if  node != nil {
			stack.Push(node)
			fmt.Printf("push:%v\n", node)
			node = node.Right
		} else {
			node := stack.Pop()

			total += node.Val
			node.Val = total

			node = node.Left
			fmt.Printf("%v%v\n", node, len(stack.S) != 0 || node != nil)
		}

	}
	return root
}

type Stack struct {
	S []*TreeNode
}

func NewStack() *Stack {
	return &Stack{S: make([]*TreeNode, 0)}
}

func (s *Stack) Push(x *TreeNode) {
	s.S = append(s.S, x)
}

func (s *Stack) Pop() *TreeNode {
	n := len(s.S)
	v := s.S[n - 1]
	s.S = s.S[:n - 1]
	return v
}
