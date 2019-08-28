package offer

import (
	"testing"
)

type BinaryTreeNode struct {
	value int
	left *BinaryTreeNode
	right *BinaryTreeNode
}

func GenerateBinaryTree(a []int) *BinaryTreeNode{
	i := -1
	return generateBinaryTree(a, &i)
}
func generateBinaryTree(a []int, i *int) *BinaryTreeNode {
	*i++
	if *i >= len(a) {
		return nil
	}
	if a[*i] == -1 {
		return nil
	}
	var root *BinaryTreeNode
	root = &BinaryTreeNode{value:a[*i]}
	root.left = generateBinaryTree(a, i)
	root.right = generateBinaryTree(a, i)
	return root
}

// 26
func HasSubtree(root1, root2 *BinaryTreeNode) bool {
	var result bool
	if root1 != nil && root2 != nil {
		if root1.value == root2.value {
			result = hasSubTree(root1, root2)
		}
		if !result {
			result = HasSubtree(root1.left, root2)
		}
		if !result {
			result = HasSubtree(root1.right, root2)
		}
	}

	return result
}
func hasSubTree(root1, root2 *BinaryTreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root1.value != root2.value {
		return false
	}

	return hasSubTree(root1.left, root2.left) && hasSubTree(root1.right, root2.right)
}

func TestHasSubtree(t *testing.T) {
	inputs := []*BinaryTreeNode{
		GenerateBinaryTree([]int{8, 8, 9, 2, 4, 7, 7}),
	}
	helper := [][]*BinaryTreeNode {{
		GenerateBinaryTree([]int{8, 9, 2}),
		GenerateBinaryTree([]int{8, 9, 3}),
		},
	}
	wants := [][]bool{
		{true, false,},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := HasSubtree(inputs[i],
					helper[i][j]); got != wants[i][j] {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}
