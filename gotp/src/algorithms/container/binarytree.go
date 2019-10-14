package container

import (
	"fmt"
)

type TreeNode struct {
	Left *TreeNode
	Val int
	Right *TreeNode
}

func InOrderTraverse(t *TreeNode) {
	if t == nil {
		return
	}
	InOrderTraverse(t.Left)
	visit(t.Val)
	InOrderTraverse(t.Right)
}
func InOrderIteration(t *TreeNode) {
	if t == nil {
		return
	}
	s := new(Stack)
	for {
		if t != nil {
			s.Push(t)
			t = t.Left
		} else {
			if len(*s) == 0 {
				break
			}
			t = s.Pop()
			visit(t.Val)
			t = t.Right
		}
	}
}

func visit(i int) {
	fmt.Printf("%d ", i)
}