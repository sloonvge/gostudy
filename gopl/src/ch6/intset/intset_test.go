package intset

import (
	"fmt"
	"testing"
)

/** 
* Created by wanjx in 2019/4/14 22:54
**/

// 1 << a 和 a << 1不一样
func TestBit(t *testing.T) {
	var b uint = 5
	c := b << 1
	d := b >> 1
	fmt.Printf("%v-%v-%v\n", b, c, d)
	fmt.Printf("%b-%b-%b\n", b, c, d)
	fmt.Printf("%T-%T-%T\n", b, c, d)
	fmt.Printf("%v\n", 1 >> 0)
}

func TestIntSet_Add(t *testing.T) {
	var x IntSet

	x.Add(1)
	x.Add(10)
	x.Add(8)
	fmt.Printf("%v\n", x)
	fmt.Println(x.Has(10))
	fmt.Println(x.String())
	fmt.Printf("Len x: %d\n", x.Len())
	x.Remove(1)
	fmt.Printf("Remove : %s\n", x.String())
	x.Remove(4)
	x.Clear()
	fmt.Printf("Clear: %s\n", x.String())
	x.AddAll(1, 8, 10)

	y := x.Copy()
	fmt.Printf("%s\n", y.String())
	x.Remove(1)
	fmt.Printf("x:%s\ty:%s\n", x.String(), y.String())
}

func TestIntSet(t *testing.T) {
	var s, u, v IntSet
	s.AddAll(1, 2, 3)

	u.AddAll(0, 4, 5)
	v.AddAll(1, 4, 5)
	t.Run("IntersectWith", func(t *testing.T) {

		s.IntersectWith(&u)
		t.Logf("%v %d\n", s.String(), s.Len())

	})

	t.Run("DifferenceWith", func(t *testing.T) {
		s.DifferenceWith(&u)
		u.DifferenceWith(&v)
		t.Logf("%v\n%v\n", s.String(), u.String())
	})

	t.Run("SymmetricDifference", func(t *testing.T) {
		s.SymmetricDifference(&u)
		u.SymmetricDifference(&v)
		t.Logf("%v\n%v\n", s.String(), u.String())

	})

	t.Run("Elems", func(t *testing.T) {
		fmt.Printf("%d\n", s.Elems())
		for _, e := range s.Elems() {
			t.Logf("%d\t", e)
		}
		t.Log("\n")
	})

}

func TestIntSet_Len(t *testing.T) {
	a := 32 << (^uint(0) >> 63)
	print(a, "\n")
	fmt.Printf("%b %v\n", 32 << 0, 32<<0)
}
