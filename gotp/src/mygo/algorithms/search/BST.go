package search

import "fmt"

type BST struct {
	root *BSTNode
}

type BSTNode struct {
	Key Key
	Value Value
	Left *BSTNode
	Right *BSTNode
	N int
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) NewNode(k Key, v Value) *BSTNode{
	return &BSTNode{
		Key:k,
		Value:v,
		N:1,
	}
}

func (b *BST) Get(k Key) Value {
	return b.get(k, b.root)
}

func (b *BST) get(k Key, tree *BSTNode) (v Value) {
	if tree == nil {
		return -1
	}

	if tree.Key == k {
		v = tree.Value
		return
	} else if k < tree.Key {
		v = b.get(k, tree.Left)
	} else {
		v = b.get(k, tree.Right)
	}

	return
}


func (b *BST) Put(k Key, v Value) {
	b.root = b.put(k, v, b.root)
}

func (b *BST) put(k Key, v Value, tree *BSTNode) *BSTNode{
	if tree == nil {
		newNode := b.NewNode(k, v)
		tree = newNode
	}

	if k == tree.Key {
		tree.Value = v
	} else if k < tree.Key {
		tree.Left = b.put(k, v, tree.Left)
	} else {
		tree.Right = b.put(k, v, tree.Right)
	}
	tree.N = b.Size(tree.Left) + b.Size(tree.Right) + 1
	return tree
}

func (b *BST) Size(node *BSTNode) int {
	if node == nil {
		return 0
	}

	return node.N
}

func (b *BST) Show() {
	b.show(b.root)
}

func (b *BST) show(tree *BSTNode) {
	if tree == nil {
		return
	}
	fmt.Printf("%d:%d:%d\n", tree.Key, tree.Value, tree.N)
	if tree.Left != nil {
		b.show(tree.Left)
	}

	if tree.Right != nil {
		b.show(tree.Right)
	}
}


