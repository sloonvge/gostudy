package tree

import (
	"testing"
)

/** 
* Created by wanjx in 2019/5/21 23:29
**/

func Setup() {

}

func initTree(a []int, lo int) *Tree {
	if a[lo] == -1 {
		return nil
	}
	tree := &Tree{
		Val:a[lo],
	}
	l := 2 * lo + 1
	r := 2 * lo + 2
	if l > len(a) -1 {
		tree.Left = nil
	} else {
		tree.Left = initTree(a, l)
	}
	if r > len(a) -1 {
		tree.Right = nil
	} else {
		tree.Right = initTree(a, r)
	}
	return tree
}

//      5
//   4     6
// 2  3   7  8
//     1
func TestTraverse(t *testing.T) {
	// N := 8
	// input := make([]int, 0)
	// for i := 0; i < N; i++ {
	// 	input = append(input, i)
	// }
	input := []int{5, 4, 6, 2, 3, 7, 8, -1, -1, -1, 1}
	tree := initTree(input, 0)
	_ = tree
}