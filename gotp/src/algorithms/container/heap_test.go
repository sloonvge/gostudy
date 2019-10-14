package container

import (
	"fmt"
	"testing"
)

func TestInitHeap(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var h InitHeap
		h = MaxPQ(6)
		a := []int{1, 10, 2, 7, 3, 8}
		for i := 0; i < len(a); i++ {
			h.Insert(a[i])
			fmt.Println(h)
		}
		for i := 0; i < len(a); i++ {
			x := h.DelMax()
			fmt.Printf("%d ", x)
		}
	})
}
