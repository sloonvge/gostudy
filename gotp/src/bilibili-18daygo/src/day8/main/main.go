package main

import (
	"fmt"
	"runtime"
)

/** 
* Created by wanjx in 2019/7/7 21:48
**/
 
func main() {
	num := runtime.NumCPU()
	fmt.Println(num)
}