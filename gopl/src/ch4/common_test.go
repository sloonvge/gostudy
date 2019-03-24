package main

import (
	"testing"
)

func TestRangeTmpVar(t *testing.T) {
	a := [3]int{1, 2, 3}
	for i, v := range a {
		b := v
		t.Log(i, &b, &a[i], "\n")
	}
	t.Logf("%x", a)
}

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i , j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}
// 翻转数组
func TestReverseArray(t *testing.T) {
	s := []int{1 , 2, 3, 4, 5}
	reverse(s)
	t.Logf("%v\n", s)
}
// 向左旋转n个元素
func TestRotateLeftN( t*testing.T) {
	s := []int{1 , 2, 3, 4, 5}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	t.Logf("%v\n", s)
}