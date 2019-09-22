package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	var k int
// 	m := 0
// 	fmt.Scan(&m)
// 	set := make(map[int]struct{})
// 	min := 0
// 	for i := 0; i < m; i++ {
// 		fmt.Scan(&k)
// 		if i == 0 {
// 			min = k
// 		}
// 		set[k] = struct {}{}
// 		if k < min {
// 			min = k
// 		}
// 	}
//
// 	n := min
// 	max := 0
// 	l := 1
// 	delete(set, n)
// 	for len(set) != 0 {
// 		if _, ok := set[n + 1]; ok {
// 			l++
// 			delete(set, n + 1)
// 		} else {
// 			if l > max {
// 				max = l
// 			}
// 			l = 0
// 		}
// 		n++
// 	}
//
// 	if l == m {
// 		fmt.Println(l)
// 	} else {
// 		fmt.Println(max)
// 	}
// }
