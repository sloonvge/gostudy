package tree

import "fmt"

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func PreOrderTraverse(t *Tree) {
	if t == nil {return }
	visit(t.Value)
	PreOrderTraverse(t.Left)
	PreOrderTraverse(t.Right)
}
func PreOrderIteration(t *Tree) {
	if t == nil {
		return
	}
	s := new(Stack)
	s.Push(t)
	for len(*s) != 0{
		t = s.Pop()
		visit(t.Value)
		if t.Right != nil {
			s.Push(t.Right)
		}
		if t.Left != nil {
			s.Push(t.Left)
		}
	}
}

func InOrderTraverse(t *Tree) {
	if t == nil {
		return
	}
	InOrderTraverse(t.Left)
	visit(t.Value)
	InOrderTraverse(t.Right)
}
func InOrderIteration(t *Tree) {
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
			visit(t.Value)
			t = t.Right
		}
	}
}

func visit(i int) {
	fmt.Printf("%d ", i)
}

