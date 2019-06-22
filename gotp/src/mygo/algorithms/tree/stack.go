package tree

type Stack []*Tree

func (s *Stack) Push(t *Tree) {
	*s = append(*s, t)
}

func (s *Stack) Pop() *Tree {
	n := len(*s)
	t := (*s)[n - 1]
	*s = (*s)[:n - 1]
	return t
}

func (s *Stack) Top() *Tree {
	return (*s)[len(*s) - 1]
}


