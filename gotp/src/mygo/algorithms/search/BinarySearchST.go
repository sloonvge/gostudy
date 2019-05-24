package search


type BinarySearchST struct {
	N int
	keys []Key
	values []Value
}

func NewBinarySearchST(n int) *BinarySearchST{
	keys := make([]Key, n)
	values := make([]Value, n)
	return &BinarySearchST{
		keys: keys,
		values:values,
		N:0,
	}
}

func (b *BinarySearchST) Get(k Key) Value {
	if b.N == 0 {
		return 0
	}

	i := b.Rank(k)

	if i < b.N && b.keys[i] == k {
		return b.values[i]
	} else {
		return 0
	}
}

func (b *BinarySearchST) Put(k Key, v Value) {
	i := b.Rank(k)

	if i < b.N && b.keys[i] == k {
		b.values[i] = v
		return
	}

	for j := b.N; j > i; j-- {
		b.keys[j] = b.keys[j - 1]
		b.values[j] = b.values[j - 1]
	}
	b.keys[i] = k
	b.values[i] = v
	b.N++
}

func (b *BinarySearchST) Rank(k Key) int {
	lo := 0
	hi := b.N - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		if b.keys[mid] == k {
			return mid
		} else if b.keys[mid] < k {
			lo = mid +1
		} else {
			hi = mid - 1
		}
	}
	return lo
}



