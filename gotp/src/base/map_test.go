package base

import (
	"fmt"
	"reflect"
	"testing"
	"time"
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
		f := func(k, v int) {
			fmt.Printf("%v,%v\n",k,v)
		}
		for k,v := range m  {
			f(k, v)
		}
	})
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
	t.Run("", func(t *testing.T) {
		m := make(map[int]int)
		for i := 0; i < 5; i++ {
			m[i] =i
		}
		for i := 5; i < 10; i++ {
			m[i] = i
			delete(m, i - 5)
			fmt.Println(m)
		}
	})
	t.Run("concurrency", func(t *testing.T) {
		m := make(map[int]int)
		go func() {
			for i := 0; i < 100; i++ {
				m[i] = i
			}
		}()
		for i := 0; i < 100; i++ {
			go func() {
				delete(m, i)
			}()
		}
		go func() {
			fmt.Println(m)
		}()
		time.Sleep(time.Second)
	})
	t.Run("len cap", func(t *testing.T) {
		m := make(map[int]int)
		mt := reflect.ValueOf(m)
		fmt.Printf("len=%d addr=%v\n", len(m), mt.Pointer())
		m[0] = 0
		fmt.Printf("len=%d addr=%v\n", len(m), mt.Pointer())
		for i := 1; i < 10; i++ {
			m[i] = i
		}
		mt = reflect.ValueOf(m)
		fmt.Printf("len=%d addr=%v\n", len(m), mt.Pointer())
	})
}
