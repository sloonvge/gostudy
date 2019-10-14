package container

type Stack []*TreeNode

func (s *Stack) Push(t *TreeNode) {
	*s = append(*s, t)
}

func (s *Stack) Pop() *TreeNode {
	n := len(*s)
	t := (*s)[n - 1]
	*s = (*s)[:n - 1]
	return t
}

func (s *Stack) Top() *TreeNode {
	return (*s)[len(*s) - 1]
}


