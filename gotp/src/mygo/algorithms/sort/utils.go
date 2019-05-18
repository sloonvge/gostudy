package sort

import "fmt"

/** 
* Created by wanjx in 2019/4/6 23:31
**/

func Show(a []int) {
	fmt.Printf("%v\n", a)
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}