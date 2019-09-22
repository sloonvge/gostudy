package offer

import (
	"fmt"
	"testing"
)

type BinaryTreeNode struct {
	value int
	left *BinaryTreeNode
	right *BinaryTreeNode
	parent *BinaryTreeNode
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

// 7
func RebuildBinaryTree(pre, in []int) *BinaryTreeNode{
	return rebuildBinaryTree(pre, in, 0, len(pre) - 1, 0, len(pre) - 1)
}
func rebuildBinaryTree(pre, in []int, s1, e1, s2, e2 int) *BinaryTreeNode {
	root := &BinaryTreeNode{}
	root.value = pre[s1]

	if s1 == e1 {
		if s2 == e2 && pre[s1] == in[s2]{
			return root
		} else {
			panic("invalid input")
		}
	}

	i := s2
	for i <= e2 && in[i] != pre[s1] {
		i++
	}
	if i == e2 && in[i] != pre[s1] {
		panic("invalid input")
	}
	leftLength := i - s2
	leftPreorderEnd := s1 + leftLength
	if leftLength > 0 {
		root.left = rebuildBinaryTree(pre, in, s1 + 1, leftPreorderEnd, s2, i - 1)
	}
	if leftLength < e1 - s1 {
		root.right = rebuildBinaryTree(pre, in, leftPreorderEnd + 1, e1, i + 1, e2)
	}

	return root

}

func TestRebuildBinaryTree(t *testing.T) {
	head := RebuildBinaryTree([]int{1,2,4,7,3,5,6,8},
		[]int{4,7,2,1,5,3,8,6})
	PrintBinaryTreeFront(head)
	fmt.Println()
	PrintBinaryTreeMid(head)
}

// 8
func GetNext(node *BinaryTreeNode) *BinaryTreeNode {
	next := &BinaryTreeNode{}
	if node.right != nil {
		right := node.right
		for right.left != nil {
			right = right.left
		}

		next = right
	} else if node.parent != nil {
		current := node
		parent := node.parent
		for parent != nil && current == parent.right {
			current = parent
			parent = parent.parent
		}
		next = parent
	}

	return next
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

// 28
func IsSymmetrical(root *BinaryTreeNode) bool {
	return isSymmetrical(root, root)
}
func isSymmetrical(root1, root2 *BinaryTreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == nil && root2 == nil
	}
	if root1.value != root2.value {
		return false
	}
	fmt.Println(root1.value, root1.left.value, root1.right.value, root2.value)
	return isSymmetrical(root1.left, root2.right) &&
		isSymmetrical(root1.right, root2.left)
}

func TestIsSymmetrical(t *testing.T) {
	inputs := []*BinaryTreeNode{
		GenerateBinaryTree([]int{8, 6, 5, 7, 6, 7, 5}),
	}
	wants := [][]bool{
		{
			true,
		},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			input := inputs[i]
			// PrintBinaryTreeFront(input)
			// fmt.Println()
			for j := 0; j < len(wants[i]); j++ {
				want := wants[i][j]
				if got := IsSymmetrical(input); got != want {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, want, got)
				}
			}
		}
	})
}

// 29
func MatrixClockwisely(a [][]int, rows, cols int) {
	if len(a) == 0 || cols <= 0 || rows <= 0 {
		return
	}
	start := 0

	for cols > start * 2 && rows > start *2 {
		printMatrixInCycle(a, rows, cols, start)
		start++
	}
	fmt.Println()
}
func printMatrixInCycle(a [][]int, rows, cols, start int) {
	endX := cols - 1 - start
	endY := rows - 1- start
	for i := start; i <= endX; i++ {
		fmt.Printf("%d ", a[start][i])
	}
	if start < endY {
		for i := start + 1; i <= endY; i++ {
			fmt.Printf("%d ", a[i][endX])
		}
	}
	if start < endX && start < endY {
		for j := endX - 1; j >= start; j-- {
			fmt.Printf("%d ", a[endY][j])
		}
	}
	if start < endX && start < endY - 1 {
		for i := endY - 1; i >= start + 1; i-- {
			fmt.Printf("%d ", a[i][start])
		}
	}
}

func TestMatrixClockwisely(t *testing.T) {
	inputs := [][][]int{
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		{
			{1},
			{5},
			{9},
			{13},
		},
	}

	t.Run("", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			input := inputs[i]
			MatrixClockwisely(input, len(input), len(input[0]))
		}
	})
}

// 54
func KthNode(root *BinaryTreeNode, k int) *BinaryTreeNode{
	return kthNode(root, &k)
}
func kthNode(root *BinaryTreeNode, k *int) *BinaryTreeNode{
	var target *BinaryTreeNode
	if root == nil {
		return nil
	}


	target = kthNode(root.left, k)
	if target == nil {
		if *k == 1 {
			target = root
		}
		*k--
	}
	if target == nil {
		target = kthNode(root.right, k)
	}

	// if root.left != nil {
	// 	target = kthNode(root.left, k)
	// }
	// if target == nil {
	// 	if *k == 1 {
	// 		target = root
	// 	}
	// 	*k--
	// }
	//
	// if target == nil && root.right != nil {
	// 	target = kthNode(root.right, k)
	// }
	return target
}

func TestKthNode(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		a4 := &BinaryTreeNode{
			value:1,
		}
		a3 := &BinaryTreeNode{
			value:4,
		}
		a2 := &BinaryTreeNode{
			value:2,
			left:a4,
		}
		a1 := &BinaryTreeNode{
			value:3,
			left:a2,
			right:a3,
		}
		node := KthNode(a1, 4)
		fmt.Println(*node)
	})
}

// 55
func TreeDepth(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	l := TreeDepth(root.left)
	r := TreeDepth(root.right)
	if l > r {
		return l + 1
	}
	return r + 1
}

func TestTreeDepth(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		a4 := &BinaryTreeNode{
			value:1,
		}
		a3 := &BinaryTreeNode{
			value:4,
		}
		a2 := &BinaryTreeNode{
			value:2,
			left:a4,
		}
		a1 := &BinaryTreeNode{
			value:3,
			left:a2,
			right:a3,
		}
		d := TreeDepth(a1)
		fmt.Println(d)
	})
}

func IsBalanced(root *BinaryTreeNode) bool {
	if root == nil {
		return true
	}
	l := TreeDepth(root.left)
	r := TreeDepth(root.right)
	diff := l - r
	if diff > 1 || diff < -1 {
		return false
	}

	return IsBalanced(root.left) && IsBalanced(root.right)
}
func IsBalanced2(root *BinaryTreeNode) bool {
	depth := 0
	return isBalanced2(root, &depth)
}
func isBalanced2(root *BinaryTreeNode, depth *int) bool{
	if root == nil {
		*depth = 0
		return true
	}
	var left, right int
	if isBalanced2(root.left, &left) && isBalanced2(root.right, &right) {
		diff := left - right
		if diff <= 1 || diff >= -1 {
			if left > right {
				*depth += left
			} else {
				*depth += right
			}
			return true
		}
	}

	return false
}
func TestIsBalanced(t *testing.T) {
	a4 := &BinaryTreeNode{
		value:1,
	}
	a3 := &BinaryTreeNode{
		value:4,
	}
	a2 := &BinaryTreeNode{
		value:2,
		left:a4,
	}
	a1 := &BinaryTreeNode{
		value:3,
		left:a2,
		right:a3,
	}
	t.Run("1", func(t *testing.T) {
		d := IsBalanced(a1)
		fmt.Println(d)
	})
	t.Run("2", func(t *testing.T) {
		d := IsBalanced2(a1)
		fmt.Println(d)
	})
}

