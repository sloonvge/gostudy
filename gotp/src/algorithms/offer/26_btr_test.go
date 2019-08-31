package offer

import (
	"fmt"
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

func PrintBinaryTreeFront(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.value)
	PrintBinaryTreeFront(root.left)
	PrintBinaryTreeFront(root.right)
}
func PrintBinaryTreeMid(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	PrintBinaryTreeMid(root.left)
	fmt.Printf("%d ", root.value)
	PrintBinaryTreeMid(root.right)
}
func PrintBinaryTreeEnd(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	PrintBinaryTreeEnd(root.left)
	PrintBinaryTreeEnd(root.right)
	fmt.Printf("%d ", root.value)
}

func EqualValueBinaryTree(root1, root2 *BinaryTreeNode) (equal bool){
	if root1 == nil || root2 == nil {
		return root1 == nil && root2 == nil
	}
	equal = root1.value == root2.value
	return equal && EqualValueBinaryTree(root1.left, root2.left) &&
		EqualValueBinaryTree(root1.right, root2.right)
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

// 27
func MirrorRecursively(root *BinaryTreeNode){
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		return
	}
	// temp := root.left
	// root.left = root.right
	// root.right = temp
	root.left, root.right = root.right, root.left
	if root.left != nil {
		MirrorRecursively(root.left)
	}
	if root.right != nil {
		MirrorRecursively(root.right)
	}
}

func TestMirrorRecursively(t *testing.T) {
	inputs := []*BinaryTreeNode{
		GenerateBinaryTree([]int{8, 6, 5, 7, 10, 9, 11}),
	}
	wants := [][]*BinaryTreeNode{
		{
			GenerateBinaryTree([]int{8, 6, 5, 7, 10, 9, 11}),
		},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			in := inputs[i]
			PrintBinaryTreeMid(in)
			fmt.Println()
			for j := 0; j < len(wants[i]); j++ {
				wa := wants[i][j]
				MirrorRecursively(in)
				if !EqualValueBinaryTree(in, wa) {
					PrintBinaryTreeMid(in)
					fmt.Println()
					PrintBinaryTreeMid(wa)
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wa, in)
				}
			}
		}
	})
	t.Run("1", func(t *testing.T) {
		a4 := &BinaryTreeNode{
			value:4,
		}
		a3 := &BinaryTreeNode{
			value:3,
		}
		a2 := &BinaryTreeNode{
			value:2,
			left:a4,
		}
		a1 := &BinaryTreeNode{
			value:1,
			left:a2,
			right:a3,
		}
		PrintBinaryTreeFront(a1)
		fmt.Println()
		PrintBinaryTreeMid(a1)
		fmt.Println()
		PrintBinaryTreeEnd(a1)
		fmt.Println()
	})
}
