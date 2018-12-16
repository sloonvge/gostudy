package bank1

/** 
* Created by wanjx in 2018/12/1 22:22
**/

var balance int

func Deposit(amount int) { balance = balance + amount }

func Balance() int { return balance }
 
