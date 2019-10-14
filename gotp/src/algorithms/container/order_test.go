package container

import (
	"testing"
)



func setup() *TreeNode {
	//      1
	//   2    3
	// 4   5     6
	//   7   8 9
	// PreOrder: 1 2 4 5 7 8 3 6 9
	// InOrder: 4 2 7 5 8 1 3 6 9
	node9 := &TreeNode{nil, 9, nil}
	node8 := &TreeNode{nil, 8, nil}
	node7 := &TreeNode{nil, 7, nil}
	node6 := &TreeNode{node9, 6, nil}
	node5 := &TreeNode{node7, 5, node8}
	node4 := &TreeNode{nil, 4, nil}
	node3 := &TreeNode{nil, 3, node6}
	node2 := &TreeNode{node4, 2, node5}
	node1 := &TreeNode{node2, 1, node3}

	return node1
}

func PreOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	visit(root.Val)
	PreOrderTraverse(root.Left)
	PreOrderTraverse(root.Right)
}

func PreOrderIter(root *TreeNode) {
	if root == nil {
		return
	}
	s := new(Stack)
	s.Push(root)
	for len(*s) != 0 {
		node := s.Pop()
		visit(node.Val)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
}

func InOrderTrave(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTrave(root.Left)
	visit(root.Val)
	InOrderTrave(root.Right)
}

func InOrderIter(root *TreeNode) {
	if root == nil {
		return
	}
	s := new(Stack)
	for {
		if root != nil {
			s.Push(root)
			root = root.Left
		} else {
			if len(*s) == 0 {
				break
			}
			root = s.Pop()
			visit(root.Val)
			root = root.Right
		}
	}
}

func TestOrder(t *testing.T) {
	root := setup()
	t.Run("PreOrder", func(t *testing.T) {
		PreOrderTraverse(root) // 1 2 4 5 7 8 3 6 9
	})
	t.Run("PreOrder", func(t *testing.T) {
		PreOrderIter(root) // 1 2 4 5 7 8 3 6 9
	})
	t.Run("InOrder", func(t *testing.T) {
		InOrderTrave(root) // 4 2 7 5 8 1 3 9 6
	})
	t.Run("InOrder", func(t *testing.T) {
		InOrderIter(root) // 4 2 7 5 8 1 3 9 6
	})
}
