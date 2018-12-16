package main

import (
	"10templatemathod/go/pattern"
)

/**
* Created by wanjx in 2018/12/8 21:08
**/

func main() {
	paperA := pattern.NewPaperA()
	paperB := pattern.NewPaperB()

	paperA.Question1()
	paperA.Question2()

	paperB.Question1()
	paperB.Question2()

}
