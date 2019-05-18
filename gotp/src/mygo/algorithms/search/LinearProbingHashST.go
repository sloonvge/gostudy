package search

type LinearProbingHashST struct {
	N int
	M int
	keys []Key
	values []Value

}

func NewLinearProbingHashST() *LinearProbingHashST {
	st := &LinearProbingHashST{}
	st.M = 16
	st.keys = make([]Key, st.M)
	st.values = make([]Value, st.M)
	return st
}

func (lp *LinearProbingHashST) hash(k Key) int {
	return int(k) % 997
}

func (lp *LinearProbingHashST) Get(k Key) Value {
	for i := lp.hash(k); lp.keys[i] != 0; i = (i + 1) % lp.M {
		if lp.keys[i] == k {
			return lp.values[i]
		}
	}

	return -1
}

func (lp *LinearProbingHashST) Put(k Key, v Value) {
	var i  int
	for i := lp.hash(k); lp.keys[i] != 0; i = (i + 1) % lp.M {
		if lp.keys[i] == k {
			lp.values[i] = v
			return
		}
	}
	lp.keys[i] = k
	lp.values[i] = v
	lp.N++
}
