package search

import (
	"testing"
	"fmt"
	"crypto/sha256"
	"strconv"
)

func TestSequentialSearchST_Get(t *testing.T) {
	n := 1000
	st := SequentialSearchST{}
	for i:=0; i < n; i++ {
		st.Put(Key(i), Value(i))
	}

	for i := 0; i < n; i++ {
		fmt.Printf("k: %d - v: %d\n", i, st.Get(Key(i)))
	}
}

func TestBinarySearchST_Get(t *testing.T) {
	n := 1000
	st := NewBinarySearchST(n)
	for i:=0; i < n; i++ {
		st.Put(Key(i), Value(i))
	}

	for i := 0; i < n; i++ {
		fmt.Printf("k: %d - v: %d\n", i, st.Get(Key(i)))
	}
}

func TestBST_Put(t *testing.T) {
	n := 10000
	st := NewBST()

	for i := 0; i < n; i++ {
		st.Put(Key(i), Value(i))
	}

	st.Show()


	fmt.Printf("%d %d\n", 5, st.Get(5))
	fmt.Printf("%d %d\n", 4, st.Get(4))
}

func TestRedBlackBST_Put(t *testing.T) {
	n := 10000
	st := NewRedBlackBST()

	for i := 0; i < n; i++ {
		st.Put(Key(i), Value(i))
	}

	st.Show()

	fmt.Printf("%d %d\n", 5, st.Get(5))
	fmt.Printf("%d %d\n", 4, st.Get(4))
}

func BenchmarkRedBlackBST_Put(b *testing.B) {
	st := NewBST()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		st.Put(Key(i), Value(i))
	}

	// st.Show()
}

func BenchmarkBST_Put(b *testing.B) {
	st := NewRedBlackBST()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		st.Put(Key(i), Value(i))
	}
}

func TestHashSHA236(t *testing.T) {
	s := sha256.Sum256([]byte(strconv.Itoa(111)))
	fmt.Printf("%v\n", s)

}

func TestSeparateChainingHashST_Put(t *testing.T) {
	n := 10000
	st := NewSeparateChiningHashST()
	for i := 0; i < n; i++ {
		st.Put(Key(i), Value(i))
	}
	fmt.Printf("%v\n", st)

	fmt.Printf("%d %d\n", 5, st.Get(5))
	fmt.Printf("%d %d\n", 4, st.Get(4))
}