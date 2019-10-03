package base

import (
	"fmt"
	"reflect"
	"testing"
)

type cat struct {
	name string
	age int
}

func TestStruct(t *testing.T) {
	t.Run("", func(t *testing.T) {
		c1 := cat{"cat1", 1}
		c2 := cat{"cat1", 1}
		fmt.Println(c1 == c2)
		fmt.Println(&c1 == &c2)
		fmt.Println(reflect.DeepEqual(c1, c2))
		fmt.Println(reflect.DeepEqual(&c1, &c2))
	})
}
