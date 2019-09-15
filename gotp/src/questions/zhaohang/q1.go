package main
//
// import (
// 	"fmt"
// 	"math"
// )
//
// func main() {
// 	x := ""
// 	fmt.Scan(&x)
// 	mods := make(map[int]int)
// 	answer := make([]int, len(x))
// 	for i := 0; i < len(x); i++ {
// 		if x[i] == 'R' {
// 			for j := i + 1; j < len(x); j++ {
// 				if x[j] == 'L'{
// 					mods[i] = j - i - 1
// 					m := (int(math.Pow(10, 100) ) - mods[i]) % 2
// 					if m == 0 {
// 						answer[j - 1]++
// 					} else {
// 						answer[j]++
// 					}
// 					break
// 				}
// 			}
// 		} else if x[i] == 'L' {
// 			for j := i - 1; j >= 0; j-- {
// 				if x[j] == 'R' {
// 					mods[i] = i - j - 1
// 					m := (int(math.Pow(10, 100) ) - mods[i]) % 2
// 					if m == 0 {
// 						answer[j + 1]++
// 					} else {
// 						answer[j]++
// 					}
// 					break
// 				}
// 			}
// 		}
// 	}
//
// 	for i := 0; i < len(x); i++ {
// 		fmt.Printf("%d ", answer[i])
// 	}
// }
//
//
//
