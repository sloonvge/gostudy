package shape

import (
	"fmt"
	"testing"
)

/** 
* Created by wanjx in 2019/4/14 21:47
**/
 
func TestPoint_Distance(t *testing.T) {
	p := Point{0, 3}
	q := Point{4, 0}

	f := p.Distance

	var zero Point
	fmt.Println(f(&q))

	fmt.Println(f(&zero))

	t.Run("FuncExpr", func(t *testing.T) {
		f := (*Point).Distance

		fmt.Println(f(&p, &q))
		fmt.Println(f(&q, &p))
	})

}

