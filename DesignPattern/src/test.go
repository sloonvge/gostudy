package main

import (
	"fmt"
	"strings"
)

/** 
* Created by wanjx in 2019/3/10 16:39
**/
 
func main() {
	// n := 0
	// input := make([]string, 2)
	le := "(()"//input[0]
	ri := "())"//input[1]
	all := ""
	// fmt.Scan(&n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&input[i])
	// }
	les := strings.Split(le)
	ris := strings.Split(ri)
	for i := 0; i < len(ris); i++ {
		for j := 0; j < len(les); j++ {

		}
	}
	all := le + ri
	ans := checkValid(all)
	fmt.Println(ans)
}


func checkValid(s string) bool {
	ret := true

	s = strings.TrimSpace(s)
	words := strings.Split(s, "")
	countsMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		countsMap[words[i]]++
	}
	if countsMap["("] != countsMap[")"] {
		ret = false
		return ret
	}

	if words[0] == ")" || words[len(words) - 1] == "(" {
		ret = false
		return ret
	}

	return ret
}


// func main() {
// 	n:=0
// 	ans:=0
//
// 	fmt.Scan(&n)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			x:=0
// 			fmt.Scan(&x)
// 			ans = ans + x
// 		}
// 	}
// 	fmt.Printf("%d\n",ans)
// }