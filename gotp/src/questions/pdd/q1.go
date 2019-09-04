package main
//
// import (
// 	"fmt"
// 	"sort"
// )
//
// func main() {
// 	odds := make([]int, 0)
// 	evens := make([]int, 0)
// 	var num, N int
// 	for {
// 		n, _ := fmt.Scanf("%d,", &num)
// 		if n == 0 {
// 			break
// 		} else {
// 			if num & 1 == 1 {
// 				odds = append(odds, num)
// 			} else {
// 				evens = append(evens, num)
// 			}
// 		}
// 	}
// 	N = num
// 	if num & 1 == 1 {
// 		odds = odds[:len(odds) - 1]
// 	} else {
// 		evens = evens[:len(evens) - 1]
// 	}
// 	answer := make([]int, 0)
// 	sort.Ints(odds)
// 	sort.Ints(evens)
// 	for i := len(evens) - 1; i >= 0; i-- {
// 		if N == 0 {
// 			break
// 		}
// 		answer = append(answer, evens[i])
// 		N--
// 	}
// 	if N != 0 {
// 		for i := len(odds) - 1; i >= 0; i-- {
// 			if N == 0 {
// 				break
// 			}
// 			answer = append(answer, odds[i])
// 			N--
// 		}
// 	}
// 	if len(answer) > 0 {
// 		for i := 0; i < len(answer) - 1; i++ {
// 			fmt.Printf("%d,", answer[i])
// 		}
// 		fmt.Println(answer[len(answer) - 1])
// 	}
// }
