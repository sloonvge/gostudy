package search

import "fmt"

const RED  = "red"
const BLACK  = "black"

type RedBlackBST struct {
	root *RedBlackNode
}

type RedBlackNode struct {
	Key Key
	Value Value
	Left *RedBlackNode
	Right *RedBlackNode
	Color string
	N int
}

func NewRedBlackBST() *RedBlackBST {
	return &RedBlackBST{}
}

func NewRedBlackNode(k Key, v Value, c string) *RedBlackNode {
	return &RedBlackNode{
		Key:k,
		Value:v,
		N:1,
		Color:c,
	}
}

func (r *RedBlackNode) isRed() bool {
	if r == nil {
		return false
	}
	return r.Color == RED
}

func (r *RedBlackNode) rotateLeft() *RedBlackNode{
	t := r.Right
	r.Right = t.Left
	t.Left = r

	t.Color = r.Color
	r.Color = RED
	t.N = r.N
	r.N = 1 + r.Left.size() + r.Right.size()

	return t
}

func (r *RedBlackNode) rotateRight() *RedBlackNode{
	t := r.Left
	r.Left = t.Right
	t.Right = r

	t.Color = r.Color
	r.Color = RED
	t.N = r.N
	r.N = 1 + r.Left.size() + r.Right.size()
	return t
}

func (r *RedBlackNode) flipColors() {
	r.Color = RED
	r.Left.Color = BLACK
	r.Right.Color = BLACK
}

func (r *RedBlackNode) size() int {
	if r == nil {
		return 0
	}

	return r.N
}

func (r *RedBlackBST) Get(k Key) Value {
	return r.get(k, r.root)
}

func (r *RedBlackBST) get(k Key, tree *RedBlackNode) (v Value) {
	if tree == nil {
		return -1
	}

	if k < tree.Key {
		v = r.get(k, tree.Left)
	} else if k > tree.Key {
		v = r.get(k, tree.Right)
	} else {
		v = tree.Value
	}

	return
}

func (r *RedBlackBST) Put(k Key, v Value) {
	r.root = r.put(k, v, r.root)
	r.root.Color = BLACK
}

func (r *RedBlackBST) put(k Key, v Value, tree *RedBlackNode) *RedBlackNode{
	if tree == nil {
		return NewRedBlackNode(k, v, RED)
	}

	if k < tree.Key {
		tree.Left = r.put(k, v, tree.Left)
	} else if k > tree.Key {
		tree.Right = r.put(k, v, tree.Right)
	} else {
		tree.Value = v
	}

	if tree.Right.isRed() && !tree.Left.isRed() {
		tree = tree.rotateLeft()
	}
	if tree.Left.isRed() && tree.Left.Left.isRed() {
		tree = tree.rotateRight()
	}
	if tree.Left.isRed() && tree.Right.isRed() {
		tree.flipColors()
	}

	return tree
}

func (r *RedBlackBST) Show() {
	r.show(r.root)
}

func (r *RedBlackBST) show(tree *RedBlackNode) {
	if tree == nil {
		return
	}
	fmt.Printf("%d:%d:%d\n", tree.Key, tree.Value, tree.N)
	if tree.Left != nil {
		r.show(tree.Left)
	}

	if tree.Right != nil {
		r.show(tree.Right)
	}
}

