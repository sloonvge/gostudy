package tree

import (
	"algorithms/container"
	"testing"
)

func setup() *container.Tree {
	//     1
	//   2   3
	// 4  5    6
	//   7 8  9
	node9 := &container.Tree{nil, 9, nil}
	node8 := &container.Tree{nil, 8, nil}
	node7 := &container.Tree{nil, 7, nil}
	node6 := &container.Tree{node9, 6, nil}
	node5 := &container.Tree{node7, 5, node8}
	node4 := &container.Tree{nil, 4, nil}
	node3 := &container.Tree{nil, 3, node6}
	node2 := &container.Tree{node4, 2, node5}
	node1 := &container.Tree{node2, 1, node3}

	return node1
}

func teardown() {

}

func TestTree(t *testing.T) {
	tree := setup()
	t.Run("PreOrder traversal", func(t *testing.T) {
		container.PreOrderTraverse(tree) // 1 2 4 5 7 8 3 6 9
	})
	t.Run("PreOrderIteration", func(t *testing.T) {
		container.PreOrderIteration(tree)
	})

	t.Run("InOrderTraversal", func(t *testing.T) {
		container.InOrderTraverse(tree) // 4 2 7 5 8 1 3 9 6
	})
	t.Run("InorderIteration", func(t *testing.T) {
		container.InOrderIteration(tree)
	})
	teardown()
}
