package base

import (
	"fmt"
	"reflect"
	"testing"
)

type Err struct {
	err string
}

func (e *Err) Error() string {
	return e.err
}

func returnErr() *Err {
	return nil
}

func ErrorNil() (err error) {
	if e := returnErr(); e != nil {
		return e
	}
	return
}


func TestNilError(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var err error
		err = returnErr()
		fmt.Println(err, err != nil) // <nil> true
	})
	t.Run("", func(t *testing.T) {
		var err error
		err = ErrorNil()
		fmt.Println(err, err != nil) // <nil> false
	})
	t.Run("", func(t *testing.T) {
		err := returnErr()
		fmt.Println(err, err != nil) // <nil> false
		fmt.Println(reflect.TypeOf(err))
	})
}

type Summer interface{
	Sum() int
}

type ints []int
func (i ints) Sum() int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}

func TestNilInterface(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var s Summer
		fmt.Println(s == nil)
	})
	t.Run("", func(t *testing.T) {
		var i ints
		var s Summer = i
		fmt.Println(i == nil, s.Sum()) // true 0
	})
}

