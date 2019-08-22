package redis

import (
	"fmt"
	"testing"
)

func poker(a []int) (ret []int){
	n := len(a)
	for len(a) != 0 {
		ret = append(ret, a[0])
		if len(a) == 1 {
			break
		}
		tmp := a[1]
		a = a[2:]
		a = append(a, tmp)
	}
	if len(ret) != n {
		fmt.Println(len(ret), n)
		panic("error in compute")
	}
	return ret
}

func TestPoker(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	out := poker(in)

	fmt.Printf("%v\n", in)
	fmt.Printf("%v\n", out)
}
