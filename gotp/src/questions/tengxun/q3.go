package main
//
// import (
// 	"fmt"
// 	"math"
// )
//
// // 2
// // 4
// // 5 9 4 7
// // 8
// // 9 13 18 10 12 4 18 3
//
// // 12 13
// // 43 44
// func main() {
// 	var T, n, x, sum int
// 	fmt.Scan(&T)
// 	for i := 0; i < T; i++ {
// 		fmt.Scan(&n)
// 		sum = 0
// 		numbers := make([]int, n)
// 		for j := 0; j < n; j++ {
// 			fmt.Scan(&x)
// 			numbers[j] = x
// 			sum += x
// 		}
// 		t, n := minGroups(numbers, n, sum)
// 		fmt.Println(n, t)
// 	}
//
// }
//
// func minGroups(nums []int, n int, sum int) (t, no int) {
// 	min := math.MaxInt64
// 	t = 0
// 	no = 0
// 	helper(nums, sum, 0,0, &min,&t, &no, 0, 0)
// 	return
// }
//
// func helper(nums []int, sum, m, curpos int, min, t, n *int, taken, noTaken int) {
// 	if curpos >= len(nums) {
// 		return
// 	}
// 	if m == len(nums) / 2 {
// 		// fmt.Println(m, taken, noTaken)
// 		if 2 * taken - sum > 0 && 2 * taken - sum < *min {
// 			*min = 2 * taken - sum
// 			*t = taken
// 			*n = sum - taken
// 		} else if sum - 2 * taken > 0 && sum - 2 * taken < *min {
// 			*min = sum - 2 * taken
// 			*t = sum - taken
// 			*n = taken
// 		}
// 		return
// 	}
// 	helper(nums, sum, m + 1, curpos + 1, min, t, n, taken + nums[curpos], noTaken)
// 	helper(nums, sum, m, curpos + 1, min, t,n,taken, noTaken + nums[curpos])
// }