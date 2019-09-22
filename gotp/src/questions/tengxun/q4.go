package main
//
// import (
// 	"fmt"
// 	"sort"
// )
//
// func main() {
// 	var n, k int
// 	fmt.Scan(&n)
// 	fmt.Scan(&k)
// 	numbers := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		x := 0
// 		fmt.Scan(&x)
// 		numbers[i] = x
// 	}
// 	sort.Ints(numbers)
// 	answer := make([]int, k)
// 	j := 0
// 	var pre, cur int
// 	for i := 0; i < n; i++ {
// 		if numbers[i] > 0 {
// 			cur = numbers[i]
// 			for i + 1 < n && numbers[i + 1] == cur {
// 				i++
// 			}
// 			cur = numbers[i]
// 			if j == 0 {
// 				answer[j] = cur
// 			} else {
// 				answer[j] = cur - pre
// 			}
// 			pre = cur
// 			j++
// 		}
// 		if j == k {
// 			break
// 		}
// 	}
// 	for _, x := range answer {
// 		fmt.Println(x)
// 	}
// }
