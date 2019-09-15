package main

import "fmt"

type BinaryTreeNode struct {
	left *BinaryTreeNode
	right *BinaryTreeNode
	value int
	path int
}

func NewBinaryTreeNode(n int) *BinaryTreeNode{
	return &BinaryTreeNode{value:n}
}

func searchNodeAndAdd(root *BinaryTreeNode, parent, i int) {

	node := searchParent(root, parent)
	if node == nil {
		return
	}
	if node.left == nil {
		node.left = NewBinaryTreeNode(i)
	} else if node.right == nil {
		node.right = NewBinaryTreeNode(i)
	}
}

func searchParent(root *BinaryTreeNode, n int) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	if root.value == n {
		return root
	}
	node := searchParent(root.left, n)
	if node != nil {
		return node
	}
	node  = searchParent(root.right, n)
	return node
}

var values map[int]map[int]int

func main() {
	var n, i, j, v int
	values = make(map[int]map[int]int, 0)
	fmt.Scan(&n)
	root := NewBinaryTreeNode(1)
	for k := 0; k < n - 1; k++ {
		fmt.Scanf("%d %d %d", &i, &j, &v)
		if _, ok := values[i]; !ok {
			values[i] = make(map[int]int)
		}
		values[i][j] = v
		searchNodeAndAdd(root, i, j)
	}
	longestPath(root)
	answer := make([]int, n)
	recordPath(root, &answer)
	for i := 0; i < len(answer); i++ {
		fmt.Printf("%d ", answer[i])
	}
}

func longestPath(root *BinaryTreeNode) int{
	if root == nil || (root.left == nil && root.right == nil) {
		return 0
	}
	root.left.path = longestPath(root.left)
	root.right.path = longestPath(root.right)
	root.path =  max(root.left.path + values[root.value][root.left.value],
		root.right.path + values[root.value][root.right.value])

	return root.path
}

func recordPath(root *BinaryTreeNode, ans *[]int) {
	if root == nil {
		return
	}
	(*ans)[root.value - 1] = root.path
	recordPath(root.left, ans)
	recordPath(root.right, ans)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
