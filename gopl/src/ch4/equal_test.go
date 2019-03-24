package main

import (
	"testing"
	"reflect"
)

func TestEqualArray(t *testing.T)  {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2, 4}

	t.Logf("a == b %v\n", a == b)
	t.Logf("a == c %v\n", a == c)
}

func TestEqualSlice(t *testing.T)  {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}

	t.Logf("a == b %v\n", reflect.DeepEqual(a, b))
	t.Logf("a == c %v\n", reflect.DeepEqual(a, c))
}