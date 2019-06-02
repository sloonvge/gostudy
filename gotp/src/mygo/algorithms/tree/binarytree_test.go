package tree

import "testing"

func setup() *Tree {
	//     1
	//   2   3
	// 4  5    6
	//   7 8  9
	node9 := &Tree{nil, 9, nil}
	node8 := &Tree{nil, 8, nil}
	node7 := &Tree{nil, 7, nil}
	node6 := &Tree{node9, 6, nil}
	node5 := &Tree{node7, 5, node8}
	node4 := &Tree{nil, 4, nil}
	node3 := &Tree{nil, 3, node6}
	node2 := &Tree{node4, 2, node5}
	node1 := &Tree{node2, 1, node3}

	return node1
}

func teardown() {

}

func TestTree(t *testing.T) {
	tree := setup()
	t.Run("PreOrder traversal", func(t *testing.T) {
		PreOrderTraverse(tree) // 1 2 4 5 7 8 3 6 9
	})
	t.Run("PreOrderIteration", func(t *testing.T) {
		PreOrderIteration(tree)
	})

	t.Run("InOrderTraversal", func(t *testing.T) {
		InOrderTraverse(tree) // 4 2 7 5 8 1 3 9 6
	})
	t.Run("InorderIteration", func(t *testing.T) {
		InOrderIteration(tree)
	})
	teardown()
}
