package main
//
// import (
// 	"fmt"
// 	"sort"
// )
//
// 2
// 4
// 5 9 4 7
// 8
// 9 13 18 10 12 4 18 3
// 10
// func main()  {
// 	var n, a, b, M int
// 	fmt.Scan(&n)
// 	numbers := make([]int, 0)
// 	for i := 0; i < n; i++ {
// 		fmt.Scanf("%d %d", &a, &b)
// 		for j := 0; j < a; j++ {
// 			numbers = append(numbers, b)
// 		}
// 		M += a
// 	}
// 	sort.Ints(numbers)
// 	answer := 0
// 	for i := 0; i < M / 2; i++ {
// 		k := numbers[i] + numbers[M - 1 - i]
// 		if k > answer {
// 			answer = k
// 		}
// 	}
//
// 	fmt.Println(answer)
// }
