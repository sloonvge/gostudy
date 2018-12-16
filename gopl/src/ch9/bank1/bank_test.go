package bank1_test

import (
	"testing"
	"ch9/bank1"
	)

/** 
* Created by wanjx in 2018/12/1 22:29
**/

func TestBank(t *testing.T) {
	//done := make(chan struct{})

	// A 存款
	go func() {
		bank1.Deposit(200)
		t.Log("=", bank1.Balance())
		//done <- struct{}{}
	}()

	// B 存款
	go func() {
		bank1.Deposit(100)
		//done <- struct{}{}
	}()

	//<-done
	//<-done

	for i := 0; i < 100; i++ {
		if got, want := bank1.Balance(), 300; got != want {
			t.Errorf("Banlance = %d, want %d", got, want)
		}
	}

}
 
