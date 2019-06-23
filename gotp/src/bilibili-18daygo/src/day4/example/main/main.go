package main

import (
	"strings"
	"fmt"
)

/** 
* Created by wanjx in 2019/6/22 23:17
**/
 
func adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func makeSuffix(name string) func(string) string {
	return func(suffix string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func close() {
	f := makeSuffix("domo")
	fmt.Println(f(".jpg"))
	fmt.Println(f(".png"))
}

func main() {
	// f := adder()
	// fmt.Println(f(1))
	// fmt.Println(f(10))
	// close()
	
}