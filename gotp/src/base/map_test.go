package base

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("close func", func(t *testing.T) {
		m := make(map[int]int)
		// m := map[string]int{
		// 	"a":1,
		// 	"b":2,
		// 	"c":3,
		// }
		for i := 0; i < 1024; i++ {
			m[i] = i
		}
		for k,v := range m  {
			go func(kk *int,vv *int) {
				fmt.Printf("%v,%v\n",*kk,*vv)
			}(&k,&v)
		}
	})
}
