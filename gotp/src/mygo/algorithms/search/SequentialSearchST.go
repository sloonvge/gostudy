package search

type SequentialSearchST struct {
	first *Node
}

type Value int
type Key int

type Node struct {
	Key Key
	Value Value
	Next *Node
}

func (s *SequentialSearchST) NewNode(k Key, v Value, n *Node) {
	node := &Node{
		Key:k,
		Value:v,
		Next:n,
	}
	s.first = node
}

func (s *SequentialSearchST) Get(k Key) Value {
	for n := s.first; n != nil; n = n.Next {
		if n.Key == k {
			return n.Value
		}
	}
	return 0
}

func (s *SequentialSearchST) Put(k Key, v Value) {
	for n := s.first; n != nil; n = n.Next {
		if n.Key == k {
			n.Value = v
		}
	}
	s.NewNode(k, v, s.first)
}
