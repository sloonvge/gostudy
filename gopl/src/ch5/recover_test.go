package ch5

import (
	"math/rand"
	"testing"
)

/** 
* Created by wanjx in 2019/3/19 23:11
**/

func recoverPanic() (r int){
	v := rand.Intn(100)
	defer func() {
		if p := recover(); p != nil {
			r = v
		}
	}()
	panic(v)
	return
}

func TestRecoverPanic(t *testing.T) {
	t.Logf("%d", recoverPanic())
	t.Logf("%d", rand.Intn(100))
}
