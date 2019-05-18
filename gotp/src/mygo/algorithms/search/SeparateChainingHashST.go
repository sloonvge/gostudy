package search


type SeparateChainingHashST struct {
	N int
	M int
	st []SequentialSearchST
}

func NewSeparateChiningHashST() *SeparateChainingHashST{
	return separateChiningHashST(997)
}

func separateChiningHashST(m int) *SeparateChainingHashST {
	return &SeparateChainingHashST{
		M: m,
		st: make([]SequentialSearchST, m),

	}
}

func (sc *SeparateChainingHashST) hash(k Key) int {
	return int(k) / sc.M
}

func (sc *SeparateChainingHashST) Get(k Key) Value {
	return sc.st[sc.hash(k)].Get(k)
}

func (sc *SeparateChainingHashST) Put(k Key, v Value) {
	sc.st[sc.hash(k)].Put(k, v)
}
